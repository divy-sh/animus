package commands

import (
	"fmt"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/generics"
	"github.com/divy-sh/animus/resp"
)

func copy(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	val, err := generics.Copy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(val)}
}

func del(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	keys := make([]string, len(args))
	for i := 0; i < len(args); i += 2 {
		keys[i] = args[i].Bulk
	}
	generics.Delete(keys)
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}
