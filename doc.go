/*
Package animus provides an in-memory database (similar to Redis) implemented in Go.

Animus supports various data types like strings, hashes, and lists, along with features
such as expiration handling (TTL, LRU), and basic commands.

Key Features:
  - **String Commands**: APPEND, DECR, DECRBY, GET, GETDEL, GETEX, GETRANGE, SET
  - **Hash Commands**: HGET, HSET
  - **List Commands**: RPUSH, RPOP
  - **Expiry Mechanism**: Supports TTL (Time-To-Live) and LRU (Least Recently Used) strategies for data expiration.

Data Types and Commands:
  - **Strings**: Basic key-value pairs supporting commands like SET, GET, APPEND, DECR, and DECRBY.
  - **Hashes**: Maps between string fields and values, supporting HSET and HGET commands.
  - **Lists**: Ordered collections of strings, supporting RPUSH and RPOP commands.

Expiry Mechanism:
  - **Time-To-Live (TTL)**: Set an expiration time for keys, after which they are automatically deleted.
  - **Least Recently Used (LRU)**: Automatically evicts the least recently accessed keys when memory is full.

Test Coverage:
  - To generate a test coverage report:
    1. Run `go test -v -coverprofile cover.out ./...` to execute tests and generate a coverage profile.
    2. Run `go tool cover -html cover.out -o cover.html` to generate an HTML coverage report.

Roadmap:
  - Advanced data structures (Sets, Sorted Sets)
  - Master-Slave replication
  - Pub/Sub for messaging
  - Performance optimizations
  - Clustering and sharding
*/
package main
