package hashcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestHsetAndHGet(t *testing.T) {
	input := []resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "key",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "value",
		},
	}
	result := HSet(input)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected success but got type: %s, value: %s", result.Typ, result.Str)
	}
	result = HGet([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "key",
		},
	})
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected success but got type: %s, value: %s", result.Typ, result.Str)
	}
}

func TestHGetWithoutHset(t *testing.T) {
	result := HGet([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "not_set",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "not_set",
		},
	})
	expected := common.ERR_HASH_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestHsetInvalidCommandSize(t *testing.T) {
	input := []resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "key",
		},
	}
	result := HSet(input)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'Hset' command but got %v", result)
	}
}

func TestHGetInvalidCommandSize(t *testing.T) {
	result := HGet([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
	})
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'HGet' command but got %v", result)
	}
}
