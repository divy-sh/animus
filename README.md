# animus
[![Build and Test](https://github.com/divy-sh/animus/actions/workflows/go.yml/badge.svg)](https://github.com/divy-sh/animus/actions/workflows/go.yml)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/divy-sh/animus)
[![License](https://img.shields.io/badge/License-GNU20GPL-blue?style=flat-square)](https://raw.githubusercontent.com/divy-sh/animus/master/LICENSE)

Animus is an in-memory database (like redis) written in go. 

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
LCS - LCS [KEY1] [KEY2] [LEN] [IDX] [MINMATCHLEN MIN-MATCH-LEN] [WITHMATCHLEN]
```
Finds the Longest Common Subsequence between the value of two keys.
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

- Expiry mechanism (LRU, TTL)

- key separation between 

# Roadmap

- add configuration support so that users can set expirations, max memory, etc.

- Add pools for lock keys to curb high memory usage, also add key based locking to essentias

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