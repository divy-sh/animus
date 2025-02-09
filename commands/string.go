package commands

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/divy-sh/animus/resp"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func append(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'append' command"}
	}
	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] += value
	SETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}

func decr(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}
	key := args[0].Bulk

	SETsMu.Lock()
	if val, ok := SETs[key]; ok {
		val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return resp.Value{Typ: "error", Str: "ERR cannot decrement a non integer value"}
		}
		SETs[key] = fmt.Sprint(val - 1)
	} else {
		SETs[key] = "0"
	}
	SETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}

func decrby(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	decrVal, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return resp.Value{Typ: "error", Str: "ERR invalid decrement value"}
	}

	SETsMu.Lock()
	if val, ok := SETs[key]; ok {
		val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return resp.Value{Typ: "error", Str: "ERR cannot decrement a non integer value"}
		}
		SETs[key] = fmt.Sprint(val - decrVal)
	} else {
		SETs[key] = "0"
	}
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

func getdel(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}
	key := args[0].Bulk
	SETsMu.RLock()
	value, ok := SETs[key]
	delete(SETs, key)
	SETsMu.RUnlock()
	if !ok {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}

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
