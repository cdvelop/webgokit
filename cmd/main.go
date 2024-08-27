package main

import (
	"github.com/cdvelop/webgokit/browser"
	"github.com/cdvelop/webgokit/compiler"
	"github.com/cdvelop/webgokit/console"
	"github.com/cdvelop/webgokit/gui"
	"github.com/cdvelop/webgokit/server"
	// "github.com/cdvelop/webgokit/watch"
)

func main() {

	cmd := console.New()

	comp := compiler.New(cmd)

	bwr := browser.New()

	svr := server.New()

	newGui := gui.New(cmd, comp, bwr, svr)

	// sComp := compiler.New(newGui)

	// watchFiles := watch.New(nil, nil, nil, newGui)

	newGui.Run()

}
