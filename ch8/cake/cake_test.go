package cake_test

import (
	"github.com/prince1809/go-programming/ch8/cake"
	"testing"
	"time"
)

var defaults = cake.Shop{
	Verbose:      testing.Verbose(),
	Cakes:        5,
	BakeTime:     1 * time.Second,
	NumIcers:     1,
	IceTime:      1 * time.Second,
	InscribeTime: 1 * time.Second,
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber.
	// Each step takes exactly 10ms. No buffers
	cakeshop := defaults
	cakeshop.Verbose = true
	cakeshop.Work(2)
}
