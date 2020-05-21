package main

import (
	dsa "algorithms/cats_and_dogs_queue/src"
	"fmt"
)

func main() {
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
		modi := j % 2
		pet := stack.PollPet()
		fmt.Println(modi, pet.GetPetType())
	}
}
