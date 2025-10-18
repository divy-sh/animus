package command

import (
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
