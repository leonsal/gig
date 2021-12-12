package imgui

// #include <stdlib.h>
import "C"

import (
	"bytes"
	"unsafe"
)

type CString []byte

//func NewCString(s string) *CString {
//
//	var cs CString
//	cs.SetString(s)
//	return &cs
//}

// Len returns the length of the CString not considering the final null terminator
func (cs *CString) Len() int {

	if *cs == nil {
		return 0
	}
	return len(*cs) - 1
}

// SetStrings sets this CString with the contents of the specified string
func (cs *CString) SetString(s string) {

	if len(s) >= len(*cs) {
		*cs = make([]byte, len(s)+1)
	}
	copy(*cs, s)
	(*cs)[len(s)] = 0
}

// String return the contents of this CString as a string
func (cs *CString) String() string {

	idx := bytes.IndexByte(*cs, 0)
	if idx < 0 {
		return ""
	}
	return string((*cs)[:idx])
}

// CString returns the pointer to this null terminated C string
func (cs *CString) CString() *C.char {

	if len(*cs) == 0 {
		return nil
	}
	return (*C.char)(unsafe.Pointer(&(*cs)[0]))
}

type CStrArr []*C.char

func (csa *CStrArr) Append(items ...string) {

	for _, it := range items {
		cs := C.CString(it)
		*csa = append(*csa, cs)
	}
}

func (csa *CStrArr) Clear() {

	for i := 0; i < len(*csa); i++ {
		C.free(unsafe.Pointer((*csa)[i]))
	}
	*csa = nil
}

func (csa *CStrArr) Ptr() **C.char {

	return (**C.char)(&(*csa)[0])
}
