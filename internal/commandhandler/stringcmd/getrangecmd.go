package stringcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/strings"
)

func GetRange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.GetRange(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	if err != nil {
		return resp.Value{Typ: common.NULL_TYPE}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func init() {
	commandhandler.RegisterCommand("GETRANGE", GetRange, `GETRANGE [KEY] [START] [END]
	Gets a substring of the string stored at a key.`)
}
