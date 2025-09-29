package main

import (
	"fmt"
	"reflect"
)

type test struct {
	age int16
}

type stats struct {
	two   uint32
	one   uint8
	three uint8
}

type contact struct {
	userID       string // 16bytes
	sendingLimit int32  // 4bytes
	age          int32
}

type perms struct {
	permissionLevel int  // 8bytes -> in 64-bit system
	canSend         bool // 1byte
	canReceive      bool
	canMessage      bool
}

type wu struct {
	userID       string  // 16 bytes
	sendingLimit int32   // 4 bytes
	age          int16   // 2 bytes
	_            [2]byte // Padding to align to 8-byte boundary
}

func main() {
	typ := reflect.TypeOf(stats{})
	fmt.Println(typ.Size())
}
