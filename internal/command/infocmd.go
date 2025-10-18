package command

import (
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

// Info implements the Redis INFO command (minimal version)
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
