package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/cdvelop/webgokit/config"
)

type compilerAdapter interface {
	InputDirectoryChange(newDirectory string)
	OutputDirectoryChange(newDirectory string)
}

type browserAdapter interface {
	BrowserClose() error
}

type serverAdapter interface {
	ServerClose() error
}

type consoleAdapter interface {
	Console() fyne.CanvasObject
	Info(message string)
	Success(message string)
	Error(error)
}

type watcherAdapter interface {
	FileWatcherStart()
	RegisterFiles(folderToWatch string)
}

type handler struct {
	conf *config.Handler

	window fyne.Window
	consoleAdapter

	watcher  watcherAdapter
	compiler compilerAdapter
	browser  browserAdapter
	server   serverAdapter
}

func New(conf *config.Handler, cons consoleAdapter, w watcherAdapter, c compilerAdapter, b browserAdapter, s serverAdapter) *handler {

	a := app.NewWithID("com.webgokit.cdvelop.github")
	a.Settings().SetTheme(newFysionTheme())

	h := &handler{
		conf:           conf,
		window:         a.NewWindow("WebGoKit"),
		consoleAdapter: cons,
		watcher:        w,
		compiler:       c,
		browser:        b,
		server:         s,
	}

	return h
}
