package main
import (
  "io"
  "encoding/binary"
)

const PCM int = 1

type Audio struct {
  Data []byte
}

type RiffHeader struct {
  chunkID       string
  chunkSize     int
  format        string
}
type FmtChunk struct {
  chunkID       string
  chunkSize     int
  audioFormat   int
  numChannels   int
  sampleRate    int
  byteRate      int
  blockAlign    int
  bitsPerSample int
}
type DataChunk struct {
  chunkID       string
  chunkSize     int
  data          []byte
}

type WAV struct {
  riffHeader  RiffHeader
  fmtChunk    FmtChunk
  dataChunk   DataChunk
}

func (audio Audio) ToWAV() WAV {
  sampleRate := 8000
  numChannels := 1
  bitsPerSample := 8
  return WAV{
    riffHeader: RiffHeader{
      chunkID: "RIFF",
      chunkSize: 36 + len(audio.Data),
      format: "WAVE",
    },
    fmtChunk: FmtChunk{
      chunkID: "fmt ",
      chunkSize: 16,
      audioFormat: PCM,
      numChannels: 2,
      sampleRate: sampleRate,
      byteRate: sampleRate * numChannels * bitsPerSample/8,
      blockAlign: numChannels * bitsPerSample/8,
      bitsPerSample: bitsPerSample,
    },
    dataChunk: DataChunk{
      chunkID: "data",
      chunkSize: len(audio.Data),
      data: audio.Data,
    },
  }
}

func (x RiffHeader) Write(f io.Writer) {
  b32 := make([]byte, 4)
  f.Write([]byte(x.chunkID))
  binary.LittleEndian.PutUint32(b32, uint32(x.chunkSize))
  f.Write(b32)
  f.Write([]byte(x.format))
}

func (x FmtChunk) Write(f io.Writer) {
  b32 := make([]byte, 4)
  b16 := make([]byte, 2)
  f.Write([]byte(x.chunkID))
  binary.LittleEndian.PutUint32(b32, uint32(x.chunkSize))
  f.Write(b32)
  binary.LittleEndian.PutUint16(b16, uint16(x.audioFormat))
  f.Write(b16)
  binary.LittleEndian.PutUint16(b16, uint16(x.numChannels))
  f.Write(b16)
  binary.LittleEndian.PutUint16(b32, uint16(x.sampleRate))
  f.Write(b32)
  binary.LittleEndian.PutUint16(b32, uint16(x.byteRate))
  f.Write(b32)
  binary.LittleEndian.PutUint16(b16, uint16(x.blockAlign))
  f.Write(b16)
  binary.LittleEndian.PutUint16(b16, uint16(x.bitsPerSample))
  f.Write(b16)
}

func (x DataChunk) Write(f io.Writer) {
  b32 := make([]byte, 4)
  f.Write([]byte(x.chunkID))
  binary.LittleEndian.PutUint32(b32, uint32(x.chunkSize))
  f.Write(b32)
  f.Write(x.data)
}

func (x WAV) Write(f io.Writer) {
  x.riffHeader.Write(f)
  x.fmtChunk.Write(f)
  x.dataChunk.Write(f)
}
