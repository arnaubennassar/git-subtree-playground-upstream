package main

import (
	"github.com/arnaubennassar/git-subtree-playground-upstream/hello"
	"github.com/arnaubennassar/git-subtree-playground-upstream/world"
)

func main() {
	hello.SayHello()
	world.World()
	hello.SayHelloFromInternalDependency()
	hello.SayHelloFromForkDep()
}
