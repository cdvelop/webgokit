package main

import (
	"github.com/cdvelop/webgokit/gui"
	// "github.com/cdvelop/webgokit/compiler"
	// "github.com/cdvelop/webgokit/watch"
)

func main() {

	newGui := gui.New()

	// sComp := compiler.New(newGui)

	// watchFiles := watch.New(nil, nil, nil, newGui)

	newGui.Run()

}
