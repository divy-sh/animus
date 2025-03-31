# animus
[![Build and Test](https://github.com/divy-sh/animus/actions/workflows/go.yml/badge.svg)](https://github.com/divy-sh/animus/actions/workflows/go.yml)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/divy-sh/animus)
[![License](https://img.shields.io/badge/License-GNU30GPL-blue?style=flat-square)](https://raw.githubusercontent.com/divy-sh/animus/master/LICENSE)

Animus is an in-memory database (like Redis) written in Go. It offers a lightweight, high-performance storage solution with support for various commands and data structures.

To see all the supported commands, go to doc.go file.

# Features

- Expiration Mechanism: Support for TTL (Time-to-Live) and LRU (Least Recently Used) expiry policies.
- Data Types: Support for strings, lists, hashes, and advanced data structures.
- Key Management: Automatic key expiration, deletion, and manipulation.
- Concurrency: Optimized for high concurrency and scalability.

# Roadmap

The following features are planned for future releases:

- Configuration Support: Allow users to configure expirations, max memory, etc.
- Key Locking: Implement pools for key locks to handle high memory usage and key-based locking for essentials.
- Advanced Data Structures: Expand support for additional data structures like sets, sorted sets, and more.
- Replication: Master-slave replication to enable high availability and fault tolerance.
- Pub/Sub: Real-time messaging with Publish/Subscribe functionality.
- Performance Optimizations: Optimizations to the event loop for enhanced performance.
- Clustering & Sharding: Scalable architecture with clustering and sharding.


# Test Coverage
- Test coverage can be verified by generating a cover.html file.

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

### v0.0.3

- Add these functions
    - GETSET, INCR, INCRBY, INCRBYFLOAT

- Added key based locking instead of locking the store globally, leading to huge increase in concurrency.

### v0.1.0

- Finalize how to handle the store, and expose the store locks so that they can be used by other commands as required.

- Add many other commands, details can be found in doc.go file.
# License

Animus is licensed under the GNU General Public License v3.0 (GPL-3.0). See the LICENSE file for more information.