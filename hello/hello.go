package hello

import (
	"fmt"

	"github.com/arnaubennassar/git-subtree-playground-upstream/internaldependency"
)

func SayHello() {
	fmt.Println("Simple... hello")
}

func SayHelloFromInternalDependency() {
	internaldependency.InternalDependency()
}

func SayHelloToConflicts() {
	fmt.Println("helloc conflicts")
}
