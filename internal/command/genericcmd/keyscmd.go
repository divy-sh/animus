package genericcmd

import (
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/generics"
)

func Keys(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	values, err := generics.Keys(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	response := make([]resp.Value, len(*values))
	for i, val := range *values {
		response[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: response}
}
