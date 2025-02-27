package resp

import (
	"bytes"
	"testing"
)

func TestMarshalInt(t *testing.T) {
	v := Value{Typ: "integer", Num: 42}
	expected := []byte{INTEGER, byte(42), '\r', '\n'}
	result := v.marshalInt()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalString(t *testing.T) {
	v := Value{Typ: "string", Str: "Hello"}
	expected := append([]byte{STRING}, "Hello"...)
	expected = append(expected, '\r', '\n')
	result := v.marshalString()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalBulk(t *testing.T) {
	v := Value{Typ: "bulk", Bulk: "data"}
	expected := append([]byte{BULK}, "4\r\ndata\r\n"...)
	result := v.marshalBulk()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalArray(t *testing.T) {
	v := Value{Typ: "array", Array: []Value{
		{Typ: "string", Str: "one"},
		{Typ: "integer", Num: 2},
	}}
	expected := append([]byte{ARRAY}, "2\r\n"...)
	expected = append(expected, Value{Typ: "string", Str: "one"}.Marshal()...)
	expected = append(expected, Value{Typ: "integer", Num: 2}.Marshal()...)
	result := v.marshalArray()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalError(t *testing.T) {
	v := Value{Typ: "error", Str: "ERR error"}
	expected := append([]byte{ERROR}, "ERR error\r\n"...)
	result := v.marshallError()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalNull(t *testing.T) {
	v := Value{Typ: "null"}
	expected := []byte("$-1\r\n")
	result := v.marshallNull()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	v := Value{Typ: "string", Str: "test"}

	err := w.Write(v)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := v.Marshal()
	if !bytes.Equal(buf.Bytes(), expected) {
		t.Errorf("Expected %q, got %q", expected, buf.Bytes())
	}
}
