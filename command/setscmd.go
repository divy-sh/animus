package command

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/sets"
)

func Sadd(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	key := args[0].Bulk
	var values []string
	for _, arg := range args[1:] {
		values = append(values, arg.Bulk)
	}
	count := sets.Sadd(key, values)
	return resp.Value{Typ: common.INTEGER_TYPE, Num: count}
}

func Scard(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	key := args[0].Bulk
	count := sets.Scard(key)
	return resp.Value{Typ: common.INTEGER_TYPE, Num: count}
}

func Sismember(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	key := args[0].Bulk
	value := args[1].Bulk
	isMember := sets.Sismember(key, value)
	var num int64
	if isMember {
		num = 1
	} else {
		num = 0
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: num}
}
