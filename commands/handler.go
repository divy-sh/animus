package commands

import (
	"github.com/divy-sh/animus/resp"
)

/* All supported commands, grouped by command group, sorted alphabetically */
var Handlers = map[string]func([]resp.Value) resp.Value{
	// ping commands
	"PING": ping,

	// string commands
	"APPEND": append,
	"DECR":   decr,
	"DECRBY": decrby,
	"GET":    get,
	"GETDEL": getdel,
	"GETEX":  getex,
	"SET":    set,

	// hash commands
	"HSET": hset,
	"HGET": hget,
}
