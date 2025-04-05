package commands

import (
	"fmt"

	"github.com/divy-sh/animus/generics"
	"github.com/divy-sh/animus/resp"
)

func copy(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'copy' command"}
	}
	val, err := generics.Copy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: "error", Str: err.Error()}
	}
	return resp.Value{Typ: "bulk", Bulk: fmt.Sprint(val)}
}

func del(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'del' command"}
	}
	keys := make([]string, len(args))
	for i := 0; i < len(args); i += 2 {
		keys[i] = args[i].Bulk
	}
	generics.Delete(keys)
	return resp.Value{Typ: "bulk", Bulk: "OK"}
}
