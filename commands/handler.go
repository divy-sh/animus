package commands

import (
	"github.com/divy-sh/animus/resp"
)

var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	"HSET": hset,
	"HGET": hget,
}
