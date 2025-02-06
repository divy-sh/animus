package commands

import (
	"sync"

	"github.com/divy-sh/animus/resp"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}

func get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}
	key := args[0].Bulk
	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()
	if !ok {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}
