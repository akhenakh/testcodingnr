# Triplestat

## Running and Building

To execute the command:

```sh
cd cmd/triplestat
go run main.go
```

Run the test:

```sh
go test -v ./...
```

Or use the Taskfile.yml (See [Taskfile](https://taskfile.dev/#/)).

```sh
task -l
```

## Example

Example usage:

```sh
cat testdata/test01.txt | ./cmd/tripletstat/tripletstat   
```

```sh
./cmd/tripletstat/tripletstat   testdata/test01.txt 
```

## Docker

A Dockerfile is provided, build as follow:

```sh
docker build  -t codingtestnr:latest .  
```

Testing and running:

```sh
docker run --rm -it -v $(pwd)/testdata:/data codingtestnr:latest /data/test00.txt
```

## gRPC server

To perform computation at large scale, relying on Kubernetes Jobs is probably good enough.
But to avoid containers starting penalties a gRPC server is provided.

Only one method is provided:

```proto
rpc Compute(ComputeRequest) returns (ComputeResponse) {}
```

Run the server:

```sh
./cmd/tripletserver/tripletserver 
```

And the client:

```sh
./tripletclient -path ../../testdata/test01.txt 
```


## Protobuf

Generate with [buf](https://docs.buf.build/installation/)
```
buf generate
```

## What could be done

- The main func is not unitested
- Context cancelation not implemented
- To reduce memory usage, rather than storing the triplet using 3 strings in memory, use a Hash, beware of collision

## Organization of the code

3 packages:
	- `wordsplit` takes an `io.Reader` and returns an iterator that returns 3 words.
	  Independent, tests provided.
	- `worstat` takes triplet as input and increment occurence
	  Independent, test provided.
	- `tripletsvc` a gRPC service & server implementation to compute stats

2 commands:
	- `tripletstat` the CLI has requested by coding test part 1
	- `tripletserver` gRPC server
	- `tripletclient` gRPC client


