package server

import (
	"bytes"
	"errors"
)

var (
	ErrorWrongMsgType = errors.New("wrong msg type")
	ErrorMsgTooLong   = errors.New("message too long")
)

type Codec interface {
	Encode(msg any) (any, error)
	Decode(msg any) (any, error)
	Next() Codec
}

type lineCodec struct {
	buf      []byte
	capacity int

	next Codec
}

func NewLineCodec(capacity int, next Codec) Codec {
	return &lineCodec{
		buf:      make([]byte, 0, capacity),
		capacity: capacity,
		next:     next,
	}
}

func (c *lineCodec) Encode(msg any) (any, error) {
	if s, ok := msg.(string); ok {
		return []byte(s), nil
	} else {
		return nil, ErrorWrongMsgType
	}
}

func (c *lineCodec) Decode(msg any) (any, error) {
	if b, ok := msg.([]byte); ok {
		if len(b)+len(c.buf) > c.capacity {
			res := make([]byte, 0, 1024)
			part := b[:c.capacity-len(c.buf)]
			if p, err := c.Decode(part); err == nil {
				res = append(res, p.([]byte)...)
				if p, err := c.Decode(b[len(part):]); err == nil {
					return append(res, p.([]byte)...), nil
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		if len(c.buf) > c.capacity {
			return nil, ErrorMsgTooLong
		}
		c.buf = append(c.buf, b...)
		for {
			idx := bytes.Index(c.buf, []byte("\n"))
			if idx == -1 {
				return nil, nil
			}

		}

		//return string(c.buf[:pos]), nil
	} else {
		return nil, ErrorWrongMsgType
	}

	return nil, nil
}

func (c *lineCodec) Next() Codec {
	return c.next
}
