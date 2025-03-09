package resp

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(rd io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(rd)}
}

func (r *Reader) Read() (Value, error) {
	valEssentia, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}
	switch valEssentia {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		log.Printf("Unknown type: %v", valEssentia)
		return Value{}, nil
	}
}

func (r *Reader) readArray() (Value, error) {
	len, err := r.readInt()
	if err != nil {
		return Value{}, err
	}
	v := Value{
		Typ:   "array",
		Array: make([]Value, len),
	}
	for i := 0; i < len; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.Array[i] = val
	}
	return v, nil
}

func (r *Reader) readBulk() (Value, error) {
	len, err := r.readInt()
	if err != nil {
		return Value{}, err
	}
	bulk := make([]byte, len)
	r.reader.Read(bulk)
	r.readLine()
	return Value{
		Typ:  "bulk",
		Bulk: string(bulk),
	}, nil
}

func (r *Reader) readInt() (val int, err error) {
	line, _, err := r.readLine()
	if err != nil {
		return 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

func (r *Reader) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		line = append(line, b)
		if len(line) > 1 && line[len(line)-1] == '\n' && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], len(line), nil
}
