package hello

import (
	"fmt"

	"github.com/arnaubennassar/git-subtree-playground-upstream/internaldependency"
)

func SayHello() {
	fmt.Println("BRAND NEW HELLO FUNC, but still... hello")
}

func SayHelloFromInternalDependency() {
	internaldependency.InternalDependency()
}
