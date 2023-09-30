package observer

type Mediator struct {
	subjects  []Subject
	observers []Observer
}

func NewMediator() *Mediator {
	return &Mediator{
		subjects:  []Subject{},
		observers: []Observer{},
	}
}

func (m *Mediator) AddSubject(s Subject) {
	m.subjects = append(m.subjects, s)
}

func (m *Mediator) AddObserver(o Observer) {
	m.observers = append(m.observers, o)
}

func (m *Mediator) NotifyObservers() {
	for _, subject := range m.subjects {
		subject.NotifyObservers()
	}
}
