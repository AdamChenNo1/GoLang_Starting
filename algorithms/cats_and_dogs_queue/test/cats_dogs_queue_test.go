package cats_dogs_queue

import (
	dsa "algorithms/cats_and_dogs_queue/src"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	stack := dsa.NewDogCatQueue()
	for i := 0; i < 10; i++ {
		id1 := fmt.Sprintf("%d", i)
		stack.Add(dsa.NewCat(id1))
	}

	for i := 0; i < 10; i++ {
		id2 := fmt.Sprintf("%d", i)
		stack.Add(dsa.NewDog(id2))
	}

	rig := stack.CatSize() == 10 && stack.DogSize() == 10 && stack.Size() == 20
	if rig {
		t.Log("Add succeed")
	} else {
		t.Error("Add failed")
	}
}

func TestPollPet(t *testing.T) {
	stack := dsa.NewDogCatQueue()
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		if i%2 == 0 {
			stack.Add(dsa.NewCat(id))
		} else {
			stack.Add(dsa.NewDog(id))
		}
	}

	for j := 0; j < 6; j++ {
		pet := stack.PollPet()

		modi := j % 2
		EvenIsCat := modi == 0 && pet.GetPetType() == "cat"
		OddIsDog := modi == 1 && pet.GetPetType() == "dog"

		if !EvenIsCat && !OddIsDog {
			t.Error("PollPet failed")
			break
		}
	}
	t.Log("PollPet succeed")
}

func TestPollCat(t *testing.T) {
	stack := dsa.NewDogCatQueue()
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		if i%2 == 0 {
			stack.Add(dsa.NewCat(id))
		} else {
			stack.Add(dsa.NewDog(id))
		}
	}

	for j := 0; j < 10; j += 2 {
		pet := stack.PollCat()

		modi := j % 2
		EvenIsCat := modi == 0 && pet.GetPetType() == "cat"
		if !EvenIsCat {
			t.Error("PollCat failed")
			break
		}
	}
	t.Log("PollCat succeed")
}

func TestPollDog(t *testing.T) {
	stack := dsa.NewDogCatQueue()
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		if i%2 == 0 {
			stack.Add(dsa.NewCat(id))
		} else {
			stack.Add(dsa.NewDog(id))
		}
	}

	for i := 1; i < 10; i += 2 {
		pet := stack.PollDog()

		modi := i % 2
		OddIsDog := modi == 1 && pet.GetPetType() == "dog"
		if !OddIsDog {
			t.Error("PollDog failed")
			break
		}
	}
	t.Log("PollDog succeed")
}

func TestIsEmpty(t *testing.T) {
	stack := dsa.NewDogCatQueue()
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		if i%2 == 0 {
			stack.Add(dsa.NewCat(id))
		} else {
			stack.Add(dsa.NewDog(id))
		}
	}
	if stack.IsEmpty() || stack.IsCatEmpty() || stack.IsDogEmpty() {
		t.Error("IsEmpty failed")
		return
	}

	for j := 0; j < 10; j += 2 {
		stack.PollCat()
	}
	if !stack.IsCatEmpty() {
		t.Error("IsCatEmpty failed")
		return
	}

	for j := 1; j < 10; j += 2 {
		stack.PollDog()
	}
	if !stack.IsDogEmpty() {
		t.Error("IsDogEmpty failed")
		return
	}

	t.Log("IsEmpty succeed")
}
