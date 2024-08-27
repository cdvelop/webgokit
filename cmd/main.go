package main

import (
	"log"

	"github.com/cdvelop/webgokit/browser"
	"github.com/cdvelop/webgokit/gui"
	"github.com/cdvelop/webgokit/server"
	// "github.com/cdvelop/webgokit/compiler"
	// "github.com/cdvelop/webgokit/watch"
)

func main() {

	bwr := browser.New()

	svr := server.New()

	newGui := gui.New(bwr, svr)

	log.SetOutput(newGui)

	// sComp := compiler.New(newGui)

	// watchFiles := watch.New(nil, nil, nil, newGui)

	newGui.Run()

}
