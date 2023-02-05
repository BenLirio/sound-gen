package main

import (
  "os"
)

func main() {
  audio := Audio{
    Samples: []Sample{},
  }
  d := 1
  v := 0
  for i := 0; i < 1<<20; i++ {
    v += d
    if v == 0 { d = 1 }
    if v == 100 { d = -1 }
    audio.Samples = append(audio.Samples, Sample(v))
  }
  audio = audio.GenData()
  wav := audio.ToWAV()
  f,err := os.Create("data/out.wav")
  if err != nil { panic(err) }
  wav.Write(f)
}
