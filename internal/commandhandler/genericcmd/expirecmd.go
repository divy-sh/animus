package genericcmd

import (
	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
	"github.com/divy-sh/animus/internal/types/generics"
)

func Expire(args []resp.Value) resp.Value {
	if len(args) == 2 {
		err := generics.Expire(args[0].Bulk, args[1].Bulk, "")
		if err != nil {
			return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
		}
	} else if len(args) == 3 {
		err := generics.Expire(args[0].Bulk, args[1].Bulk, args[2].Bulk)
		if err != nil {
			return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
		}
	} else {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}

func init() {
	commandhandler.RegisterCommand("EXPIRE", Expire, `EXPIRE key seconds [NX XX GT LT]
	Sets a timeout on key. After the timeout, the key gets deleted.
	NX - Only set timeout if the key has no previous expiry.
	XX - Only set timeout if the key has a previous expiry.
	GT - Only set timeout if the new time is greater than the existing expiry.
	LT - Only set timeout if the new time is less than the existing expiry.`)
}
