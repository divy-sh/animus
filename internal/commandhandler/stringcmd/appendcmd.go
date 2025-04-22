package stringcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/strings"
)

func Append(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	strings.Append(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func init() {
	commandhandler.RegisterCommand("APPEND", Append, `APPEND [KEY] [VALUE]
	Appends a value to a key and returns the new length of the string.`)
}
