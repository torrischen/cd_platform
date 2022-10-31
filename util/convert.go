package util

import "unsafe"

func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		cap int
	}{s, len(s)}))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func OwnerToNS(owner string) string {
	return owner + "-" + "workspace"
}
