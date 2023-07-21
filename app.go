package main

import (
	"context"
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"

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

func addEndBounceHook() {
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("esc")
		isBouncing = false
		hook.End()
	})

	s := hook.Start()
	hook.Process(s)
}

var isBouncing = false

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(ctx, "key_pressed", func(data ...interface{}) {
		fmt.Println("key pressed", data)
		// robotgo.TypeStr("Hello Worldzzz")
		isBouncing = true
		dirX, dirY := 1, 1
		sizeX, sizeY := robotgo.GetScreenSize()
		addEndBounceHook()
		// x, y := robotgo.GetMousePos()
		for {
			x, y := robotgo.GetMousePos()
			if dirX == 1 && x+dirX >= sizeX {
				dirX = -1
			} else if x+dirX <= 0 {
				dirX = 1
			}

			if dirY == 1 && y+dirY >= sizeY {
				dirY = -1
			} else if y+dirY <= 0 {
				dirY = 1
			}
			robotgo.MoveSmooth(x+dirX, y+dirY)
			if !isBouncing {
				fmt.Println("stop bounce")
				break
			}
		}
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

func (a *App) BounceMouse() {
	x, y := robotgo.GetMousePos()
	robotgo.MoveSmoothRelative(300, 300, 1.0, 10.0)
	fmt.Println("pos:", x, y)
}
