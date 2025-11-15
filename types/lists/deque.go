package lists

type Deque[T any] struct {
	buf        []T
	head, tail int
	size       int
}

func NewDeque[T any](capHint int) *Deque[T] {
	if capHint < 4 {
		capHint = 4
	}
	return &Deque[T]{
		buf: make([]T, capHint),
	}
}

func (d *Deque[T]) grow() {
	newBuf := make([]T, len(d.buf)*2)
	if d.head < d.tail {
		copy(newBuf, d.buf[d.head:d.tail])
	} else {
		n := copy(newBuf, d.buf[d.head:])
		copy(newBuf[n:], d.buf[:d.tail])
	}
	d.head = 0
	d.tail = d.size
	d.buf = newBuf
}

func (d *Deque[T]) PushFront(v T) {
	if d.size == len(d.buf) {
		d.grow()
	}
	d.head = (d.head - 1 + len(d.buf)) % len(d.buf)
	d.buf[d.head] = v
	d.size++
}

func (d *Deque[T]) PushBack(v T) {
	if d.size == len(d.buf) {
		d.grow()
	}
	d.buf[d.tail] = v
	d.tail = (d.tail + 1) % len(d.buf)
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	v := d.buf[d.head]
	d.head = (d.head + 1) % len(d.buf)
	d.size--
	return v, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	d.tail = (d.tail - 1 + len(d.buf)) % len(d.buf)
	v := d.buf[d.tail]
	d.size--
	return v, true
}

func (d *Deque[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= d.size {
		return zero, false
	}
	return d.buf[(d.head+i)%len(d.buf)], true
}

func (d *Deque[T]) Set(i int, v T) bool {
	if i < 0 || i >= d.size {
		return false
	}
	d.buf[(d.head+i)%len(d.buf)] = v
	return true
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) ToSlice() []T {
	out := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		out[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	return out
}

func (d *Deque[T]) SliceRange(start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end >= d.size {
		end = d.size - 1
	}
	if start > end {
		return []T{}
	}
	out := make([]T, end-start+1)
	for i := 0; i < len(out); i++ {
		out[i] = d.buf[(d.head+start+i)%len(d.buf)]
	}
	return out
}

func (d *Deque[T]) InsertAt(idx int, v T) bool {
	if idx < 0 || idx > d.size {
		return false
	}
	if idx == 0 {
		d.PushFront(v)
		return true
	}
	if idx == d.size {
		d.PushBack(v)
		return true
	}

	// Make space
	if d.size == len(d.buf) {
		d.grow()
	}

	// Shift items either left or right
	// Decide cheaper direction
	if idx < d.size/2 {
		// Shift front part left
		d.head = (d.head - 1 + len(d.buf)) % len(d.buf)
		for i := 0; i < idx; i++ {
			from := (d.head + i + 1) % len(d.buf)
			to := (d.head + i) % len(d.buf)
			d.buf[to] = d.buf[from]
		}
		d.buf[(d.head+idx)%len(d.buf)] = v
	} else {
		// Shift back part right
		for i := d.size; i > idx; i-- {
			from := (d.head + i - 1) % len(d.buf)
			to := (d.head + i) % len(d.buf)
			d.buf[to] = d.buf[from]
		}
		d.buf[(d.head+idx)%len(d.buf)] = v
		d.tail = (d.tail + 1) % len(d.buf)
	}

	d.size++
	return true
}

func (d *Deque[T]) RemoveAt(idx int) (T, bool) {
	var zero T
	if idx < 0 || idx >= d.size {
		return zero, false
	}

	v, _ := d.Get(idx)

	// Same logic: shift cheapest way
	if idx < d.size/2 {
		// Shift left side right
		for i := idx; i > 0; i-- {
			from := (d.head + i - 1) % len(d.buf)
			to := (d.head + i) % len(d.buf)
			d.buf[to] = d.buf[from]
		}
		d.head = (d.head + 1) % len(d.buf)
	} else {
		// Shift right side left
		for i := idx; i < d.size-1; i++ {
			from := (d.head + i + 1) % len(d.buf)
			to := (d.head + i) % len(d.buf)
			d.buf[to] = d.buf[from]
		}
		d.tail = (d.tail - 1 + len(d.buf)) % len(d.buf)
	}

	d.size--
	return v, true
}
