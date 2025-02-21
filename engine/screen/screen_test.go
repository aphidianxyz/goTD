package screen

import (
    "testing"
)

type spyRaylibWrapper struct {
    screenWidth, screenHeight int
    screenTitle string
    windowReady, windowClosed bool
}

func (r *spyRaylibWrapper) IsWindowReady() bool {
    return r.windowReady
}

func (r *spyRaylibWrapper) InitWindow(width, height int32, title string) {
    r.screenWidth = int(width)
    r.screenHeight = int(height)
    r.screenTitle = title
    r.windowReady = true
}

func (r *spyRaylibWrapper) CloseWindow() {
    r.windowClosed = true
}

func (r spyRaylibWrapper) GetScreenWidth() int {
    return r.screenWidth
}

func (r spyRaylibWrapper) GetScreenHeight() int {
    return r.screenHeight
}

func (r spyRaylibWrapper) GetScreenTitle() string {
    return r.screenTitle
}

func TestScreenInit(t *testing.T) {
    t.Run("basic screen init", func(t *testing.T) {
        rlw := spyRaylibWrapper{}
        screen := Screen{Width: 100, Height: 100, Title: "Game"}
        err := screen.Init(&rlw)

        got := screen
        want := Screen{uint16(rlw.screenWidth), uint16(rlw.screenHeight), rlw.screenTitle}

        assertNoErrors(t, err)
        if got != want {
            t.Errorf("got %v, expected %v", got, want)
        }
    })

    t.Run("fail on creating 2 screens", func(t *testing.T) {
        rlw := spyRaylibWrapper{}
        screen := Screen{Width: 100, Height: 100, Title: "Game1"}
        screen2 := Screen{Width: 1000, Height: 1000, Title: "Game2"}

        screen.Init(&rlw)
        err := screen2.Init(&rlw)

        if err.Error() != ErrInitMultipleScreens {
            t.Errorf("expected an error: %v, got %v", ErrInitMultipleScreens, err.Error())
        }
    })

    t.Run("close window", func(t *testing.T) {
        rlw := spyRaylibWrapper{}
        screen := Screen{Width: 100, Height: 100}

        err := screen.Init(&rlw)
        err2 := screen.Close(&rlw)

        assertNoErrors(t, err)
        assertNoErrors(t, err2)
        if !rlw.windowClosed {
            t.Errorf("Expected window to close")
        }
    })

    t.Run("attempting to close a window that's not initialized", func(t *testing.T) {
        rlw := spyRaylibWrapper{}
        screen := Screen{Width: 100, Height: 100, Title: "Game"}

        err := screen.Close(&rlw)

        if err == nil {
            t.Errorf("Expected error: '%v', got nothing", "No screen to close!")
        }
    })
}

func assertNoErrors(t *testing.T, got error) {
    t.Helper()
    if got != nil {
        t.Errorf("Expected no errors, got '%v'", got)
    }
}
