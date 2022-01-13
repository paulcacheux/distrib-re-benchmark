package bench_init

import (
	_ "github.com/distribution/distribution/v3/reference"
    "testing"
	_ "unsafe"
)

//go:linkname _initdone github.com/distribution/distribution/v3/reference.initdone·
var _initdone uint8

//go:linkname _init github.com/distribution/distribution/v3/reference.init
func _init()

func BenchmarkPR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_initdone = 0
		_init()
	}
}
