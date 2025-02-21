package raylibwrapper

import rl "github.com/gen2brain/raylib-go/raylib"

type RaylibWrapper interface {
    IsWindowReady() bool
    InitWindow(width, height int32, title string)
    CloseWindow()
    GetScreenWidth() int
    GetScreenHeight() int
    GetScreenTitle() string
}

type RealRaylibWrapper struct {
    title string
}

func (r RealRaylibWrapper) IsWindowReady() bool {
    return rl.IsWindowReady()
}

func (r RealRaylibWrapper) InitWindow(width, height int32, title string) {
    rl.InitWindow(width, height, title)
}

func (r RealRaylibWrapper) CloseWindow() {
    rl.CloseWindow()
}

func (r RealRaylibWrapper) GetScreenWidth() int {
    return rl.GetScreenWidth()
}

func (r RealRaylibWrapper) GetScreenHeight() int {
    return rl.GetScreenHeight()
}

func (r RealRaylibWrapper) GetScreenTitle() string {
    return r.title
}
