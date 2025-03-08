package commands

import (
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types"
)

var listTypes = types.NewListType()

func rpop(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'rpop' command"}
	}
	var values []string
	var err error
	if len(args) == 1 {
		values, err = listTypes.RPop(args[0].Bulk, "1")
	} else {
		values, err = listTypes.RPop(args[0].Bulk, args[1].Bulk)
	}
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	respArr := make([]resp.Value, len(values))
	for i, val := range values {
		respArr[i] = resp.Value{Typ: "bulk", Bulk: val}
	}
	return resp.Value{Typ: "array", Array: respArr}
}

func rpush(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'rpush' command"}
	}
	values := make([]string, len(args)-1)
	for i, val := range args[1:] {
		values[i] = val.Bulk
	}
	listTypes.RPush(args[0].Bulk, &values)
	return resp.Value{Typ: "string", Str: "OK"}
}
