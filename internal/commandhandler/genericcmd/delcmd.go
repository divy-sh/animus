package genericcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/generics"
)

func Del(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	keys := make([]string, len(args))
	for i := 0; i < len(args); i += 2 {
		keys[i] = args[i].Bulk
	}
	generics.Delete(&keys)
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}

func init() {
	commandhandler.RegisterCommand("DEL", Del, `DEL key1 [keys...]
	Deletes all the keys passes as argument.
	If a key doesn't exist, it is ignored.`)
}
