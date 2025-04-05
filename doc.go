/*
Package animus provides an in-memory database (similar to Redis) implemented in Go.

Animus supports various data types like strings, hashes, and lists, along with features
such as expiration handling (TTL, LRU), and basic commands.

Key Features:
  - **PING (String)**: PING [ARGUMENT]
	Returns PONG to test server responsiveness.
  - **APPEND (String)**: APPEND [KEY] [VALUE]
	Appends a value to a key and returns the new length of the string.
  - **DECR (String)**: DECR [KEY]
	Decrements the integer value of a key by one.
  - **DECRBY (String)**: DECRBY [KEY] [DECREMENT]
	Decrements the integer value of a key by the given amount.
  - **GET (String)**: GET [KEY]
	Gets the value of a key.
  - **GETDEL (String)**: GETDEL [KEY]
	Gets the value of a key and deletes it.
  - **GETEX (String)**: GETEX [KEY] [EXPIRATION]
	Gets the value of a key and sets an expiration.
  - **GETRANGE (String)**: GETRANGE [KEY] [START] [END]
	Gets a substring of the string stored at a key.
  - **GETSET (String)**: GETSET [KEY] [VALUE]
	Gets the previous key value and then sets it to the passed value.
  - **INCR (String)**: INCR [KEY]
	Increments the integer value of a key by one.
  - **INCRBY (String)**: INCRBY [KEY] [INCREMENT]
	Increments the integer value of a key by the given amount.
  - **INCRBYFLOAT (String)**: INCRBYFLOAT [KEY] [INCREMENT]
	Increments the float value of a key by the given amount.
  - **LCS (List)**: LCS [KEY1] [KEY2] LEN
	Finds the Longest Common Subsequence between the value of two keys.
	Send the optional LEN argument to get just the length
  - **MGET (String)**: MGET key [key ...]
	Returns the values for all the keys.
	Returns nil for a non-existing key.
  - **MSET (String)**: MSET key value [key1 value1 ...]
	Sets the values for all the keys value pair.
  - **SET (String)**: SET [KEY] [VALUE] [EX SECONDS|PX MILLISECONDS|KEEPTTL]
	Sets the value of a key with optional expiration.
  - **HSET (Hash)**: HSET [KEY] [FIELD] [VALUE]
	Sets a field in the hash stored at key to a value.
  - **HGET (Hash)**: HGET [KEY] [FIELD]
	Gets the value of a field in the hash stored at key.
  - **RPOP (List)**: RPOP [KEY] [COUNT]
	Removes and returns the last element(s) of the list stored at key.
  - **RPUSH (List)**: RPUSH [KEY] [VALUE] [VALUE ...]
	Inserts one or more elements at the end of the list stored at key.
  - **HELP (Help)**: HELP [COMMAND]
	Shows documentation for available commands.
  - **COPY (String)**: COPY [key1] [key2]
	Copies value(s) of key1 into key2.
	If key2 doesn't exist, creates key2 and sets its value.
  - **DEL (String)**: DEL key1 [keys...]
	Deletes all the keys passes as argument.
	If a key doesn't exist, it is ignored.

Roadmap:
  - Advanced data structures (Sets, Sorted Sets)
  - Master-Slave replication
  - Pub/Sub for messaging
  - Performance optimizations
  - Clustering and sharding
*/
package main
