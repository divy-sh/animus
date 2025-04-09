package commands

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestHsetAndHget(t *testing.T) {
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
	result := hset(input)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected success but got type: %s, value: %s", result.Typ, result.Str)
	}
	result = hget([]resp.Value{
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

func TestHgetWithoutHset(t *testing.T) {
	result := hget([]resp.Value{
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
	result := hset(input)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'hset' command but got %v", result)
	}
}

func TestHGetInvalidCommandSize(t *testing.T) {
	result := hget([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
	})
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'hget' command but got %v", result)
	}
}
