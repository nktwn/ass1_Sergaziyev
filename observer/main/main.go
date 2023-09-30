package main

import (
	"SDP/observer"
	"fmt"
)

func main() {
	mediator := observer.NewMediator()

	subject1 := observer.NewSubjectImpl("Software Design Patterns", "1", "Mukatayev Tleuzhan")
	subject2 := observer.NewSubjectImpl("Operating Systems", "2", "Apai")

	user1 := observer.NewUser("nurik@gmail.com", []string{"Software Design Patterns", "Operating Systems"})
	user2 := observer.NewUser("shurik@gmail.com", []string{"Software Design Patterns"})

	mediator.AddSubject(subject1)
	mediator.AddSubject(subject2)
	mediator.AddObserver(user1)
	mediator.AddObserver(user2)

	subject1.NotifyObservers()
	fmt.Println(subject1.Name, subject1.Code, subject1.Teacher)
	subject2.NotifyObservers()
	fmt.Println(subject2.Name, subject2.Code, subject2.Teacher)
}
