package hello

import (
	"fmt"

	"github.com/arnaubennassar/git-subtree-playground-upstream/internaldependency"
)

func SayHello() {
	fmt.Println("hello")
}

func SayHelloFromInternalDependency() {
	internaldependency.InternalDependency()
}
