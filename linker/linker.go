package linker

type ForeachFunc func(data interface{}) error

type Node interface {
	Next()
}
