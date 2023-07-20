package main

import (
	"context"
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(ctx, "key_pressed", func(data ...interface{}) {
		fmt.Println("key pressed", data)
		robotgo.TypeStr("Hello World")

		// x, y := robotgo.GetMousePos()
		// robotgo.MoveSmoothRelative(300, 300, 10.0, 10.0)
		// fmt.Println("pos:", x, y)
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Hold down the pressed key
func (a *App) HoldKey(key string) bool {
	key_return := robotgo.KeyToggle("a")
	fmt.Println("key", key_return)
	return true
}
