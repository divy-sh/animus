package genericcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/generics"
)

func CopyVal(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	val, err := generics.Copy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: val}
}

func init() {
	commandhandler.RegisterCommand("COPY", CopyVal, `COPY [key1] [key2]
	Copies value(s) of key1 into key2.
	If key2 doesn't exist, creates key2 and sets its value.`)
}
