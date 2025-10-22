package command

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/generics"
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

func Exists(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	keys := make([]string, len(args))
	for i := 0; i < len(args); i += 2 {
		keys[i] = args[i].Bulk
	}
	validKeyCount := generics.Exists(&keys)
	return resp.Value{Typ: common.INTEGER_TYPE, Num: validKeyCount}
}

func ExpireAt(args []resp.Value) resp.Value {
	if len(args) == 2 {
		err := generics.ExpireAt(args[0].Bulk, args[1].Bulk, "")
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

func ExpireTime(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	val, err := generics.ExpireTime(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: val}
}

func Keys(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	values, err := generics.Keys(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	response := make([]resp.Value, len(*values))
	for i, val := range *values {
		response[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: response}
}
