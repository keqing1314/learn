// 基准测试
package test

import (
	"learn/test/demo"
	"testing"
)

var longString = "longStringlongStringlongStringlongStringlongStringlongStringlongStringlongString"

func BenchmarkConcatDirect(b *testing.B) {
	for b.Loop() {
		demo.ConcatStringDirect(longString)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for b.Loop() {
		demo.ConcatStringWithBuilder(longString)
	}
}
