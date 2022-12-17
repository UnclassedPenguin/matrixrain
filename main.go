package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)

type Symbol struct {
  x int
  y int
  value int
}

func (s *Symbol) setToRandomSymbol() {
  rand.Seed(time.Now().UnixNano())
  //s.value = 0x30a0 + rand.Intn(96)
  s.value = 12448 + rand.Intn(97)
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

  //encoding.Register()

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }

  s.Clear()

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

  x, y := s.Size()

  go func() {
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
  }()

  for {
    var sym Symbol

    sym.setToRandomSymbol()

    writeToScreen(s, style, x/2-4, y/2+1, string(sym.value))

    time.Sleep(time.Millisecond * 100)
  }

}
