package resp_test

import (
	"bytes"
	"testing"

	"github.com/divy-sh/animus/resp"
)

func TestMarshalInt(t *testing.T) {
	v := resp.Value{Typ: "integer", Num: 42}
	expected := []byte{resp.INTEGER, byte(42), '\r', '\n'}
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalString(t *testing.T) {
	v := resp.Value{Typ: "string", Str: "Hello"}
	expected := append([]byte{resp.STRING}, "Hello"...)
	expected = append(expected, '\r', '\n')
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalBulk(t *testing.T) {
	v := resp.Value{Typ: "bulk", Bulk: "data"}
	expected := append([]byte{resp.BULK}, "4\r\ndata\r\n"...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalArray(t *testing.T) {
	v := resp.Value{Typ: "array", Array: []resp.Value{
		{Typ: "string", Str: "one"},
		{Typ: "integer", Num: 2},
	}}
	expected := append([]byte{resp.ARRAY}, "2\r\n"...)
	expected = append(expected, resp.Value{Typ: "string", Str: "one"}.Marshal()...)
	expected = append(expected, resp.Value{Typ: "integer", Num: 2}.Marshal()...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalError(t *testing.T) {
	v := resp.Value{Typ: "error", Str: "ERR error"}
	expected := append([]byte{resp.ERROR}, "ERR error\r\n"...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalNull(t *testing.T) {
	v := resp.Value{Typ: "null"}
	expected := []byte("$-1\r\n")
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	w := resp.NewWriter(&buf)
	v := resp.Value{Typ: "string", Str: "test"}

	err := w.Write(v)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := v.Marshal()
	if !bytes.Equal(buf.Bytes(), expected) {
		t.Errorf("Expected %q, got %q", expected, buf.Bytes())
	}
}
