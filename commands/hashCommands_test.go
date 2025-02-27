package commands

import (
	"testing"

	"github.com/divy-sh/animus/resp"
)

func TestHsetAndHget(t *testing.T) {
	input := []resp.Value{
		{
			Typ:  "bulk",
			Bulk: "hash",
		},
		{
			Typ:  "bulk",
			Bulk: "key",
		},
		{
			Typ:  "bulk",
			Bulk: "value",
		},
	}
	result := hset(input)
	if result.Typ != "string" || result.Str != "OK" {
		t.Errorf("Expected success but got type: %s, value: %s", result.Typ, result.Str)
	}
	result = hget([]resp.Value{
		{
			Typ:  "bulk",
			Bulk: "hash",
		},
		{
			Typ:  "bulk",
			Bulk: "key",
		},
	})
	if result.Typ != "bulk" || result.Bulk != "value" {
		t.Errorf("Expected success but got type: %s, value: %s", result.Typ, result.Str)
	}
}

func TestHgeteithoutHset(t *testing.T) {
	result := hget([]resp.Value{
		{
			Typ:  "bulk",
			Bulk: "hash",
		},
		{
			Typ:  "bulk",
			Bulk: "key",
		},
	})
	if result.Typ != "null" || result.Bulk != "" {
		t.Errorf("Expected null but got %v", result)
	}
}

func TestHsetInvalidCommandSize(t *testing.T) {
	input := []resp.Value{
		{
			Typ:  "bulk",
			Bulk: "hash",
		},
		{
			Typ:  "bulk",
			Bulk: "key",
		},
	}
	result := hset(input)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'hset' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'hset' command but got %v", result)
	}
}

func TestHGetInvalidCommandSize(t *testing.T) {
	result := hget([]resp.Value{
		{
			Typ:  "bulk",
			Bulk: "hash",
		},
	})
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'hget' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'hget' command but got %v", result)
	}
}
