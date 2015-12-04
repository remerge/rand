package rand

import (
	coreRand "math/rand"
	"testing"
	"time"
)

func BenchmarkCoreGlobal(b *testing.B) {
	var acc int64
	for i := 0; i < b.N; i++ {
		acc += coreRand.Int63()
	}
	_ = acc
}

func BenchmarkRandGlobal(b *testing.B) {
	var acc int64
	for i := 0; i < b.N; i++ {
		acc += Int63()
	}
	_ = acc
}

func BenchmarkCore(b *testing.B) {
	generator := coreRand.New(coreRand.NewSource(time.Now().UnixNano()))
	var acc int64
	for i := 0; i < b.N; i++ {
		acc += generator.Int63()
	}
	_ = acc
}

func BenchmarkRand(b *testing.B) {
	generator := New(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
	var acc int64
	for i := 0; i < b.N; i++ {
		acc += generator.Int63()
	}
	_ = acc
}

func BenchmarkCoreEach(b *testing.B) {
	var acc int64
	for i := 0; i < b.N; i++ {
		generator := coreRand.New(coreRand.NewSource(time.Now().UnixNano()))
		acc += generator.Int63()
	}
	_ = acc
}

func BenchmarkRandEach(b *testing.B) {
	var acc int64
	for i := 0; i < b.N; i++ {
		generator := New(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
		acc += generator.Int63()
	}
	_ = acc
}

func BenchmarkCoreGlobalParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var acc int64
		for pb.Next() {
			acc += coreRand.Int63()
		}
		_ = acc
	})
}

func BenchmarkRandGlobalParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var acc int64
		for pb.Next() {
			acc += Int63()
		}
		_ = acc
	})
}

func BenchmarkCoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		generator := coreRand.New(coreRand.NewSource(time.Now().UnixNano()))
		var acc int64
		for pb.Next() {
			acc += generator.Int63()
		}
		_ = acc
	})
}

func BenchmarkRandParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		generator := New(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
		var acc int64
		for pb.Next() {
			acc += generator.Int63()
		}
		_ = acc
	})
}

func BenchmarkCoreEachParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var acc int64
		for pb.Next() {
			generator := coreRand.New(coreRand.NewSource(time.Now().UnixNano()))
			acc += generator.Int63()
		}
		_ = acc
	})
}

func BenchmarkRandEachParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var acc int64
		for pb.Next() {
			generator := New(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
			acc += generator.Int63()
		}
		_ = acc
	})
}
