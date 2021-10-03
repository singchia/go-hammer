package fan

import "context"

type Fanout interface {
	In(data interface{})
	Out(context.Context) <-chan interface{}
}

type BufferedFanoutOption func(*BufferedFanout)

func OptionBufferdFanoutPipeSize(size int) BufferedFanoutOption {
	return func(bf *BufferedFanout) {
		bf.pipeSize = size
	}
}

func OptionBufferdFanoutOutSize(size int) BufferedFanoutOption {
	return func(bf *BufferedFanout) {
		bf.outSize = size
	}
}

type BufferedFanout struct {
	pipe     chan interface{}
	pipeSize int
	outSize  int
}

func NewBufferedFanout(options ...BufferedFanoutOption) Fanout {
	bf := &BufferedFanout{
		pipeSize: 0,
		outSize:  0,
	}

	for _, option := range options {
		option(bf)
	}
	bf.pipe = make(chan interface{}, bf.pipeSize)
	return bf
}

func (fanout *BufferedFanout) In(data interface{}) {
	fanout.pipe <- data
}

func (fanout *BufferedFanout) Out(ctx context.Context) <-chan interface{} {
	out := make(chan interface{}, fanout.outSize)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-fanout.pipe:
				out <- data
			}
		}
	}()
	return out
}
