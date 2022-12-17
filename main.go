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

func (sym *Symbol) rain(y int) {
  sym.y += 1
  if sym.y > y+1 {
    sym.y = 0
  }
}

func (sym *Symbol) setToRandomSymbol() {
  rand.Seed(time.Now().UnixNano())
  sym.value = 0x30a0 + rand.Intn(96)
}

func (sym *Symbol) render(s tcell.Screen, style tcell.Style, speed int) {
  rand.Seed(time.Now().UnixNano())
  if rand.Intn(10) < 2 {
    sym.setToRandomSymbol()
  }
  writeToScreen(s, style, sym.x, sym.y, string(sym.value))
  s.Sync()
  time.Sleep(time.Millisecond * time.Duration(speed))
  s.Clear()
}

// This is used just to write strings to the screen.
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), nil, style)
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

  style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

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

  var sym Symbol
  sym = Symbol{
    x: x/2,
    y: y-y,
  }

  sym.setToRandomSymbol()

  for {
    sym.render(s, style, 120)
    sym.rain(y)

  }
}
