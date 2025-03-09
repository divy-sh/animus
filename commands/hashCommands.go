package commands

import (
	"github.com/divy-sh/animus/essentias"
	"github.com/divy-sh/animus/resp"
)

var hashEssentias = essentias.NewHashEssentia()

func hset(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hashEssentias.HSet(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	return resp.Value{Typ: "string", Str: "OK"}
}

func hget(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	value, err := hashEssentias.HGet(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}
