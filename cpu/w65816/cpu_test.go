package w65816_test

import (
	"testing"

	"github.com/apoloval/bitwise/cpu/w65816"
	"github.com/apoloval/bitwise/logic"
)

func BenchmarkCPU(b *testing.B) {
	var clock logic.Clock
	var rst logic.TriStateRegister[logic.Level]

	cpu := w65816.New()
	cpu.ConnectReset(&rst)
	cpu.ConnectDin(logic.Fixed(logic.Value[uint8](0xEA)))

	clock.Register(&rst)
	clock.Register(cpu)

	rst.Drive(logic.Low)
	clock.Step(2)
	rst.Drive(logic.High)

	b.ResetTimer()
	clock.Step(b.N)
}
