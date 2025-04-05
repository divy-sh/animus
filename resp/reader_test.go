package resp_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestReadBulkString(t *testing.T) {
	input := "$5\r\nhello\r\n"
	r := resp.NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := resp.Value{Typ: common.BULK_TYPE, Bulk: "hello"}
	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadArray(t *testing.T) {
	input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	r := resp.NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := resp.Value{
		Typ:   "array",
		Array: []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}},
	}

	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadEmptyArray(t *testing.T) {
	input := "*0\r\n"
	r := resp.NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := resp.Value{
		Typ:   "array",
		Array: []resp.Value{},
	}

	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadInvalidEssentia(t *testing.T) {
	input := "?"
	r := resp.NewReader(bytes.NewBufferString(input))
	val, _ := r.Read()
	expected := resp.Value{}
	if !reflect.DeepEqual(val, expected) {
		t.Errorf("expected %v, but got %v", expected, val)
	}
}
