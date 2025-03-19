# animus
[![Build and Test](https://github.com/divy-sh/animus/actions/workflows/go.yml/badge.svg)](https://github.com/divy-sh/animus/actions/workflows/go.yml)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/divy-sh/animus)
[![License](https://img.shields.io/badge/License-GNU30GPL-blue?style=flat-square)](https://raw.githubusercontent.com/divy-sh/animus/master/LICENSE)

Animus is an in-memory database (like Redis) written in Go. It offers a lightweight, high-performance storage solution with support for various commands and data structures.

# Working Commands
```bash
GETEX - GETEX [KEY] [EXPIRATION]
```
Gets the value of a key and sets an expiration.
```bash
GETRANGE - GETRANGE [KEY] [START] [END]
```
Gets a substring of the string stored at a key.
```bash
INCR - INCR [KEY]
```
Increments the integer value of a key by one.
```bash
LCS - LCS [KEY1] [KEY2] LEN
```
Finds the Longest Common Subsequence between the value of two keys. 
```bash
MGET key [keys ...]
```
Returns the values for all the keys, nil for a non-existing key.
```bash
INCRBYFLOAT - INCRBYFLOAT [KEY] [INCREMENT]
```
Increments the float value of a key by the given amount.
```bash
RPOP - RPOP [KEY] [COUNT]
```
Removes and returns the last element(s) of the list stored at key.
```bash
DECRBY - DECRBY [KEY] [DECREMENT]
```
Decrements the integer value of a key by the given amount.
```bash
GETDEL - GETDEL [KEY]
```
Gets the value of a key and deletes it.
```bash
SET - SET [KEY] [VALUE] [EX SECONDS|PX MILLISECONDS|KEEPTTL]
```
Sets the value of a key with optional expiration.
```bash
GETSET - GETSET [KEY] [VALUE]
```
Gets the previous key value and then sets it to the passed value.
```bash
INCRBY - INCRBY [KEY] [INCREMENT]
```
Increments the integer value of a key by the given amount.
```bash
HSET - HSET [KEY] [FIELD] [VALUE]
```
Sets a field in the hash stored at key to a value.
```bash
HGET - HGET [KEY] [FIELD]
```
Gets the value of a field in the hash stored at key.
```bash
PING - PING [ARGUMENT]
```
Returns PONG to test server responsiveness.
```bash
APPEND - APPEND [KEY] [VALUE]
```
Appends a value to a key and returns the new length of the string.
```bash
DECR - DECR [KEY]
```
Decrements the integer value of a key by one.
```bash
GET - GET [KEY]
```
Gets the value of a key.
```bash
RPUSH - RPUSH [KEY] [VALUE] [VALUE ...]
```
Inserts one or more elements at the end of the list stored at key.
```bash
HELP - HELP [COMMAND]
```
Shows documentation for available commands.

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

# License

Animus is licensed under the GNU General Public License v3.0 (GPL-3.0). See the LICENSE file for more information.