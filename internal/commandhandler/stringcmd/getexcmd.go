package stringcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/strings"
)

func GetEx(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.GetEx(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.NULL_TYPE}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func init() {
	commandhandler.RegisterCommand("GETEX", GetEx, `GETEX [KEY] [EXPIRATION]
	Gets the value of a key and sets an expiration.`)
}
