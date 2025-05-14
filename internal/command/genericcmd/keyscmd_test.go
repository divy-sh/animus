package genericcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/command/stringcmd"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestKeys(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestKeys"},
		{Typ: common.BULK_TYPE, Bulk: "value"},
	})

	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "Test.eys"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}
	if len(result.Array) != 1 || result.Array[0].Bulk != "TestKeys" {
		t.Errorf("expected ['TestKeys'], got %v", result.Array)
	}
}

func TestKeysWrongArgCount(t *testing.T) {
	result := Keys([]resp.Value{})

	if result.Typ != common.ERROR_TYPE {
		t.Errorf("expected ERROR_TYPE, got %v", result.Typ)
	}
	if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("expected error '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}

func TestKeysInvalidRegex(t *testing.T) {
	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "[a-b"}})

	if result.Typ != common.ERROR_TYPE {
		t.Errorf("expected ERROR_TYPE, got %v", result.Typ)
	}
	if result.Str != common.ERR_INVALID_REGEX {
		t.Errorf("expected error '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}

func TestKeysMultipleMatches(t *testing.T) {
	stringcmd.Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:1"}, {Typ: common.BULK_TYPE, Bulk: "red"}})
	stringcmd.Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:2"}, {Typ: common.BULK_TYPE, Bulk: "green"}})
	stringcmd.Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "banana:1"}, {Typ: common.BULK_TYPE, Bulk: "yellow"}})

	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:*"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}

	got := map[string]bool{}
	for _, v := range result.Array {
		got[v.Bulk] = true
	}

	if !(got["apple:1"] && got["apple:2"]) || len(result.Array) != 2 {
		t.Errorf("expected ['apple:1', 'apple:2'], got %v", result.Array)
	}
}

func TestKeysNoMatch(t *testing.T) {
	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "nonexistent*"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}

	if len(result.Array) != 0 {
		t.Errorf("expected empty result, got %v", result.Array)
	}
}
