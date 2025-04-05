package commands

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestPingNoArg(t *testing.T) {
	result := ping([]resp.Value{})
	if result.Typ != common.STRING_TYPE || result.Str != "PONG" {
		t.Errorf("expected PONG, got %v", result)
	}
}

func TestPingWithArg(t *testing.T) {
	result := ping([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "test",
		},
	})
	if result.Typ != common.STRING_TYPE || result.Str != "test" {
		t.Errorf("expected %s, got %v", "test", result)
	}
}
