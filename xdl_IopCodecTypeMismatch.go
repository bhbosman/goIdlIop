// Code generated by me. DO NOT EDIT.

package goIdlIop

import __fmt__ "fmt"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Exception declaration: "IOP::Codec::TypeMismatch", generatedBy by: "WriteStructDcl"
// Exception Decl: true
type IopCodecTypeMismatch struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const IopCodecTypeMismatchId_Const = "IDL:omg.org/IOP/Codec/TypeMismatch:1.0"

func (self *IopCodecTypeMismatch) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IopCodecTypeMismatch) Error() string{
	return 	__fmt__.Sprintf("Error of type IopCodecTypeMismatch(%v)", self.String())
}
func (self *IopCodecTypeMismatch) GoString() string {
	return self.String()
}

func (self *IopCodecTypeMismatch) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopCodecTypeMismatch) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopCodecTypeMismatch) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IopCodecTypeMismatch_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IopCodecTypeMismatchHelper = IopCodecTypeMismatch_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IopCodecTypeMismatchId_Const,
			__goidl__.StructType,
			"IOP.idl",
			"xdl_IopCodecTypeMismatch.go",
			func() __goidl__.IIdlObject {
				return &IopCodecTypeMismatch{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IopCodecTypeMismatch{}
			},
			__reflect__.TypeOf((*IopCodecTypeMismatch)(nil))))
}
