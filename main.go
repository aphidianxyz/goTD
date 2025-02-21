package main

import (
    screen "github.com/aphidianxyz/goTD/engine/screen"
    rlw "github.com/aphidianxyz/goTD/engine/raylibwrapper"
)

func main() {
    rlw := &rlw.RealRaylibWrapper{}
    screen := screen.Screen{Width: 100, Height: 100, Title: "game"}
    screen.Init(rlw)
}

