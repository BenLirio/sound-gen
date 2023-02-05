
//line wav.rl:1
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


//line wav.rl:22



//line wav.go:23
var _WAV_actions []byte = []byte{
	0, 1, 0, 
}

var _WAV_key_offsets []byte = []byte{
	0, 0, 1, 2, 3, 4, 4, 4, 
	4, 4, 5, 6, 7, 8, 8, 
}

var _WAV_trans_keys []byte = []byte{
	82, 73, 70, 70, 87, 65, 86, 69, 
	
}

var _WAV_single_lengths []byte = []byte{
	0, 1, 1, 1, 1, 0, 0, 0, 
	0, 1, 1, 1, 1, 0, 0, 
}

var _WAV_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _WAV_index_offsets []byte = []byte{
	0, 0, 2, 4, 6, 8, 9, 10, 
	11, 12, 14, 16, 18, 20, 21, 
}

var _WAV_trans_targs []byte = []byte{
	2, 0, 3, 0, 4, 0, 5, 0, 
	6, 7, 8, 9, 10, 0, 11, 0, 
	12, 0, 13, 0, 14, 0, 
}

var _WAV_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 0, 
}

const WAV_start int = 1
const WAV_first_final int = 14
const WAV_error int = 0

const WAV_en_main int = 1


//line wav.rl:25

func main() {
  
//line wav.go:76
	{
	cs = WAV_start
	}

//line wav.rl:28
  data = make([]byte, BUF_LEN)
  p = 0
  f,err := os.Open("../data/sample.wav")
  if err != nil { panic(err) }
  f.Read(data)
  pe = len(data)
  
//line wav.go:89
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_WAV_key_offsets[cs])
	_trans = int(_WAV_index_offsets[cs])

	_klen = int(_WAV_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _WAV_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _WAV_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_WAV_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _WAV_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _WAV_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_WAV_trans_targs[_trans])

	if _WAV_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_WAV_trans_actions[_trans])
	_nacts = uint(_WAV_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _WAV_actions[_acts-1] {
		case 0:
//line wav.rl:17

    fmt.Println("MATCH")
  
//line wav.go:172
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line wav.rl:35
}
