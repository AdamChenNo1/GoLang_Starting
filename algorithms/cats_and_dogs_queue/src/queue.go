package cats_dogs_queue

import (
	dsa "algorithms/stack_queue/src"
)

type Ele struct {
	pet   Pet
	count int64
}

type DogCatQueue struct {
	dogQ  dsa.StackQueue
	catQ  dsa.StackQueue
	count int64
}

func NewDogCatQueue() *DogCatQueue {
	return &DogCatQueue{dsa.StackQueue{}, dsa.StackQueue{}, 0}
}

func (dogCatQueue *DogCatQueue) Add(pet Pet) {
	types := pet.GetPetType()
	if types == "dog" {
		dogCatQueue.count++
		dogCatQueue.dogQ.Enqueue(Ele{pet, dogCatQueue.count})
	} else if types == "cat" {
		dogCatQueue.count++
		dogCatQueue.catQ.Enqueue(Ele{pet, dogCatQueue.count})
	} else {
		panic("not dog or cat")
	}
}

func (dogCatQueue *DogCatQueue) PollPet() Pet {
	dogQEmpty := dogCatQueue.dogQ.Size() == 0
	catQEmpty := dogCatQueue.catQ.Size() == 0
	if !dogQEmpty && !catQEmpty {
		dogPeek, ok1 := dogCatQueue.dogQ.Peek().(Ele)
		catPeek, ok2 := dogCatQueue.catQ.Peek().(Ele)
		if ok1 && ok2 {
			if dogPeek.count < catPeek.count {
				return dogCatQueue.dogQ.Dequeue().(Ele).pet
			} else {
				return dogCatQueue.catQ.Dequeue().(Ele).pet
			}
		} else {
			return nil
		}
	} else if !dogQEmpty {
		if dogPeek, ok := dogCatQueue.dogQ.Peek().(Ele); ok {
			return dogPeek.pet
		} else {
			return nil
		}
	} else if !catQEmpty {
		if catPeek, ok := dogCatQueue.catQ.Peek().(Ele); ok {
			return catPeek.pet
		} else {
			return nil
		}
	} else {
		panic("queue is empty")
	}
}

func (dogCatQueue *DogCatQueue) PollDog() Dog {
	if !(dogCatQueue.dogQ.Size() == 0) {
		return dogCatQueue.dogQ.Dequeue().(Ele).pet.(Dog)
	} else {
		panic("queue is empty")
	}
}

func (dogCatQueue *DogCatQueue) PollCat() Cat {
	if !(dogCatQueue.catQ.Size() == 0) {
		return dogCatQueue.catQ.Dequeue().(Ele).pet.(Cat)
	} else {
		panic("queue is empty")
	}
}

func (dogCatQueue *DogCatQueue) IsEmpty() bool {
	return dogCatQueue.dogQ.Size() == 0 && dogCatQueue.catQ.Size() == 0
}

func (dogCatQueue *DogCatQueue) IsDogEmpty() bool {
	return dogCatQueue.dogQ.Size() == 0
}

func (dogCatQueue *DogCatQueue) IsCatEmpty() bool {
	return dogCatQueue.catQ.Size() == 0
}

func (dogCatQueue *DogCatQueue) Size() int {
	return dogCatQueue.catQ.Size() + dogCatQueue.dogQ.Size()
}

func (dogCatQueue *DogCatQueue) CatSize() int {
	return dogCatQueue.catQ.Size()
}

func (dogCatQueue *DogCatQueue) DogSize() int {
	return dogCatQueue.dogQ.Size()
}
