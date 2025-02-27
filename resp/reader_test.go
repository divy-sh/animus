package resp

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadBulkString(t *testing.T) {
	input := "$5\r\nhello\r\n"
	r := NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := Value{Typ: "bulk", Bulk: "hello"}
	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadArray(t *testing.T) {
	input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	r := NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := Value{
		Typ:   "array",
		Array: []Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: "world"}},
	}

	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadEmptyArray(t *testing.T) {
	input := "*0\r\n"
	r := NewReader(bytes.NewBufferString(input))
	val, err := r.Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := Value{
		Typ:   "array",
		Array: []Value{},
	}

	if !reflect.DeepEqual(val, expected) {
		t.Errorf("Expected %v, got %v", expected, val)
	}
}

func TestReadInvalidType(t *testing.T) {
	input := "?"
	r := NewReader(bytes.NewBufferString(input))
	val, _ := r.Read()
	expected := Value{}
	if !reflect.DeepEqual(val, expected) {
		t.Errorf("expected %v, but got %v", expected, val)
	}
}
