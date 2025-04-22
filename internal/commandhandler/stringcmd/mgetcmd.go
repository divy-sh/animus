package stringcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/strings"
)

func MGet(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	keys := []string{}
	for _, arg := range args {
		keys = append(keys, arg.Bulk)
	}
	values := strings.MGet(&keys)
	response := make([]resp.Value, len(keys))
	for i, val := range *values {
		if val == "" {
			response[i] = resp.Value{Typ: common.NULL_TYPE}
		} else {
			response[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
		}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: response}
}

func init() {
	commandhandler.RegisterCommand("MGET", MGet, `MGET key [key ...]
	Returns the values for all the keys.
	Returns nil for a non-existing key.`)
}
