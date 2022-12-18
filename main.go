//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Matrix Rain 2022
//
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin
// Description: I just wanted to create the matrix rain. Thanks to
// Emily Xie, who was a guest on the youtube channel The Coding Train.
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)

// A symbol is an individual character
type Symbol struct {
  x int
  y int
  speed int
  value int
  first bool
}

// Rain causes the symbol to advance by its speed every "frame"
// in the y direction
func (sym *Symbol) rain(s tcell.Screen) {
  _, y := s.Size()

  if sym.y > y {
    sym.y = 0
  }
  sym.y += sym.speed
}

// Picks a random Katakana character for symbol
func (sym *Symbol) setToRandomSymbol() {
  rand.Seed(time.Now().UnixNano())
  sym.value = 0x30a0 + rand.Intn(96)
}

// Draws the symbol to the screen
func (sym *Symbol) render(s tcell.Screen, style tcell.Style) {
  style2 := tcell.StyleDefault.Foreground(tcell.ColorPaleGreen)
  //style3 := tcell.StyleDefault.Foreground(tcell.ColorTurquoise)
  //style4 := tcell.StyleDefault.Foreground(tcell.ColorPaleTurquoise)
  rand.Seed(time.Now().UnixNano())
  if rand.Intn(10) < 2 {
    sym.setToRandomSymbol()
  }
  if !sym.first {
    writeToScreen(s, style, sym.x, sym.y, string(sym.value))
  } else {
    writeToScreen(s, style2, sym.x, sym.y, string(sym.value))
  }
}

// A stream is an array of symbols
type Stream struct {
  symbols []Symbol
  totalSymbols int
}

// Creates the array of symbols
func (stream *Stream) generateSymbols(x, y, speed int) {
  rand.Seed(time.Now().UnixNano())
  rand := rand.Intn(10)
  first := rand < 3
  for i := 0; i < stream.totalSymbols; i++ {
    sym := Symbol{
      x: x,
      y: y-i,
      speed: speed,
      first: first,
    }
    sym.setToRandomSymbol()
    stream.symbols = append(stream.symbols, sym)
    first = false
  }
}

// Draws a stream to the screen
func (stream *Stream) render(s tcell.Screen, style tcell.Style) {
  for i, sym := range stream.symbols {
    for j := 0; j < len(stream.symbols); j++ {
      if i != j {
        if sym.y == stream.symbols[j].y {
          stream.symbols[i].y -= 1
        }
      }
    }
  }
  for i := 0; i < stream.totalSymbols; i++ {
    stream.symbols[i].render(s, style)
    stream.symbols[i].rain(s)
  }
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

  rand.Seed(time.Now().UnixNano())

  var streams []Stream

  for i := 0; i < x; i++ {
    if i % 3 == 0 {
      stream := Stream{
        totalSymbols: rand.Intn(y-10)+10,
      }
      speed := rand.Intn(2)+1
      //speed := 1
      rany := rand.Intn(50)
      stream.generateSymbols(i, rany*-1, speed)
      streams = append(streams, stream)
    }
  }

  // Main "draw" loop
  for {
    for _, stream := range streams {
      stream.render(s, style)
    }
    s.Sync()
    time.Sleep(time.Millisecond * 60)
    s.Clear()
  }
}
