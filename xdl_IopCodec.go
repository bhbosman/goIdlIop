// Code generated by me. DO NOT EDIT.

package goIdlIop

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "IOP::Codec", generatedBy by: "WriteInterface"
type IopCodec interface {
	//Exceptions for : Decode
	//	IopCodecFormatMismatch
	// Original name: "decode"
	Decode(params IopCodecDecodeIn) (IopCodecDecodeOut, error)
	//Exceptions for : DecodeValue
	//	IopCodecFormatMismatch
	//	IopCodecTypeMismatch
	// Original name: "decode_value"
	DecodeValue(params IopCodecDecodeValueIn) (IopCodecDecodeValueOut, error)
	//Exceptions for : Encode
	//	IopCodecInvalidTypeForEncoding
	// Original name: "encode"
	Encode(params IopCodecEncodeIn) (IopCodecEncodeOut, error)
	//Exceptions for : EncodeValue
	//	IopCodecInvalidTypeForEncoding
	// Original name: "encode_value"
	EncodeValue(params IopCodecEncodeValueIn) (IopCodecEncodeValueOut, error)
}

const IopCodecDecodeOperation = "decode"
const IopCodecDecodeValueOperation = "decode_value"
const IopCodecEncodeOperation = "encode"
const IopCodecEncodeValueOperation = "encode_value"
//noinspection GoSnakeCaseUsage
type IopCodec_Helper struct {
}

//noinspection GoSnakeCaseUsage
const IopCodecId_Const = "IDL:omg.org/IOP/Codec:1.0"

func (self IopCodec_Helper) Id() string {
	return IopCodecId_Const
}

func (self IopCodec_Helper) Read(stream __goidl__.IReadAny) (IopCodec, error) {
	return nil, nil
}

func (self IopCodec_Helper) Write(stream __goidl__.IWriteAny, v IopCodec) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var IopCodecHelper = IopCodec_Helper{}

func init() {
}
