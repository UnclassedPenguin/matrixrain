package main

import (
  "os"
  "fmt"
  "github.com/gdamore/tcell/v2"
)

type Symbol struct {
  x int
  y int
  value rune
}

func (s Symbol) setToRandomSymbol() {

}

// This is used just to write strings to the screen.
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

func main() {
  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("Error in tcell.NewScreen:", err)
  }

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }

  s.Clear()

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

  x, y := s.Size()

  writeToScreen(s, style, x/2-4, y/2, "Welcome")

  b := []byte{0x30, 0xA0}

  writeToScreen(s, style, x/2-4, y/2+1, string(b))

  for {
  switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          os.Exit(0)
        }
      }
    }
  }
}
