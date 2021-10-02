# Triplestat

## Running and Building

To execute the command:
```
cd cmd/triplestat
go run main.go
```

Run the test:
```
go test -v ./...
```

Or use the Taskfile.yml (See [Taskfile](https://taskfile.dev/#/)).
```
task -l

## Example
```
Example usage:
```
cat testdata/test01.txt | ./cmd/tripletstat/tripletstat   
```

```
./cmd/tripletstat/tripletstat   testdata/test01.txt 
```

## Docker

A Dockerfile is provided build as follow:
```
nerdctl build  -t codingtestnr:latest .  
```
or with Docker:
```
docker build  -t codingtestnr:latest .  
```

## What could be done
Stopped after 3h25

- The main func is untested
- To answer the scaling question with a Kubernetes Jobs we can run this at scale
- To avoid container start penalty for each runs a simple gRPC server could be added


## Organization of the code

2 packages:
	- wordsplit takes an io.Reader and returns an iterator that returns 3 words.
	- worstat takes triplet as input and increment occurence


