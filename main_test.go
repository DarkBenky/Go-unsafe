package main_test

import (
	"math/rand"
	"testing"
	"unsafe"
)

type Vector struct {
	X, Y, Z float32
}

// Define a struct
type MyStruct struct {
	V1 , V2, V3 Vector
	R, G, B , A float32
}

const arraySize = 1_000_000

var structArray []MyStruct
var indexes []int

func init() {
	// Initialize the array with random data
	structArray = make([]MyStruct, arraySize)
	for i := 0; i < arraySize; i++ {
		structArray[i] = MyStruct{
			V1: Vector{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()},
			V2: Vector{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()},
			V3: Vector{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()},
			R:  rand.Float32(),
			G:  rand.Float32(),
			B:  rand.Float32(),
			A:  rand.Float32(),
		}
	}

	// Pre-generate random indexes
	indexes = make([]int, 100_000)
	for i := range indexes {
		indexes[i] = rand.Intn(arraySize)
	}
}

func BenchmarkStandardIndexing(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, idx := range indexes {
			_ = structArray[idx]
		}
	}
}

func BenchmarkUnsafePointerIndexing(b *testing.B) {
    b.ResetTimer()
    size := unsafe.Sizeof(MyStruct{})
    basePtr := uintptr(unsafe.Pointer(&structArray[0]))
    for i := 0; i < b.N; i++ {
        for _, idx := range indexes {
            _ = *(*MyStruct)(unsafe.Pointer(basePtr + uintptr(idx)*size))
        }
    }
}
