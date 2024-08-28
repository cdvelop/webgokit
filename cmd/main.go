package main

import (
	"github.com/cdvelop/webgokit/browser"
	"github.com/cdvelop/webgokit/compiler"
	"github.com/cdvelop/webgokit/console"
	"github.com/cdvelop/webgokit/gui"
	"github.com/cdvelop/webgokit/server"
	"github.com/cdvelop/webgokit/watch"
)

func main() {

	cmd := console.New()

	comp := compiler.New(cmd)

	bwr := browser.New()

	svr := server.New()

	w := watch.New(comp, nil, nil, cmd)

	newGui := gui.New(cmd, w, comp, bwr, svr)

	// sComp := compiler.New(newGui)

	newGui.Run()

}
