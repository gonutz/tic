package main

import (
	"os"
	"time"

	"fmt"
	"github.com/gonutz/tic/cmd"
)

func main() {
	now := time.Now()

	f, err := os.OpenFile(cmd.Path(), os.O_RDWR, 0x666)
	check(err)
	defer f.Close()

	_, err = f.Seek(-1, os.SEEK_END)
	check(err)
	var size [1]byte
	_, err = f.Read(size[:])
	check(err)
	newEnd, err := f.Seek(-1-int64(size[0]), os.SEEK_END)
	check(err)
	data := make([]byte, size[0])
	_, err = f.Read(data)
	check(err)
	var t time.Time
	check(t.UnmarshalBinary(data))
	fmt.Println(now.Sub(t))
	f.Truncate(newEnd)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
