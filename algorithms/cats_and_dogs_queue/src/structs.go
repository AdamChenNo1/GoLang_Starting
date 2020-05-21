package cats_dogs_queue

type Pet interface {
	GetPetType() string
}

type Cat struct {
	types string
	name  string
}

type Dog struct {
	types string
	name  string
}

func (cat Cat) GetPetType() string {
	return cat.types
}

func (dog Dog) GetPetType() string {
	return dog.types
}

func NewCat(name string) Cat {
	return Cat{"cat", name}
}

func NewDog(name string) Dog {
	return Dog{"dog", name}
}
