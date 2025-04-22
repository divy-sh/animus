package stringcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/strings"
)

func MSet(args []resp.Value) resp.Value {
	if len(args) < 2 || len(args)&1 == 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	kvPairs := map[string]string{}
	for i := 0; i < len(args); i += 2 {
		kvPairs[args[i].Bulk] = args[i+1].Bulk
	}
	strings.MSet(&kvPairs)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func init() {
	commandhandler.RegisterCommand("MSET", MSet, `MSET key value [key1 value1 ...]
	Sets the values for all the keys value pair.`)
}
