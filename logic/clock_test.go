package logic_test

import (
	"fmt"
	"testing"

	"github.com/apoloval/bitwise/logic"
)

func BenchmarkCPU(b *testing.B) {
	for _, size := range []int{2, 5, 10, 20, 50} {
		b.Run(fmt.Sprintf("with %d observers", size), func(b *testing.B) {
			var clock logic.Clock
			regs := make([]logic.TriStateRegister[logic.Level], size)
			for _, r := range regs {
				clock.Register(&r)
			}

			b.ResetTimer()
			clock.Step(b.N)
		})
	}
}
