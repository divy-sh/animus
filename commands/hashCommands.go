package commands

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/essentias"
	"github.com/divy-sh/animus/resp"
)

func hset(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: "ERR wrong number of arguments for 'hset' command"}
	}

	essentias.HSet(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func hget(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: "ERR wrong number of arguments for 'hget' command"}
	}

	value, err := essentias.HGet(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}
