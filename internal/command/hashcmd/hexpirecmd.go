package hashcmd

import (
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/hashes"
)

func HExpire(args []resp.Value) resp.Value {
	if len(args) == 2 {
		err := hashes.HExpire(args[0].Bulk, args[1].Bulk, "")
		if err != nil {
			return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
		}
	} else if len(args) == 3 {
		err := hashes.HExpire(args[0].Bulk, args[1].Bulk, args[2].Bulk)
		if err != nil {
			return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
		}
	} else {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}
