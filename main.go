package main

import (
  "os"
)

func main() {
  audio := Audio{
    Data: []byte{},
  }
  for i := 0; i < 1<<20; i++ {
  }
  wav := audio.ToWAV()
  f,err := os.Create("data/out.wav")
  if err != nil { panic(err) }
  wav.Write(f)
}
