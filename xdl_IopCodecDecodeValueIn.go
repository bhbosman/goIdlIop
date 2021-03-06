// Code generated by me. DO NOT EDIT.

package goIdlIop

import __goIdlCorba__ "github.com/bhbosman/goIdlCorba"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "IOP::Codec_decode_value_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type IopCodecDecodeValueIn struct {
	__goidl__.IdlObject
	Data __goIdlCorba__.CorbaOctetSeq `json:"Data"`
	Tc __goIdlCorba__.CorbaTypeCode `json:"Tc"`
}

//noinspection GoSnakeCaseUsage
const IopCodecDecodeValueInId_Const = "IDL:IOP/Codec_decode_value_In:1.0"

func (self *IopCodecDecodeValueIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IopCodecDecodeValueIn) GoString() string {
	return self.String()
}

func (self *IopCodecDecodeValueIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Data.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.Tc, err = __goIdlCorba__.CorbaTypeCodeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopCodecDecodeValueIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopCodecDecodeValueIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Data.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = __goIdlCorba__.CorbaTypeCodeHelper.Write(stream, self.Tc)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IopCodecDecodeValueIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IopCodecDecodeValueInHelper = IopCodecDecodeValueIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IopCodecDecodeValueInId_Const,
			__goidl__.StructType,
			"IOP.idl",
			"xdl_IopCodecDecodeValueIn.go",
			func() __goidl__.IIdlObject {
				return &IopCodecDecodeValueIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IopCodecDecodeValueIn{}
			},
			__reflect__.TypeOf((*IopCodecDecodeValueIn)(nil))))
}
