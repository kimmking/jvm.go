package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke interface method
type InvokeInterface struct {
	index uint
	// count uint8
	// zero uint8

	// optimization
	kMethodRef   *heap.ConstantInterfaceMethodref
	argSlotCount uint
}

func (instr *InvokeInterface) FetchOperands(reader *base.BytecodeReader) {
	instr.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (instr *InvokeInterface) Execute(frame *rtda.Frame) {
	if instr.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		instr.kMethodRef = cp.GetConstant(instr.index).(*heap.ConstantInterfaceMethodref)
		instr.argSlotCount = instr.kMethodRef.ArgSlotCount()
	}

	ref := frame.TopRef(instr.argSlotCount)
	if ref == nil {
		panic("NPE") // todo
	}

	method := instr.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread().InvokeMethod(method)
}
