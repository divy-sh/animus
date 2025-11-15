package command

import (
	"strconv"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/lists"
)

func LIndex(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	index, err := strconv.ParseInt(args[1].Bulk, 10, 64)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	val, err := lists.Lindex(args[0].Bulk, index)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: val}
}

func LInsert(args []resp.Value) resp.Value {
	if len(args) != 4 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	newLen, err := lists.Linsert(args[0].Bulk, args[1].Bulk, args[2].Bulk, args[3].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: newLen}
}

func LLen(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	length, err := lists.LLen(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: length}
}

func LMove(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	val, err := lists.Lmove(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: val}
}

// func LMpop(args []resp.Value) resp.Value {
// 	if len(args) < 1 || len(args) > 2 {
// 		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
// 	}
// 	var countStr string
// 	if len(args) == 1 {
// 		countStr = "1"
// 	} else {
// 		countStr = args[1].Bulk
// 	}
// 	values, err := lists.Lmpop(args[0].Bulk, countStr)
// 	if err != nil {
// 		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
// 	}
// 	respArr := make([]resp.Value, len(values))
// 	for i, val := range values {
// 		respArr[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
// 	}
// 	return resp.Value{Typ: common.ARRAY_TYPE, Array: respArr}
// }

func LRange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	start, err := strconv.ParseInt(args[1].Bulk, 10, 64)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	end, err := strconv.ParseInt(args[2].Bulk, 10, 64)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	values, err := lists.Lrange(args[0].Bulk, start, end)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	respArr := make([]resp.Value, len(values))
	for i, val := range values {
		respArr[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: respArr}
}

func LPop(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	var values []string
	var err error
	if len(args) == 1 {
		values, err = lists.LPop(args[0].Bulk, "1")
	} else {
		values, err = lists.LPop(args[0].Bulk, args[1].Bulk)
	}
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	respArr := make([]resp.Value, len(values))
	for i, val := range values {
		respArr[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: respArr}
}

func LPush(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	values := make([]string, len(args)-1)
	for i, val := range args[1:] {
		values[i] = val.Bulk
	}
	lists.LPush(args[0].Bulk, &values)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func RPop(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	var values []string
	var err error
	if len(args) == 1 {
		values, err = lists.RPop(args[0].Bulk, "1")
	} else {
		values, err = lists.RPop(args[0].Bulk, args[1].Bulk)
	}
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	respArr := make([]resp.Value, len(values))
	for i, val := range values {
		respArr[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: respArr}
}

func RPush(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	values := make([]string, len(args)-1)
	for i, val := range args[1:] {
		values[i] = val.Bulk
	}
	lists.RPush(args[0].Bulk, &values)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}
