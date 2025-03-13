package commands

import (
	"github.com/divy-sh/animus/resp"
)

type Command struct {
	Func          func([]resp.Value) resp.Value
	Documentation string
}

/* All supported commands, grouped by command group, sorted alphabetically */
var Handlers = map[string]Command{

	// ping commands
	"PING": {
		Func:          ping,
		Documentation: "PING - No arguments. Returns PONG to test server responsiveness.",
	},

	// string commands
	"APPEND": {
		Func:          append,
		Documentation: "APPEND key value - Appends a value to a key and returns the new length of the string.",
	},
	"DECR": {
		Func:          decr,
		Documentation: "DECR key - Decrements the integer value of a key by one.",
	},
	"DECRBY": {
		Func:          decrby,
		Documentation: "DECRBY key decrement - Decrements the integer value of a key by the given amount.",
	},
	"GET": {
		Func:          get,
		Documentation: "GET key - Gets the value of a key.",
	},
	"GETDEL": {
		Func:          getdel,
		Documentation: "GETDEL key - Gets the value of a key and deletes it.",
	},
	"GETEX": {
		Func:          getex,
		Documentation: "GETEX key [expiration] - Gets the value of a key and sets an expiration.",
	},
	"GETRANGE": {
		Func:          getrange,
		Documentation: "GETRANGE key start end - Gets a substring of the string stored at a key.",
	},
	"GETSET": {
		Func:          getset,
		Documentation: "GETSET key value - Gets the previous key value and then sets it to the passed value",
	},
	"INCR": {
		Func:          incr,
		Documentation: "INCR key - Increments the integer value of a key by one.",
	},
	"INCRBY": {
		Func:          incrby,
		Documentation: "INCRBY key increment - Increments the integer value of a key by the given amount.",
	},
	"INCRBYFLOAT": {
		Func:          incrbyfloat,
		Documentation: "INCRBYFLOAT key increment - Increments the float value of a key by the given amount.",
	},
	"SET": {
		Func:          set,
		Documentation: "SET key value [EX seconds|PX milliseconds|KEEPTTL] - Sets the value of a key with optional expiration.",
	},

	// hash commands
	"HSET": {
		Func:          hset,
		Documentation: "HSET key field value - Sets a field in the hash stored at key to a value.",
	},
	"HGET": {
		Func:          hget,
		Documentation: "HGET key field - Gets the value of a field in the hash stored at key.",
	},

	// list commands
	"RPOP": {
		Func:          rpop,
		Documentation: "RPOP key [count] - Removes and returns the last element(s) of the list stored at key.",
	},
	"RPUSH": {
		Func:          rpush,
		Documentation: "RPUSH key value [value ...] - Inserts one or more elements at the end of the list stored at key.",
	},
}
