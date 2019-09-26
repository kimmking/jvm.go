package misc

import (
	"github.com/zxh0/jvm.go/jutil/bigendian"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_unsafe(allocateMemory, "allocateMemory", "(J)J")
	_unsafe(reallocateMemory, "reallocateMemory", "(JJ)J")
	_unsafe(freeMemory, "freeMemory", "(J)V")
	_unsafe(addressSize, "addressSize", "()I")
	_unsafe(putAddress, "putAddress", "(JJ)V")
	_unsafe(getAddress, "getAddress", "(J)J")
	_unsafe(mem_putByte, "putByte", "(JB)V")
	_unsafe(mem_getByte, "getByte", "(J)B")
	_unsafe(mem_putShort, "putShort", "(JS)V")
	_unsafe(mem_getShort, "getShort", "(J)S")
	_unsafe(mem_putChar, "putChar", "(JC)V")
	_unsafe(mem_getChar, "getChar", "(J)C")
	_unsafe(mem_putInt, "putInt", "(JI)V")
	_unsafe(mem_getInt, "getInt", "(J)I")
	_unsafe(mem_putLong, "putLong", "(JJ)V")
	_unsafe(mem_getLong, "getLong", "(J)J")
	_unsafe(mem_putFloat, "putFloat", "(JF)V")
	_unsafe(mem_getFloat, "getFloat", "(J)F")
	_unsafe(mem_putDouble, "putDouble", "(JD)V")
	_unsafe(mem_getDouble, "getDouble", "(J)D")
}

// public native long allocateMemory(long bytes);
// (J)J
func allocateMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	bytes := frame.GetLongVar(1)

	address := allocate(bytes)
	frame.PushLong(address)
}

// public native long reallocateMemory(long address, long bytes);
// (JJ)J
func reallocateMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	bytes := frame.GetLongVar(3)

	newAddress := reallocate(address, bytes)
	frame.PushLong(newAddress)
}

// public native void freeMemory(long address);
// (J)V
func freeMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	free(address)
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
	// vars := frame.
	// frame.GetRefVar(0) // this

	frame.PushInt(8) // todo unsafe.Sizeof(int)
}

// public native void putAddress(long address, long x);
// (JJ)V
func putAddress(frame *rtda.Frame) {
	mem_putLong(frame)
}

// public native long getAddress(long address);
// (J)J
func getAddress(frame *rtda.Frame) {
	mem_getLong(frame)
}

// public native void putByte(long address, byte x);
// (JB)V
func mem_putByte(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutInt8(mem, int8(value.IntValue()))
}

// public native byte getByte(long address);
// (J)B
func mem_getByte(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushInt(int32(bigendian.Int8(mem)))
}

// public native void putShort(long address, short x);
// (JS)V
func mem_putShort(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutInt16(mem, int16(value.IntValue()))
}

// public native short getShort(long address);
// (J)S
func mem_getShort(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushInt(int32(bigendian.Int16(mem)))
}

// public native void putChar(long address, char x);
// (JC)V
func mem_putChar(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutUint16(mem, uint16(value.IntValue()))
}

// public native char getChar(long address);
// (J)C
func mem_getChar(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushInt(int32(bigendian.Uint16(mem)))
}

// public native void putInt(long address, int x);
// (JI)V
func mem_putInt(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutInt32(mem, value.IntValue())
}

// public native int getInt(long address);
// (J)I
func mem_getInt(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushInt(bigendian.Int32(mem))
}

// public native void putLong(long address, long x);
// (JJ)V
func mem_putLong(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutInt64(mem, value.LongValue())
}

// public native long getLong(long address);
// (J)J
func mem_getLong(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushLong(bigendian.Int64(mem))
}

// public native void putFloat(long address, float x);
// (JJ)V
func mem_putFloat(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutFloat32(mem, value.FloatValue())
}

// public native float getFloat(long address);
// (J)J
func mem_getFloat(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushFloat(bigendian.Float32(mem))
}

// public native void putDouble(long address, double x);
// (JJ)V
func mem_putDouble(frame *rtda.Frame) {
	mem, value := _put(frame)
	bigendian.PutFloat64(mem, value.DoubleValue())
}

// public native double getDouble(long address);
// (J)J
func mem_getDouble(frame *rtda.Frame) {
	mem := _get(frame)
	frame.PushDouble(bigendian.Float64(mem))
}

func _put(frame *rtda.Frame) ([]byte, heap.Slot) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	value := frame.GetLocalVar(3)

	mem := memoryAt(address)
	return mem, value
}

func _get(frame *rtda.Frame) []byte {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	return memoryAt(address)
}
