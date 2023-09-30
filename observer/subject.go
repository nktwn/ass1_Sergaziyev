package observer

type Subject interface {
	AddObserver(o Observer)
	DeleteObserver(o Observer)
	NotifyObservers()
}

type SubjectImpl struct {
	Name      string
	Code      string
	Teacher   string
	observers []Observer
}

func NewSubjectImpl(name, code, teacher string) *SubjectImpl {
	return &SubjectImpl{
		Name:      name,
		Code:      code,
		Teacher:   teacher,
		observers: []Observer{},
	}
}

func (s *SubjectImpl) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *SubjectImpl) DeleteObserver(o Observer) {
	index := -1
	for i, observer := range s.observers {
		if observer == o {
			index = i
			break
		}
	}
	if index != -1 {
		s.observers = append(s.observers[:index], s.observers[index+1:]...)
	}
}

func (s *SubjectImpl) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.Name, s.Code, s.Teacher)
	}
}
