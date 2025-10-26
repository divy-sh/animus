package command

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/hashes"
)

func HExists(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	val, err := hashes.HExists(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: val}
}

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

func HGet(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	value, err := hashes.HGet(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func HSet(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	hashes.HSet(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func HDel(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	err := hashes.HDel(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func HGetAll(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	result, err := hashes.HGetAll(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	values := make([]resp.Value, 0, len(result)*2)
	for field, value := range result {
		values = append(values, resp.Value{Typ: common.BULK_TYPE, Bulk: field})
		values = append(values, resp.Value{Typ: common.BULK_TYPE, Bulk: value})
	}

	return resp.Value{Typ: common.ARRAY_TYPE, Array: values}
}
