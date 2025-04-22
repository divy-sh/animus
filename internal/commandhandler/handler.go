package commandhandler

import (
	"github.com/divy-sh/animus/internal/resp"
)

// Command represents a command with an associated function and documentation.
type Command struct {
	Func func([]resp.Value) resp.Value
	Doc  string
}

// Handlers maps command names to their implementations.
var Handlers = map[string]Command{}

// RegisterCommand registers a command function with its documentation.
func RegisterCommand(name string, fn func([]resp.Value) resp.Value, doc string) {
	Handlers[name] = Command{Func: fn, Doc: doc}
}
