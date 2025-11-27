/*
Package animus provides an in-memory database (similar to Redis) implemented in Go.

Animus supports various data types like strings, hashes, and lists, along with features
such as expiration handling (TTL, LRU), and basic commands.

Key Features:
  - **PING (String)**: PING [ARGUMENT]
    Returns PONG to test server responsiveness.
  - **COMMAND (String)**: COMMAND
    Returns metadata about all registered commands.
  - **INFO (String)**: INFO
    Returns information and statistics about the server.
  - **CONFIG (String)**: CONFIG
    command to handle server configuration
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
  - **LCS (String)**: LCS [KEY1] [KEY2] LEN
    Finds the Longest Common Subsequence between the value of two keys.
    Send the optional LEN argument to get just the length
  - **MGET (String)**: MGET key [key ...]
    Returns the values for all the keys.
    Returns nil for a non-existing key.
  - **MSET (String)**: MSET key value [key1 value1 ...]
    Sets the values for all the keys value pair.
  - **SET (String)**: SET [KEY] [VALUE]
    Sets the value of a key.
  - **SETRANGE (String)**: SETRANGE key offset value
  - **SETEX (String)**: SET [KEY] [VALUE] [EX SECONDS]
    Sets the value of a key with expiration in seconds.
  - **STRLEN (String)**: STRLEN [KEY]
    Returns the length of the string value stored at key.
  - **HSET (String)**: HSET [KEY] [FIELD] [VALUE]
    Sets a field in the hash stored at key to a value.
  - **HGET (String)**: HGET [KEY] [FIELD]
    Gets the value of a field in the hash stored at key.
  - **HEXISTS (String)**: HEXISTS [KEY] [FIELD]
    Checks if the hash and the field combination exists in the store.
  - **HEXPIRE (String)**: HEXPIRE key seconds [NX XX GT LT]
    Sets a timeout on hash key. After the timeout, the key gets deleted.
    NX - Only set timeout if the key has no previous expiry.
    XX - Only set timeout if the key has a previous expiry.
    GT - Only set timeout if the new time is greater than the existing expiry.
    LT - Only set timeout if the new time is less than the existing expiry.
  - **HDEL (String)**: HDEL [KEY] [FIELD]
    Deletes a field from the hash stored at key.
  - **HGETALL (String)**: HGETALL [KEY]
    Returns all fields and values of the hash stored at key.
  - **RPOP (String)**: RPOP [KEY] [COUNT]
    Removes and returns the last element(s) of the list stored at key.
  - **RPUSH (String)**: RPUSH [KEY] [VALUE] [VALUE ...]
    Inserts one or more elements at the end of the list stored at key.
  - **LINDEX (String)**: LINDEX [KEY] [INDEX]
    Returns the element at index INDEX in the list stored at key.
  - **LINSERT (String)**: LINSERT [KEY] [BEFORE|AFTER] [PIVOT] [VALUE]
    Inserts VALUE in the list stored at KEY either before or after the PIVOT element.
  - **LMOVE (String)**: LMOVE [SOURCE] [DESTINATION] [LEFT|RIGHT]
    Removes an element from the source list and pushes it to the destination list from the specified direction.
  - **LRANGE (String)**: LRANGE [KEY] [START] [END]
    Returns the specified elements of the list stored at key.
  - **LLEN (String)**: LLEN [KEY]
    Returns the length of the list stored at key.
  - **LPOP (String)**: LPOP [KEY] [COUNT]
    Removes and returns the first element(s) of the list stored at key.
  - **LPUSH (String)**: LPUSH [KEY] [VALUE] [VALUE ...]
    Inserts one or more elements at the beginning of the list stored at key.
  - **SADD (String)**: SADD [KEY] [MEMBER] [MEMBER ...]
    Adds one or more members to the set stored at key.
  - **SCARD (String)**: SCARD [KEY]
    Returns the number of members in the set stored at key.
  - **SDIFF (String)**: SDIFF [KEY] [KEY ...]
    Returns the members of the set resulting from the difference between the first set and all the successive sets.
  - **SISMEMBER (String)**: SISMEMBER [KEY] [MEMBER]
    Returns if member is a member of the set stored at key.
  - **HELP (Help)**: HELP [COMMAND]
    Provides details on how to use a command and what the command actually does.
  - **COPY (String)**: COPY [key1] [key2]
    Copies value(s) of key1 into key2.
    If key2 doesn't exist, creates key2 and copies the value of key1 into key2.
  - **DEL (String)**: DEL key1 [keys...]
    Deletes all the keys passed as argument. Ignores the keys in the argument that don't exist.
  - **EXISTS (String)**: EXISTS key1 [keys...]
    Returns an integer denoting how many of the passed keys exist in the cache.
  - **EXPIRE (String)**: EXPIRE key seconds [NX XX GT LT]
    Sets a timeout on key. After the timeout, the key gets deleted.
    NX - Only set timeout if the key has no previous expiry.
    XX - Only set timeout if the key has a previous expiry.
    GT - Only set timeout if the new time is greater than the existing expiry.
    LT - Only set timeout if the new time is less than the existing expiry.
  - **EXPIREAT (String)**: EXPIREAT key unix-time-seconds [NX XX GT LT]
    Sets the timeout of a key to the unix time stamp in seconds. After the timeout, the key gets deleted.
    NX - Only set timeout if the key has no previous expiry.
    XX - Only set timeout if the key has a previous expiry.
    GT - Only set timeout if the new time is greater than the existing expiry.
    LT - Only set timeout if the new time is less than the existing expiry.
  - **EXPIRETIME (String)**: EXPIRETIME key
    Returns the expire time of a key in unix epoch seconds.
    -1 If the key doesn't have an expiry set
    -2 If the key doesn't exist
  - **KEYS (String)**: KEYS
    Returns the keys that exist in the store.

Roadmap:
  - Advanced data structures (Sets, Sorted Sets)
  - Master-Slave replication
  - Pub/Sub for messaging
  - Performance optimizations
  - Clustering and sharding
*/
package main
