package rpc

import "rpc/codec"

const MagicNumber = 13

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: 13,
	CodecType:   codec.Type,
}
