package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gonutz/tic/cmd"
)

func main() {
	corrupt := false
	defer func() {
		if corrupt {
			fmt.Println("time stack file is corrupt, deleting it, please call tic again")
			os.Remove(cmd.Path())
		}
	}()
	now := time.Now()

	f, err := os.OpenFile(cmd.Path(), os.O_RDWR, 0x666)
	if err != nil {
		fmt.Println("call tic before toc")
		return
	}
	defer f.Close()

	_, err = f.Seek(-1, os.SEEK_END)
	if err != nil {
		fmt.Println("call tic before toc")
		return
	}
	var size [1]byte
	_, err = f.Read(size[:])
	if err != nil {
		corrupt = true
		return
	}
	newEnd, err := f.Seek(-1-int64(size[0]), os.SEEK_END)
	if err != nil {
		corrupt = true
		return
	}
	data := make([]byte, size[0])
	_, err = f.Read(data)
	if err != nil {
		corrupt = true
		return
	}
	var t time.Time
	err = t.UnmarshalBinary(data)
	if err != nil {
		corrupt = true
		return
	}
	fmt.Println(now.Sub(t))
	f.Truncate(newEnd)
}
