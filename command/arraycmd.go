package command

import (
	"fmt"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/arrays"
)

func ArCount(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk

	count, err := arrays.ArCount(key)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: count}
}

func ArDel(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk
	index := args[1].Num

	err := arrays.ArDel(key, index)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}

func ArDelRange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk
	start := args[1].Num
	end := args[2].Num

	err := arrays.ArDelRange(key, start, end)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: "OK"}
}

func ArGet(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk
	index := args[1].Num

	value, err := arrays.ArGet(key, index)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	return resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(value)}
}

func ArGetRange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk
	start := args[1].Num
	end := args[2].Num

	values, err := arrays.ArGetRange(key, start, end)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	arrayValues := make([]resp.Value, len(values))
	for i, val := range values {
		arrayValues[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(val)}
	}

	return resp.Value{Typ: common.ARRAY_TYPE, Array: arrayValues}
}

func ArGrep(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	key := args[0].Bulk
	pattern := args[1].Bulk

	values, err := arrays.ArGrep(key, pattern)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	arrayValues := make([]resp.Value, len(values))
	for i, val := range values {
		arrayValues[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(val)}
	}

	return resp.Value{Typ: common.ARRAY_TYPE, Array: arrayValues}
}
