package pingcmd

import (
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func Ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: common.STRING_TYPE, Str: "PONG"}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: args[0].Bulk}
}
