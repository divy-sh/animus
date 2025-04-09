package commands

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestHandler_Help_NoArgs(t *testing.T) {
	response := Help([]resp.Value{})
	if response.Typ != common.BULK_TYPE {
		t.Errorf("expected %s type for Help, got %v", common.BULK_TYPE, response.Typ)
	}
}

func TestHandler_Help(t *testing.T) {
	for key, val := range Handlers {
		response := Help([]resp.Value{{Typ: common.BULK_TYPE, Bulk: key}})
		if response.Typ != common.BULK_TYPE || response.Bulk != key+" - "+val.Doc {
			t.Errorf("expected response %s for command %s, got %v", val.Doc, key, response)
		}
	}
}

func TestHandler_Help_InvalidCommand(t *testing.T) {
	response := Help([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "invalid_command"}})
	if response.Typ != common.ERROR_TYPE || response.Str != "Unknown command: INVALID_COMMAND" {
		t.Errorf("expected error %s, got %v", "Unknown command: INVALID_COMMAND", response)
	}
}
