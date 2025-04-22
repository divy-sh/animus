package command

import (
	"github.com/divy-sh/animus/internal/command/genericcmd"
	"github.com/divy-sh/animus/internal/command/hashcmd"
	"github.com/divy-sh/animus/internal/command/listcmd"
	"github.com/divy-sh/animus/internal/command/pingcmd"
	"github.com/divy-sh/animus/internal/command/stringcmd"
	"github.com/divy-sh/animus/internal/resp"
)

// Command represents a command with an associated function and documentation.
type Command struct {
	Func func([]resp.Value) resp.Value
	Doc  string
}

// Handlers maps command names to their implementations.
var Handlers = map[string]Command{}

// RegisterCommand registers a command function with its documentation.
func RegisterCommand(name string, fn func([]resp.Value) resp.Value, doc string) {
	Handlers[name] = Command{Func: fn, Doc: doc}
}

// Initialize commands with their documentation
func init() {
	// Connection
	RegisterCommand("PING", pingcmd.Ping, `PING [ARGUMENT]
	Returns PONG to test server responsiveness.`)

	// Strings
	RegisterCommand("APPEND", stringcmd.Append, `APPEND [KEY] [VALUE]
	Appends a value to a key and returns the new length of the string.`)
	RegisterCommand("DECR", stringcmd.Decr, `DECR [KEY]
	Decrements the integer value of a key by one.`)
	RegisterCommand("DECRBY", stringcmd.DecrBy, `DECRBY [KEY] [DECREMENT]
	Decrements the integer value of a key by the given amount.`)
	RegisterCommand("GET", stringcmd.Get, `GET [KEY]
	Gets the value of a key.`)
	RegisterCommand("GETDEL", stringcmd.GetDel, `GETDEL [KEY]
	Gets the value of a key and deletes it.`)
	RegisterCommand("GETEX", stringcmd.GetEx, `GETEX [KEY] [EXPIRATION]
	Gets the value of a key and sets an expiration.`)
	RegisterCommand("GETRANGE", stringcmd.GetRange, `GETRANGE [KEY] [START] [END]
	Gets a substring of the string stored at a key.`)
	RegisterCommand("GETSET", stringcmd.GetSet, `GETSET [KEY] [VALUE]
	Gets the previous key value and then sets it to the passed value.`)
	RegisterCommand("INCR", stringcmd.Incr, `INCR [KEY]
	Increments the integer value of a key by one.`)
	RegisterCommand("INCRBY", stringcmd.IncrBy, `INCRBY [KEY] [INCREMENT]
	Increments the integer value of a key by the given amount.`)
	RegisterCommand("INCRBYFLOAT", stringcmd.IncrByFloat, `INCRBYFLOAT [KEY] [INCREMENT]
	Increments the float value of a key by the given amount.`)
	RegisterCommand("LCS", stringcmd.LCS, `LCS [KEY1] [KEY2] LEN
	Finds the Longest Common Subsequence between the value of two keys.
	Send the optional LEN argument to get just the length`)
	RegisterCommand("MGET", stringcmd.MGet, `MGET key [key ...]
	Returns the values for all the keys.
	Returns nil for a non-existing key.`)
	RegisterCommand("MSET", stringcmd.MSet, `MSET key value [key1 value1 ...]
	Sets the values for all the keys value pair.`)
	RegisterCommand("SET", stringcmd.Set, `SET [KEY] [VALUE] [EX SECONDS|PX MILLISECONDS|KEEPTTL]
	Sets the value of a key with optional expiration.`)

	// Hashes
	RegisterCommand("HSET", hashcmd.HSet, `HSET [KEY] [FIELD] [VALUE]
	Sets a field in the hash stored at key to a value.`)
	RegisterCommand("HGET", hashcmd.HGet, `HGET [KEY] [FIELD]
	Gets the value of a field in the hash stored at key.`)

	// Lists
	RegisterCommand("RPOP", listcmd.RPop, `RPOP [KEY] [COUNT]
	Removes and returns the last element(s) of the list stored at key.`)
	RegisterCommand("RPUSH", listcmd.RPush, `RPUSH [KEY] [VALUE] [VALUE ...]
	Inserts one or more elements at the end of the list stored at key.`)

	// Help
	RegisterCommand("HELP", Help, `HELP [COMMAND]
	Shows documentation for available commands.`)

	// Generics
	RegisterCommand("COPY", genericcmd.CopyVal, `COPY [key1] [key2]
	Copies value(s) of key1 into key2.
	If key2 doesn't exist, creates key2 and sets its value.`)
	RegisterCommand("DEL", genericcmd.Del, `DEL key1 [keys...]
	Deletes all the keys passes as argument.
	If a key doesn't exist, it is ignored.`)
	RegisterCommand("EXISTS", genericcmd.Exists, `EXISTS key1 [keys...]
	Returns an integer denoting how many of the passed keys exist in the cache.`)
	RegisterCommand("EXPIRE", genericcmd.Expire, `EXPIRE key seconds [NX XX GT LT]
	Sets a timeout on key. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`)
}
