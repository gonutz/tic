package main

import (
	"os"
	"time"

	"github.com/gonutz/tic/cmd"
)

func main() {
	f, err := os.OpenFile(cmd.Path(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	check(err)
	defer f.Close()

	data, err := time.Now().MarshalBinary()
	check(err)
	_, err = f.Write(append(data, byte(len(data))))
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
