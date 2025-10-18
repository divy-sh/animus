package command

import (
	"strings"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

// A simple in-memory config store
var serverConfig = map[string]string{
	"maxmemory": "0",
	"timeout":   "0",
	"save":      "3600 1 300 100 60 10000", // simulate default Redis RDB save points
}

// This is here just so that redis-benchmark doesn't complain.
// ConfigCmd implements the Redis CONFIG command.
// It supports CONFIG GET <parameter> and CONFIG SET <parameter> <value>
func ConfigCmd(args []resp.Value) resp.Value {
	if len(args) < 2 {
		return resp.Value{
			Typ: common.ERROR_TYPE,
			Str: common.ERR_WRONG_ARGUMENT_COUNT,
		}
	}

	subcommand := strings.ToUpper(args[0].Bulk)
	switch subcommand {
	case "GET":
		if len(args) != 2 {
			return resp.Value{
				Typ: common.ERROR_TYPE,
				Str: common.ERR_WRONG_ARGUMENT_COUNT,
			}
		}

		param := strings.ToLower(args[1].Bulk)
		// support wildcard "*"
		if param == "*" {
			array := make([]resp.Value, 0, len(serverConfig)*2)
			for k, v := range serverConfig {
				array = append(array, resp.Value{Typ: common.BULK_TYPE, Bulk: k})
				array = append(array, resp.Value{Typ: common.BULK_TYPE, Bulk: v})
			}
			return resp.Value{
				Typ:   common.ARRAY_TYPE,
				Array: array,
			}
		}

		value, ok := serverConfig[param]
		if !ok {
			return resp.Value{
				Typ: common.ERROR_TYPE,
				Str: "ERR unknown configuration parameter",
			}
		}

		return resp.Value{
			Typ: common.ARRAY_TYPE,
			Array: []resp.Value{
				{Typ: common.BULK_TYPE, Bulk: param},
				{Typ: common.BULK_TYPE, Bulk: value},
			},
		}

	case "SET":
		if len(args) != 3 {
			return resp.Value{
				Typ: common.ERROR_TYPE,
				Str: common.ERR_WRONG_ARGUMENT_COUNT,
			}
		}

		param := strings.ToLower(args[1].Bulk)
		value := args[2].Bulk // fix: value is at index 2

		if _, ok := serverConfig[param]; !ok {
			return resp.Value{
				Typ: common.ERROR_TYPE,
				Str: "ERR unknown configuration parameter",
			}
		}

		serverConfig[param] = value

		return resp.Value{
			Typ: common.STRING_TYPE,
			Str: "OK",
		}

	default:
		return resp.Value{
			Typ: common.ERROR_TYPE,
			Str: "ERR unknown subcommand, must be GET or SET",
		}
	}
}
