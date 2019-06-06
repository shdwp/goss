package game

type Event interface {}

type GlobalEventType int
const (
	GlobalUpdateEvent GlobalEventType = iota + 1
)

type GlobalEvent struct {
	Kind GlobalEventType
}

type InputEvent struct {
	Input rune
}
