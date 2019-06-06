package game

// public types
type ProcessResult int
const (
	DestroyComponent ProcessResult = iota + 1
	SpawnUpdateEvent ProcessResult = iota + 1
	DoNothing        ProcessResult = iota + 1
)

type Component interface {
	Render(canvas RuneCanvas)
	Process(event Event, ch chan ProcessResult)
}

type Loop struct {
    canvas RuneCanvas

	componentStack []Component
}

func NewLoop(canvas RuneCanvas) Loop {
	var stack []Component
	return Loop{
		canvas,
		stack,
	}
}

func (self *Loop) DestroyComponent(c Component) {
	for i, component := range self.componentStack {
		if component == c {
			self.componentStack = append(self.componentStack[:i], self.componentStack[i+1:]...)
			return
		}
	}
}

func (self *Loop) AddComponent(c Component) {
	self.componentStack = append(self.componentStack, c)
}

func (self *Loop) Input(r rune) {
	event := InputEvent{r}
	if len(self.componentStack) == 0 {
        return
	}

	var chans resultChannelPool
	chans.process(event, self.componentStack[0])

	for len(chans.pool) > 0 {
		ch, comp := chans.pop()
		result := <- ch
		switch result {
		case SpawnUpdateEvent:
			for _, component := range self.componentStack {
				if comp == component {
					continue
				}

				chans.process(GlobalEvent{GlobalUpdateEvent}, component)
			}

		case DestroyComponent:
			self.DestroyComponent(chans.remove(ch))

		case DoNothing:
			continue
		}
	}
}

func (self *Loop) Render() {
	for _, c := range self.componentStack {
		c.Render(self.canvas)
	}
}

// private
type resultChannelEntry struct {
	component Component
	channel chan ProcessResult
}

type resultChannelPool struct {
	pool []resultChannelEntry
}

func (self *resultChannelPool) process(event Event, in Component) {
	ch := make(chan ProcessResult)
	self.pool = append(self.pool, resultChannelEntry{in, ch})
    go in.Process(event, ch)
}

func (self *resultChannelPool) remove(of chan ProcessResult) Component {
	idx := 0
	for i, e := range self.pool {
		if e.channel == of {
            idx = i
			break
		}
	}

	component := self.pool[idx].component
	self.pool = append(self.pool[:idx], self.pool[idx+1:]...)
    return component
}

func (self *resultChannelPool) pop() (chan ProcessResult, Component) {
    ch, c := self.pool[0].channel, self.pool[0].component
    self.pool = self.pool[1:]
    return ch, c
}

