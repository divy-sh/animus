package hashcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func Test_HExists(t *testing.T) {
	HSet([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "TestHExists",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "TestHExists",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "value",
		},
	})
	result := HExists([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "TestHExists",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "TestHExists",
		},
	})
	if result.Typ != common.INTEGER_TYPE || result.Num != 1 {
		t.Errorf("Expected hash to exist, got %v", result)
	}
}

func Test_HExists_Nope(t *testing.T) {
	result := HExists([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "Test_HExists_Nope",
		},
		{
			Typ:  common.BULK_TYPE,
			Bulk: "Test_HExists_Nope",
		},
	})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected hash to not exist, got %v", result)
	}
}

func TestHExistsInvalidCommandSize(t *testing.T) {
	result := HExists([]resp.Value{
		{
			Typ:  common.BULK_TYPE,
			Bulk: "hash",
		},
	})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'HGet' command but got %v", result)
	}
}
