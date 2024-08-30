package main

import (
	"github.com/cdvelop/webgokit/browser"
	"github.com/cdvelop/webgokit/compiler"
	"github.com/cdvelop/webgokit/config"
	"github.com/cdvelop/webgokit/console"
	"github.com/cdvelop/webgokit/gui"
	"github.com/cdvelop/webgokit/server"
	"github.com/cdvelop/webgokit/watch"
)

func main() {
	cmd := console.New()

	conf, err := config.New(cmd)
	if err != nil {
		panic(err)
	}

	comp := compiler.New(cmd)

	bwr := browser.New()

	svr := server.New()

	w := watch.New(comp, nil, nil, cmd)

	newGui := gui.New(conf, cmd, w, comp, bwr, svr)

	// sComp := compiler.New(newGui)

	newGui.Run()

}
