package commands

import (
	"fmt"
	"strings"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
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

// Help command: Displays documentation for all or a specific command.
func Help(args []resp.Value) resp.Value {
	if len(args) == 0 {
		var docs []string
		for cmd, handler := range Handlers {
			docs = append(docs, fmt.Sprintf("%s - %s", cmd, handler.Doc))
		}
		return resp.Value{Typ: common.BULK_TYPE, Bulk: strings.Join(docs, "\n")}
	}

	cmd := strings.ToUpper(args[0].Bulk)
	if handler, exists := Handlers[cmd]; exists {
		return resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprintf("%s - %s", cmd, handler.Doc)}
	}
	return resp.Value{Typ: common.ERROR_TYPE, Str: "Unknown command: " + cmd}
}

// Initialize commands with their documentation
func init() {
	// Connection
	RegisterCommand("PING", ping, `PING [ARGUMENT]
	Returns PONG to test server responsiveness.`)

	// Strings
	RegisterCommand("APPEND", appendCmd, `APPEND [KEY] [VALUE]
	Appends a value to a key and returns the new length of the string.`)
	RegisterCommand("DECR", decr, `DECR [KEY]
	Decrements the integer value of a key by one.`)
	RegisterCommand("DECRBY", decrby, `DECRBY [KEY] [DECREMENT]
	Decrements the integer value of a key by the given amount.`)
	RegisterCommand("GET", get, `GET [KEY]
	Gets the value of a key.`)
	RegisterCommand("GETDEL", getdel, `GETDEL [KEY]
	Gets the value of a key and deletes it.`)
	RegisterCommand("GETEX", getex, `GETEX [KEY] [EXPIRATION]
	Gets the value of a key and sets an expiration.`)
	RegisterCommand("GETRANGE", getrange, `GETRANGE [KEY] [START] [END]
	Gets a substring of the string stored at a key.`)
	RegisterCommand("GETSET", getset, `GETSET [KEY] [VALUE]
	Gets the previous key value and then sets it to the passed value.`)
	RegisterCommand("INCR", incr, `INCR [KEY]
	Increments the integer value of a key by one.`)
	RegisterCommand("INCRBY", incrby, `INCRBY [KEY] [INCREMENT]
	Increments the integer value of a key by the given amount.`)
	RegisterCommand("INCRBYFLOAT", incrbyfloat, `INCRBYFLOAT [KEY] [INCREMENT]
	Increments the float value of a key by the given amount.`)
	RegisterCommand("LCS", lcs, `LCS [KEY1] [KEY2] LEN
	Finds the Longest Common Subsequence between the value of two keys.
	Send the optional LEN argument to get just the length`)
	RegisterCommand("MGET", mget, `MGET key [key ...]
	Returns the values for all the keys.
	Returns nil for a non-existing key.`)
	RegisterCommand("MSET", mset, `MSET key value [key1 value1 ...]
	Sets the values for all the keys value pair.`)
	RegisterCommand("SET", set, `SET [KEY] [VALUE] [EX SECONDS|PX MILLISECONDS|KEEPTTL]
	Sets the value of a key with optional expiration.`)

	// Hashes
	RegisterCommand("HSET", hset, `HSET [KEY] [FIELD] [VALUE]
	Sets a field in the hash stored at key to a value.`)
	RegisterCommand("HGET", hget, `HGET [KEY] [FIELD]
	Gets the value of a field in the hash stored at key.`)

	// Lists
	RegisterCommand("RPOP", rpop, `RPOP [KEY] [COUNT]
	Removes and returns the last element(s) of the list stored at key.`)
	RegisterCommand("RPUSH", rpush, `RPUSH [KEY] [VALUE] [VALUE ...]
	Inserts one or more elements at the end of the list stored at key.`)

	// Help
	RegisterCommand("HELP", Help, `HELP [COMMAND]
	Shows documentation for available commands.`)

	// Generics
	RegisterCommand("COPY", copyVal, `COPY [key1] [key2]
	Copies value(s) of key1 into key2.
	If key2 doesn't exist, creates key2 and sets its value.`)
	RegisterCommand("DEL", del, `DEL key1 [keys...]
	Deletes all the keys passes as argument.
	If a key doesn't exist, it is ignored.`)
	RegisterCommand("EXISTS", exists, `EXISTS key1 [keys...]
	Returns an integer denoting how many of the passed keys exist in the cache.`)
	RegisterCommand("EXPIRE", expire, `EXPIRE key seconds [NX XX GT LT]
	Sets a timeout on key. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`)
}
