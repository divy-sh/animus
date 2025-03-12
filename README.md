# animus
[![Build and Test](https://github.com/divy-sh/animus/actions/workflows/go.yml/badge.svg)](https://github.com/divy-sh/animus/actions/workflows/go.yml)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/divy-sh/animus)
[![License](https://img.shields.io/badge/License-GNU20GPL-blue?style=flat-square)](https://raw.githubusercontent.com/divy-sh/animus/master/LICENSE)

Animus is an in-memory database (like redis) written in go. 
I am actively working on the initial stages of this project.

# Working

- PING

- String Essentia - APPEND, DECR, DECRBY, GET, GETDEL, GETEX, GETRANGE, GETSET, SET, INCR, INCRBY

- Hash Essentia - HGET, HSET

- List Essentia - RPush, RPop

- Expiry mechanism (LRU, TTL)

# Roadmap

- Support for advanced data structures (Lists, Sets, Hashes)

- Master-Slave replication

- Pub/Sub for real-time messaging

- Optimize event loop for performance

- Clustering and sharding for scalability

# Test Coverage
- Test coverage can be verified by generating cover.html file

- To generate the cover.html file, run these commands:
```bash
go test -v -coverprofile cover.out ./...
```
```bash
go tool cover -html cover.out -o cover.html
```

## Release Notes

### v0.0.1

- First package release. Can be found here - https://pkg.go.dev/github.com/divy-sh/animus@v0.0.1

### v0.0.2
- Refactor packages to have better names so that there are no potential naming conflicts - https://pkg.go.dev/github.com/divy-sh/animus