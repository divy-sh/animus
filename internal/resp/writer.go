package resp

import (
	"io"
	"strconv"

	"github.com/divy-sh/animus/internal/common"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (v Value) Marshal() []byte {
	switch v.Typ {
	case common.ARRAY_TYPE:
		return v.marshalArray()
	case common.BULK_TYPE:
		return v.marshalBulk()
	case common.INTEGER_TYPE:
		return v.marshalInt()
	case common.STRING_TYPE:
		return v.marshalString()
	case common.NULL_TYPE:
		return v.marshallNull()
	case common.ERROR_TYPE:
		return v.marshallError()
	default:
		return []byte{}
	}
}

func (v Value) marshalInt() []byte {
	var bytes []byte
	bytes = append(bytes, INTEGER)
	bytes = append(bytes, byte(v.Num))
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalArray() []byte {
	len := len(v.Array)
	var bytes []byte
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}

	return bytes
}

func (v Value) marshallError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}

func (w *Writer) Write(v Value) error {
	var bytes = v.Marshal()
	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
