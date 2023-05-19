package w65816

import "github.com/apoloval/bitwise/logic"

type CPU struct {
	reset logic.Wire[logic.TriStateLevel]
	din   logic.Wire[logic.TriState[uint8]]

	addr logic.TriStateRegister[uint16]
	dout logic.TriStateRegister[uint8]
	rwb  logic.TriStateRegister[logic.Level]

	regs  RegsBank
	state func(logic.ClockEvent)
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) ConnectClock(clk *logic.Clock) {
	clk.Register(cpu)
}

func (cpu *CPU) ConnectReset(w logic.Wire[logic.TriStateLevel]) {
	cpu.reset = w
}

func (cpu *CPU) ConnectDin(w logic.Wire[logic.TriState[uint8]]) {
	cpu.din = w
}

func (cpu *CPU) Addr() logic.Wire[logic.TriState[uint16]] {
	return cpu.addr
}

func (cpu *CPU) Data() logic.Wire[logic.TriState[uint8]] {
	return cpu.dout
}

func (cpu *CPU) OnClockEvent(ev logic.ClockEvent) {
	if val := cpu.reset.Sample(); val.Is(logic.Low) {
		cpu.state = cpu.onReset
	}
	if cpu.state != nil {
		cpu.state(ev)
	}
}

func (cpu *CPU) onReset(ev logic.ClockEvent) {
	switch ev {
	case logic.RisingEdge:
		// Reset internal registers
		cpu.regs.D = 0
		cpu.regs.DBR = 0
		cpu.regs.PBR = 0
		cpu.regs.S = cpu.regs.S&0x0011 | 0x0100             // Set SH to 0x01
		cpu.regs.X = cpu.regs.X & 0x0011                    // Set XH to 0x00
		cpu.regs.Y = cpu.regs.Y & 0x0011                    // Set YH to 0x00
		cpu.regs.P = cpu.regs.P&0b11_0000_11 | 0b00_1101_00 // Set P central bits to 1101
		cpu.regs.PC = 0xFFFC

		// TODO: reset control signals

		cpu.state = cpu.onOpCodeFetch
	}
}

func (cpu *CPU) onOpCodeFetch(ev logic.ClockEvent) {
	switch ev {
	case logic.FallingEdge:
		cpu.addr.Drive(cpu.regs.PC)
		cpu.dout.Drive(cpu.regs.PBR)
		cpu.rwb.Drive(logic.High)
	case logic.RisingEdge:
		cpu.dout.Undrive()
		cpu.rwb.Undrive()
		cpu.state = cpu.onOpCodeDecode
	}
}

func (cpu *CPU) onOpCodeDecode(ev logic.ClockEvent) {
	switch ev {
	case logic.FallingEdge:
		cpu.regs.IR = cpu.din.Sample().Value
	case logic.RisingEdge:
		// TODO: assuming here we will get a NOP (0xEA)
		cpu.execNop()
		cpu.state = cpu.onOpCodeFetch
	}
}

func (cpu *CPU) execNop() {
	cpu.regs.PC++
}
