package hashcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/hashes"
)

func HSet(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	hashes.HSet(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func init() {
	commandhandler.RegisterCommand("HSET", HSet, `HSET [KEY] [FIELD] [VALUE]
	Sets a field in the hash stored at key to a value.`)
}
