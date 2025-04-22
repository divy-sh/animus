package command

import (
	"fmt"
	"strings"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

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
