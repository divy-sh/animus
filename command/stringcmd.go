package command

import (
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/types/strings"
)

func Append(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	strings.Append(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func DecrBy(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	err := strings.DecrBy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func Decr(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	err := strings.Decr(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func Get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.Get(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func GetDel(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.GetDel(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.NULL_TYPE}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func GetEx(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.GetEx(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.NULL_TYPE}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func GetRange(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	value, err := strings.GetRange(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	if err != nil {
		return resp.Value{Typ: common.NULL_TYPE}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: value}
}

func GetSet(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	val, err := strings.GetSet(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: val}
}

func IncrBy(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	err := strings.IncrBy(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func IncrByFloat(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	err := strings.IncrByFloat(args[0].Bulk, args[1].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func Incr(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	err := strings.Incr(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func LCS(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	commands := []string{}
	for _, arg := range args[2:] {
		commands = append(commands, arg.Bulk)
	}
	val, err := strings.Lcs(args[0].Bulk, args[1].Bulk, commands)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.BULK_TYPE, Bulk: val}
}

func MGet(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}
	keys := []string{}
	for _, arg := range args {
		keys = append(keys, arg.Bulk)
	}
	values := strings.MGet(&keys)
	response := make([]resp.Value, len(keys))
	for i, val := range *values {
		if val == "" {
			response[i] = resp.Value{Typ: common.NULL_TYPE}
		} else {
			response[i] = resp.Value{Typ: common.BULK_TYPE, Bulk: val}
		}
	}
	return resp.Value{Typ: common.ARRAY_TYPE, Array: response}
}

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

func Set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	strings.Set(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func SetEx(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	err := strings.SetEx(args[0].Bulk, args[1].Bulk, args[2].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.STRING_TYPE, Str: "OK"}
}

func StrLen(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: common.ERROR_TYPE, Str: common.ERR_WRONG_ARGUMENT_COUNT}
	}

	length, err := strings.StrLen(args[0].Bulk)
	if err != nil {
		return resp.Value{Typ: common.ERROR_TYPE, Str: err.Error()}
	}
	return resp.Value{Typ: common.INTEGER_TYPE, Num: length}
}
