package chapter_3

import (
	"errors"
	"time"
)

type Animal interface {
	GetArrivalTime() int64
}

type Cat struct {
	arrivalTime int64
}

type Dog struct {
	arrivalTime int64
}

func (c Cat) GetArrivalTime() int64 {
	return c.arrivalTime
}

func NewCat() Cat {
	return Cat{time.Now().Unix()}
}

func (d Dog) GetArrivalTime() int64 {
	return d.arrivalTime
}

func NewDog() Dog {
	return Dog{time.Now().Unix()}
}

type CatQueue chan Cat
type DogQueue chan Dog

type AnimalShelter struct {
	cats, dogs *Queue
}

func NewAnimalShelter() AnimalShelter {
	return AnimalShelter{NewQueue(), NewQueue()}
}

func (q CatQueue) dequeueCat() (Cat, bool) {
	select {
	case c := <- q:
		return c, true
	default:
		return Cat{}, false
	}
}

func (q CatQueue) enqueueCat(c Cat) bool {
	select {
	case q <- c:
		return true
	default:
		return false
	}
}

func (q DogQueue) enqueueDog(d Dog) bool {
	select {
	case q <- d:
		return true
	default:
		return false
	}
}

func (a AnimalShelter) IsEmpty() bool {
	return a.cats.IsEmpty() && a.dogs.IsEmpty()
}

func (a AnimalShelter) Enqueue(animal Animal) error {
	if cat, ok := animal.(Cat); ok {
		a.cats.Enqueue(cat)
	} else if dog, ok := animal.(Dog); ok {
		a.dogs.Enqueue(dog)
	} else {
		return errors.New("this shelter doesn't house that type of animal")
	}
	return nil
}

func (a AnimalShelter) DequeueDog() Dog {
	if a.dogs.IsEmpty() {
		panic("DequeueDog() called with no dogs")
	}
	return a.dogs.Dequeue().(Dog)
}

func (a AnimalShelter) DequeueCat() Cat {
	if a.cats.IsEmpty() {
		panic("DequeueDog() called with no cats")
	}
	return a.cats.Dequeue().(Cat)
}

func (a AnimalShelter) DequeueAny() Animal {
	noCats, noDogs := a.cats.IsEmpty(), a.dogs.IsEmpty()
	if noCats && noDogs {
		panic("DequeueAny() called with no animals in shelter")
	} else if noCats {
		return a.DequeueDog()
	} else if noDogs {
		return a.DequeueCat()
	}
	cat := a.cats.Peek().(Cat)
	dog := a.dogs.Peek().(Dog)
	if cat.GetArrivalTime() < dog.GetArrivalTime() {
		return a.DequeueCat()
	} else {
		return a.DequeueDog()
	}
}
