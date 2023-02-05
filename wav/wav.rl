package main

import (
  "os"
  "fmt"
)

var eof int
var cs int
var BUF_LEN int = 256
var p int
var pe int
var data []byte

%%{
  machine WAV;
  action match {
    fmt.Println("MATCH")
  }
  header = "RIFF" "WAVE" % match;
  main := header extend;
}%%

%% write data;

func main() {
  %% write init;
  data = make([]byte, BUF_LEN)
  p = 0
  f,err := os.Open("../data/sample.wav")
  if err != nil { panic(err) }
  f.Read(data)
  pe = len(data)
  %% write exec;
}
