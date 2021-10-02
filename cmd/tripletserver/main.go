package main

import (
	"context"
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/akhenakh/codingtestnr/tripletsvc"
)

const appName = "tripletserver"

var (
	version = "no version from LDFLAGS"

	httpMetricsPort = flag.Int("httpMetricsPort", 8888, "http port")
	grpcPort        = flag.Int("grpcPort", 9200, "gRPC API port")

	grpcServer        *grpc.Server
	httpMetricsServer *http.Server
)

func main() {
	flag.Parse()

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "app", appName)
	logger = level.NewFilter(logger, level.AllowAll())

	stdlog.SetOutput(log.NewStdlibAdapter(logger))

	level.Info(logger).Log("msg", "Starting app", "version", version)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// catch termination
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	// web server metrics
	g.Go(func() error {
		httpMetricsServer = &http.Server{
			Addr:         fmt.Sprintf(":%d", *httpMetricsPort),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		level.Info(logger).Log("msg", fmt.Sprintf("HTTP Metrics server serving at :%d", *httpMetricsPort))

		// Register Prometheus metrics handler.
		http.Handle("/metrics", promhttp.Handler())

		if err := httpMetricsServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	// gRPC Server
	g.Go(func() error {
		addr := fmt.Sprintf(":%d", *grpcPort)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			level.Error(logger).Log("msg", "gRPC server: failed to listen", "error", err)
			os.Exit(2)
		}

		grpcServer = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
			)),
		)

		s := &tripletsvc.Server{}
		tripletsvc.RegisterTripletServiceServer(grpcServer, s)
		level.Info(logger).Log("msg", fmt.Sprintf("gRPC server serving at %s", addr))

		return grpcServer.Serve(ln)
	})

	select {
	case <-interrupt:
		cancel()
		break
	case <-ctx.Done():
		break
	}

	level.Warn(logger).Log("msg", "received shutdown signal")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if httpMetricsServer != nil {
		_ = httpMetricsServer.Shutdown(shutdownCtx)
	}

	if err := g.Wait(); err != nil {
		level.Error(logger).Log("msg", "server returning an error", "error", err)
		os.Exit(2)
	}
}
