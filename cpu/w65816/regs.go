package w65816

type RegsBank struct {
	P   uint8  // P is the Processor Status
	A   uint16 // A is the Accumulator
	X   uint16 // X is the index register X
	Y   uint16 // X is the index register Y
	PC  uint16 // PC is the program counter
	D   uint16 // D is the Direct Register
	S   uint16 // S is the Stack Pointer
	DBR uint8  // DBR is the Data Bank Register
	PBR uint8  // PBR is the Program Bank Register

	IR uint8 // IR is the Instruction Register
}
