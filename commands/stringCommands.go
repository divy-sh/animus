package commands

import (
	"github.com/divy-sh/animus/essentias"
	"github.com/divy-sh/animus/resp"
)

var stringEssentia = essentias.NewStringEssentia()

func appendCmd(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'append' command"}
	}
	stringEssentia.Append(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: "string", Str: "OK"}
}

func decr(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'decr' command"}
	}
	err := stringEssentia.Decr(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	return resp.Value{Typ: "string", Str: "OK"}
}

func decrby(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'decrby' command"}
	}

	err := stringEssentia.DecrBy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}

	return resp.Value{Typ: "string", Str: "OK"}
}

func get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}
	value, err := stringEssentia.Get(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}

func getdel(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'getdel' command"}
	}
	value, err := stringEssentia.GetDel(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}

func getex(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'getex' command"}
	}
	value, err := stringEssentia.GetEx(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}

func getrange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'getrange' command"}
	}
	value, err := stringEssentia.GetRange(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	if err != nil {
		return resp.Value{Typ: "null"}
	}
	return resp.Value{Typ: "bulk", Bulk: value}
}

func getset(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'getset' command"}
	}

	val, err := stringEssentia.GetSet(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	return resp.Value{Typ: "bulk", Bulk: val}
}

func incr(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'incr' command"}
	}
	err := stringEssentia.Incr(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	return resp.Value{Typ: "string", Str: "OK"}
}

func incrby(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'incrby' command"}
	}

	err := stringEssentia.IncrBy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}

	return resp.Value{Typ: "string", Str: "OK"}
}

func incrbyfloat(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'incrby' command"}
	}

	err := stringEssentia.IncrByFloat(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}

	return resp.Value{Typ: "string", Str: "OK"}
}

func lcs(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'lcs' command"}
	}
	commands := []string{}
	for _, arg := range args[2:] {
		commands = append(commands, arg.Bulk)
	}
	val, err := stringEssentia.Lcs(args[0].Bulk, args[1].Bulk, commands)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	return resp.Value{Typ: "bulk", Bulk: val}
}

func mget(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'mget' command"}
	}
	keys := []string{}
	for _, arg := range args {
		keys = append(keys, arg.Bulk)
	}
	values := stringEssentia.MGet(&keys)
	response := make([]resp.Value, len(keys))
	for i, val := range *values {
		if val == "" {
			response[i] = resp.Value{Typ: "null"}
		} else {
			response[i] = resp.Value{Typ: "bulk", Bulk: val}
		}
	}
	return resp.Value{Typ: "array", Array: response}
}

func mset(args []resp.Value) resp.Value {
	if len(args) < 2 || len(args)&1 == 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'mget' command"}
	}
	kvPairs := map[string]string{}
	for i := 0; i < len(args); i += 2 {
		kvPairs[args[i].Bulk] = args[i+1].Bulk
	}
	stringEssentia.MSet(&kvPairs)
	return resp.Value{Typ: "string", Str: "OK"}
}

func set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	stringEssentia.Set(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: "string", Str: "OK"}
}
