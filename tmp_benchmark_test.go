package main

import (
	"errors"
	"log"
	"testing"
)

func Divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}

	return a / b, nil
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(5, 4)
	}
}

func Benchmark_Time(b *testing.B) {
	b.StopTimer()
	log.Println("我是这里的准备阶段")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Divide(5, 4)
	}
}
