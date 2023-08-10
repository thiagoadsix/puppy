package puppy

import (
	"github.com/thiagoadsix/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrouUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrouUp(Barks())
}
