package observer

type Observer interface {
	Update(data interface{})
}
