package screen

import (
	"errors"

	rl "github.com/aphidianxyz/goTD/engine/raylibwrapper"
)

const ErrInitMultipleScreens = "cannot create more than 1 screen"

type Screen struct {
    // utilizing type safety to avoid failed assertions with rl when 
    // a negative dimension is defined for screen
    Width, Height   uint16  
    Title           string
}

func (s Screen) Init(rl rl.RaylibWrapper) error {
    if rl.IsWindowReady() {
        return errors.New(ErrInitMultipleScreens)
    }
    rl.InitWindow(int32(s.Width), int32(s.Height), s.Title)
    return nil
}

func (s Screen) Close(rl rl.RaylibWrapper) error {
    if !rl.IsWindowReady() {
        return errors.New("No screen to close!")
    }

    rl.CloseWindow()
    return nil
}
