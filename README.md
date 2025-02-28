# animus
Animus is an in-memory database (like redis) written in go. 
I am actively working on the initial stages of this project.

# Working

- PING

- APPEND, DECR, DECRBY, GET, GETDEL, GETEX, GETRANGE, SET

- HGET, HSET

# Roadmap

- Expiry mechanism (EXPIRE, TTL), (implemented for string type)

- Support for advanced data structures (Lists, Sets, Hashes)

- Master-Slave replication

- Pub/Sub for real-time messaging

- Optimize event loop for performance

- Clustering and sharding for scalability

# Test Coverage
- 89.4% of statements covered with tests

- Test coverage can be verified by generating cover.html file

- To generate the cover.html file, run these commands:
```bash
go test -v -coverprofile cover.out ./...
```
```bash
go tool cover -html cover.out -o cover.html
```
