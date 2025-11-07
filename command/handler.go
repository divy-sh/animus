package command

import (
	"github.com/divy-sh/animus/resp"
)

// Command represents a command with an associated function and documentation.
type Command struct {
	Func     func([]resp.Value) resp.Value
	Doc      string
	Arity    int
	Flags    []string
	FirstKey int
	LastKey  int
	Step     int
}

// Handlers maps command names to their implementations.
var Handlers = map[string]Command{}

// RegisterCommand registers a command function with its documentation.
func RegisterCommand(name string, fn func([]resp.Value) resp.Value, doc string, flags []string, arity, firstKey, lastKey, step int) {
	Handlers[name] = Command{Func: fn, Doc: doc, Flags: flags, Arity: arity, FirstKey: firstKey, LastKey: lastKey, Step: step}
}

// Initialize commands with their documentation.
// Arguments apart from name, function and documentation are for metadata purposes. They may not be completely correct.
func init() {
	// Connection
	RegisterCommand("PING", Ping, `PING [ARGUMENT]
	Returns PONG to test server responsiveness.`, []string{"readonly", "fast"}, -1, 0, 0, 0)
	RegisterCommand("COMMAND", CommandCmd, `COMMAND
	Returns metadata about all registered commands.`, []string{"readonly", "fast"}, 0, 0, 0, 0)
	RegisterCommand("INFO", Info, `INFO
    Returns information and statistics about the server.`, []string{"readonly", "fast"}, 0, 0, 0, 0)
	RegisterCommand("CONFIG", ConfigCmd, `CONFIG
	command to handle server configuration`, []string{"readonly", "fast"}, -1, 0, 0, 0)

	// Strings
	RegisterCommand("APPEND", Append, `APPEND [KEY] [VALUE]
	Appends a value to a key and returns the new length of the string.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("DECR", Decr, `DECR [KEY]
	Decrements the integer value of a key by one.`, []string{}, 2, 0, 0, 0)
	RegisterCommand("DECRBY", DecrBy, `DECRBY [KEY] [DECREMENT]
	Decrements the integer value of a key by the given amount.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("GET", Get, `GET [KEY]
	Gets the value of a key.`, []string{"readonly", "fast"}, 2, 0, 0, 0)
	RegisterCommand("GETDEL", GetDel, `GETDEL [KEY]
	Gets the value of a key and deletes it.`, []string{}, 2, 0, 0, 0)
	RegisterCommand("GETEX", GetEx, `GETEX [KEY] [EXPIRATION]
	Gets the value of a key and sets an expiration.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("GETRANGE", GetRange, `GETRANGE [KEY] [START] [END]
	Gets a substring of the string stored at a key.`, []string{"readonly", "fast"}, 4, 0, 0, 0)
	RegisterCommand("GETSET", GetSet, `GETSET [KEY] [VALUE]
	Gets the previous key value and then sets it to the passed value.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("INCR", Incr, `INCR [KEY]
	Increments the integer value of a key by one.`, []string{}, 2, 0, 0, 0)
	RegisterCommand("INCRBY", IncrBy, `INCRBY [KEY] [INCREMENT]
	Increments the integer value of a key by the given amount.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("INCRBYFLOAT", IncrByFloat, `INCRBYFLOAT [KEY] [INCREMENT]
	Increments the float value of a key by the given amount.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("LCS", LCS, `LCS [KEY1] [KEY2] LEN
	Finds the Longest Common Subsequence between the value of two keys.
	Send the optional LEN argument to get just the length`, []string{"readonly", "fast"}, 4, 0, 0, 0)
	RegisterCommand("MGET", MGet, `MGET key [key ...]
	Returns the values for all the keys.
	Returns nil for a non-existing key.`, []string{"readonly", "fast"}, -2, 0, 0, 0)
	RegisterCommand("MSET", MSet, `MSET key value [key1 value1 ...]
	Sets the values for all the keys value pair.`, []string{}, -3, 0, 0, 0)
	RegisterCommand("SET", Set, `SET [KEY] [VALUE]
	Sets the value of a key.`, []string{}, -3, 0, 0, 0)
	RegisterCommand("SETRANGE", SetRange, `SETRANGE key offset value`, []string{}, -3, 0, 0, 0)
	RegisterCommand("SETEX", SetEx, `SET [KEY] [VALUE] [EX SECONDS]
	Sets the value of a key with expiration in seconds.`, []string{}, 4, 0, 0, 0)
	RegisterCommand("STRLEN", StrLen, `STRLEN [KEY]
	Returns the length of the string value stored at key.`, []string{"readonly", "fast"}, 2, 0, 0, 0)

	// Hashes
	RegisterCommand("HSET", HSet, `HSET [KEY] [FIELD] [VALUE]
	Sets a field in the hash stored at key to a value.`, []string{}, -4, 0, 0, 0)
	RegisterCommand("HGET", HGet, `HGET [KEY] [FIELD]
	Gets the value of a field in the hash stored at key.`, []string{"readonly", "fast"}, 3, 0, 0, 0)
	RegisterCommand("HEXISTS", HGet, `HEXISTS [KEY] [FIELD]
	Checks if the hash and the field combination exists in the store.`, []string{"readonly", "fast"}, 3, 0, 0, 0)
	RegisterCommand("HEXPIRE", HExpire, `HEXPIRE key seconds [NX XX GT LT]
	Sets a timeout on hash key. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("HDEL", HDel, `HDEL [KEY] [FIELD]
	Deletes a field from the hash stored at key.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("HGETALL", HGetAll, `HGETALL [KEY]
	Returns all fields and values of the hash stored at key.`, []string{"readonly", "fast"}, 2, 0, 0, 0)

	// Lists
	RegisterCommand("RPOP", RPop, `RPOP [KEY] [COUNT]
	Removes and returns the last element(s) of the list stored at key.`, []string{}, -2, 0, 0, 0)
	RegisterCommand("RPUSH", RPush, `RPUSH [KEY] [VALUE] [VALUE ...]
	Inserts one or more elements at the end of the list stored at key.`, []string{}, -3, 0, 0, 0)

	// Sets
	RegisterCommand("SADD", Sadd, `SADD [KEY] [MEMBER] [MEMBER ...]
	Adds one or more members to the set stored at key.`, []string{}, -3, 0, 0, 0)
	RegisterCommand("Scard", Scard, `SCARD [KEY]
	Returns the number of members in the set stored at key.`, []string{"readonly", "fast"}, 2, 0, 0, 0)

	// Help
	RegisterCommand("HELP", Help, `HELP [COMMAND]
	Provides details on how to use a command and what the command actually does.`, []string{"readonly", "fast"}, -1, 0, 0, 0)

	// Generics
	RegisterCommand("COPY", CopyVal, `COPY [key1] [key2]
	Copies value(s) of key1 into key2.
	If key2 doesn't exist, creates key2 and copies the value of key1 into key2.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("DEL", Del, `DEL key1 [keys...]
	Deletes all the keys passed as argument. Ignores the keys in the argument that don't exist.`, []string{}, -2, 0, 0, 0)
	RegisterCommand("EXISTS", Exists, `EXISTS key1 [keys...]
	Returns an integer denoting how many of the passed keys exist in the cache.`, []string{"readonly", "fast"}, -2, 0, 0, 0)
	RegisterCommand("EXPIRE", Expire, `EXPIRE key seconds [NX XX GT LT]
	Sets a timeout on key. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("EXPIREAT", ExpireAt, `EXPIREAT key unix-time-seconds [NX XX GT LT]
	Sets the timeout of a key to the unix time stamp in seconds. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`, []string{}, 3, 0, 0, 0)
	RegisterCommand("EXPIRETIME", ExpireTime, `EXPIRETIME key
	Returns the expire time of a key in unix epoch seconds.
	-1 If the key doesn't have an expiry set
	-2 If the key doesn't exist`, []string{"readonly", "fast"}, 2, 0, 0, 0)
	RegisterCommand("KEYS", Keys, `KEYS
	Returns the keys that exist in the store.`, []string{"readonly", "fast"}, 1, 0, 0, 0)
}
