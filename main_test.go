package main

import (
	"testing"
)


// Функция теста производительности для функции factorial
func BenchmarkGenerateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
        generateRandomString(10)
	}
}

// Функция теста производительности для функции factorial
func BenchmarkGenerateRandomString2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        generateRandomString2(5)
    }
}
