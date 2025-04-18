package resp_test

import (
	"bytes"
	"testing"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestMarshalInt(t *testing.T) {
	v := resp.Value{Typ: common.INTEGER_TYPE, Num: 42}
	expected := []byte{resp.INTEGER, byte(42), '\r', '\n'}
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalString(t *testing.T) {
	v := resp.Value{Typ: common.STRING_TYPE, Str: "Hello"}
	expected := append([]byte{resp.STRING}, "Hello"...)
	expected = append(expected, '\r', '\n')
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalBulk(t *testing.T) {
	v := resp.Value{Typ: common.BULK_TYPE, Bulk: "data"}
	expected := append([]byte{resp.BULK}, "4\r\ndata\r\n"...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalArray(t *testing.T) {
	v := resp.Value{Typ: common.ARRAY_TYPE, Array: []resp.Value{
		{Typ: common.STRING_TYPE, Str: "one"},
		{Typ: common.ARRAY_TYPE, Num: 2},
	}}
	expected := append([]byte{resp.ARRAY}, "2\r\n"...)
	expected = append(expected, resp.Value{Typ: common.STRING_TYPE, Str: "one"}.Marshal()...)
	expected = append(expected, resp.Value{Typ: common.ARRAY_TYPE, Num: 2}.Marshal()...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalError(t *testing.T) {
	v := resp.Value{Typ: common.ERROR_TYPE, Str: "ERR error"}
	expected := append([]byte{resp.ERROR}, "ERR error\r\n"...)
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalNull(t *testing.T) {
	v := resp.Value{Typ: common.NULL_TYPE}
	expected := []byte("$-1\r\n")
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarshalUnknownType(t *testing.T) {
	v := resp.Value{Typ: ""}
	expected := []byte{}
	result := v.Marshal()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	w := resp.NewWriter(&buf)
	v := resp.Value{Typ: common.STRING_TYPE, Str: "test"}

	err := w.Write(v)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := v.Marshal()
	if !bytes.Equal(buf.Bytes(), expected) {
		t.Errorf("Expected %q, got %q", expected, buf.Bytes())
	}
}
