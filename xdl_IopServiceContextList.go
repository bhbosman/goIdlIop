// Code generated by me. DO NOT EDIT.

package goIdlIop

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "IOP::ServiceContextList", generatedBy by: "WriteStructSequenceDcl"
type IopServiceContextList struct {
	__goidl__.IdlObject
	Array []*IopServiceContext `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const IopServiceContextListId_Const = "IDL:IOP/ServiceContextList:1.0"

func (self *IopServiceContextList) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IopServiceContextList) GoString() string {
	return self.String()
}

func (self *IopServiceContextList) ReadValue(stream __goidl__.IReadAny) error {
	err := self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	var n uint32
	n, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	if n > 0 {
		self.Array = make([]*IopServiceContext, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &IopServiceContext{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *IopServiceContextList) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IopServiceContextList) Write(stream __goidl__.IWriteAny) error {
	err := self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	err = __goidl__.IdlUInt32Helper.Write(stream, uint32(len(self.Array)))
	if err != nil {
	return err
	}
	if len(self.Array) > 0 {
		for _, item := range self.Array {
			err = item.Write(stream)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IopServiceContextList_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IopServiceContextListHelper = IopServiceContextList_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IopServiceContextListId_Const,
			__goidl__.SequenceType,
			"IOP.idl",
			"xdl_IopServiceContextList.go",
			func() __goidl__.IIdlObject {
				return &IopServiceContextList{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IopServiceContextList{}
			},
			__reflect__.TypeOf((*IopServiceContextList)(nil))))
}