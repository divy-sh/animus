package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

// CommandCmd implements the Redis COMMAND command.
// It returns metadata about all registered commands.
func CommandCmd(args []resp.Value) resp.Value {
	// The COMMAND command takes no arguments
	if len(args) != 0 {
		return resp.Value{
			Typ: common.ERROR_TYPE,
			Str: common.ERR_WRONG_ARGUMENT_COUNT,
		}
	}

	results := make([]resp.Value, 0, len(Handlers))

	for cmdName, handler := range Handlers {
		// Define metadata for each command
		cmdMeta := resp.Value{
			Typ: common.ARRAY_TYPE,
			Array: []resp.Value{
				{Typ: common.BULK_TYPE, Bulk: strings.ToLower(cmdName)},    // Command name
				{Typ: common.BULK_TYPE, Bulk: strconv.Itoa(handler.Arity)}, // Arity
				{Typ: common.ARRAY_TYPE, Array: func() []resp.Value { // Flags
					flags := []resp.Value{}
					for _, flag := range handler.Flags {
						flags = append(flags, resp.Value{Typ: common.BULK_TYPE, Bulk: flag})
					}
					return flags
				}()},
				{Typ: common.BULK_TYPE, Bulk: strconv.Itoa(handler.FirstKey)}, // First key position (not implemented)
				{Typ: common.BULK_TYPE, Bulk: strconv.Itoa(handler.LastKey)},  // Last key position (not implemented)
				{Typ: common.BULK_TYPE, Bulk: strconv.Itoa(handler.Step)},     // Step count for keys (not implemented)
			},
		}
		results = append(results, cmdMeta)
	}

	return resp.Value{
		Typ:   common.ARRAY_TYPE,
		Array: results,
	}
}

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

func Help(args []resp.Value) resp.Value {
	if len(args) == 0 {
		var docs []string
		for cmd, handler := range Handlers {
			docs = append(docs, fmt.Sprintf("%s - %s", cmd, handler.Doc))
		}
		return resp.Value{Typ: common.BULK_TYPE, Bulk: strings.Join(docs, "\n")}
	}

	cmd := strings.ToUpper(args[0].Bulk)
	if handler, exists := Handlers[cmd]; exists {
		return resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprintf("%s - %s", cmd, handler.Doc)}
	}
	return resp.Value{Typ: common.ERROR_TYPE, Str: "Unknown command: " + cmd}
}

func Info(args []resp.Value) resp.Value {
	info := `# Server\n
		redis_version:0.0.1-animus\n
		redis_mode:standalone\n
		os:darwin-arm64\n`

	return resp.Value{
		Typ:  common.BULK_TYPE,
		Bulk: info,
	}
}

func Ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: common.STRING_TYPE, Str: "PONG"}
	}

	return resp.Value{Typ: common.STRING_TYPE, Str: args[0].Bulk}
}
