package hello

import (
	"fmt"

	"github.com/arnaubennassar/git-subtree-playground-fork/forkdep"
	"github.com/arnaubennassar/git-subtree-playground-upstream/internaldependency"
)

func SayHello() {
	fmt.Println("Simple... hello")
}

func SayHelloFromInternalDependency() {
	internaldependency.InternalDependency()
}

func SayHelloFromForkDep() {
	forkdep.SayHelloFromForkDep()
}

func SayHelloToConflicts() {
	fmt.Println("helloc conflicts")
}
