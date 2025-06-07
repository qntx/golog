// Package logger_test contains tests for the logger package.
//
// Verifies logger creation, level filtering, message handling, formatting, and JSON output.
package golog_test

import (
	"sync"
	"testing"

	"github.com/qntx/golog"
)

// To run benchmarks:
//   > cd C:\Users\14388\Desktop\qntx\golog
//   > go test -c .
//   > .\golog.test.exe -test.run=NONE -test.bench=. -test.benchmem > $null 2>&1
// Redirect output to $null to avoid terminal printing overhead.

func BenchmarkSimpleInfo(b *testing.B) {
	l := golog.New()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.Info("test")
	}
}

var logString = "String, IDK what to write, let's punch a keyboard. jkdlsjklfdjfklsjfklsdjaflkdjfkdjsfkldjsfkdjklfjdslfjakdfjioerjieofjofdnvonoijdfneslkffjsdfljadljfdjkfjkf"

func BenchmarkFormatInfo(b *testing.B) {
	l := golog.New()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.Infof("test %d %s", i, logString)
	}
}

func BenchmarkFormatInfoMulti(b *testing.B) {
	l := golog.New()
	var wg sync.WaitGroup
	goroutines := 16
	run := func() {
		for i := range b.N / goroutines {
			l.Infof("test %d %s", i, logString)
		}
		wg.Done()
	}
	wg.Add(goroutines)
	b.ResetTimer()
	b.ReportAllocs()
	for range goroutines {
		go run()
	}
	wg.Wait()
}
