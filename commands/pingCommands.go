package commands

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: common.STRING_TYPE, Str: "PONG"}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: args[0].Bulk}
}
