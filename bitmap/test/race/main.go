package main

import (
	"math/rand"

	"github.com/singchia/bitindex"
)

func main() {
	bi := bitindex.NewBitmap()
	x := rand.Uint32()
	c := make(chan bool)
	go func() {
		bi.Add(x)
		c <- true
	}()
	bi.Del(x)
	<-c
}
