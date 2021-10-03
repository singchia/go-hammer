package fan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_Fanout(t *testing.T) {
	fanout := NewBufferedFanout()
	go func() {
		for i := 0; i < 20; i++ {
			fanout.In(i)
		}
	}()

	out1 := fanout.Out(context.TODO())
	out2 := fanout.Out(context.TODO())
	go func() {
		for {
			out1data := <-out1
			fmt.Printf("first: %d\n", out1data.(int))
		}
	}()
	go func() {
		for {
			out2data := <-out2
			fmt.Printf("second: %d\n", out2data.(int))
		}
	}()
	time.Sleep(10 * time.Second)
}
