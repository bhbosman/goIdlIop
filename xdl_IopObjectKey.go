// Code generated by me. DO NOT EDIT.

package goIdlIop

import __goIdlCorba__ "github.com/bhbosman/goIdlCorba"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "IOP::ObjectKey", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type IopObjectKey struct {
	__goIdlCorba__.CorbaOctetSeq
}

//noinspection GoSnakeCaseUsage
const IopObjectKeyId_Const = "IDL:IOP/ObjectKey:1.0"

func (self *IopObjectKey) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IopObjectKey) GoString() string {
	return self.String()
}

func (self *IopObjectKey) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.CorbaOctetSeq.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopObjectKey) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopObjectKey) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.CorbaOctetSeq.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IopObjectKey_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IopObjectKeyHelper = IopObjectKey_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IopObjectKeyId_Const,
			__goidl__.StructType,
			"IOP.idl",
			"xdl_IopObjectKey.go",
			func() __goidl__.IIdlObject {
				return &IopObjectKey{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IopObjectKey{}
			},
			__reflect__.TypeOf((*IopObjectKey)(nil))))
}
