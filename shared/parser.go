package shared

import "fmt"

type Reader struct {
	input []rune
	pos   int
	len   int
}

func NewReader(input []rune) Reader {
	return Reader{
		input: input,
		pos:   0,
		len:   len(input),
	}
}

func (r *Reader) PeekNext() (rune, error) {
	return r.Peek(0)
}

func (r *Reader) Peek(k int) (rune, error) {
	idx := r.pos + k
	if idx >= r.len {
		return '0', fmt.Errorf("Requested index %d outside length %d", idx, r.len)
	}

	return r.input[idx], nil
}

func (r *Reader) ConsumeNext() (rune, error) {
	return r.Consume(0)
}

func (r *Reader) Consume(k int) (rune, error) {
	out, err := r.Peek(k)
	if err != nil {
		return '0', err
	}

	return out, nil
}

func (r *Reader) AtEnd() bool {
	return r.pos == r.len
}
