package genericcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/command/stringcmd"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestGenerics_ExpireTime(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"}})
	if result.Typ != common.INTEGER_TYPE || result.Num != -1 {
		t.Errorf("Expected value %d, got: %v", -1, result)
	}
}

func TestGenerics_ExpireTime_InvalidKey(t *testing.T) {
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime_InvalidKey"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireTime_InvaliArgCount(t *testing.T) {
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}
