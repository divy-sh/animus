package helpcmd

import (
	"fmt"
	"strings"

	"github.com/divy-sh/animus/internal/commandhandler"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func Help(args []resp.Value) resp.Value {
	if len(args) == 0 {
		var docs []string
		for cmd, handler := range commandhandler.Handlers {
			docs = append(docs, fmt.Sprintf("%s - %s", cmd, handler.Doc))
		}
		return resp.Value{Typ: common.BULK_TYPE, Bulk: strings.Join(docs, "\n")}
	}

	cmd := strings.ToUpper(args[0].Bulk)
	if handler, exists := commandhandler.Handlers[cmd]; exists {
		return resp.Value{Typ: common.BULK_TYPE, Bulk: fmt.Sprintf("%s - %s", cmd, handler.Doc)}
	}
	return resp.Value{Typ: common.ERROR_TYPE, Str: "Unknown command: " + cmd}
}

func init() {
	commandhandler.RegisterCommand("HELP", Help, `HELP [COMMAND]
	Shows documentation for available commands.`)
}
