# animus
Animus is an in-memory database (like redis) written in go. 
I am actively working on the initial stages of this project.

# Working

- PING

- String Type - APPEND, DECR, DECRBY, GET, GETDEL, GETEX, GETRANGE, SET

- Hash Type - HGET, HSET

- List Type - RPush, RPop

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
