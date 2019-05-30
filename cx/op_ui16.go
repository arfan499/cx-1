package cxcore

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// The built-in str function returns the base 10 string representation of operand 1.
func opUI16ToStr(expr *CXExpression, fp int) {
	outB0 := FromStr(strconv.FormatUint(uint64(ReadUI16(fp, expr.Inputs[0])), 10))
	WriteObject(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in i8 function returns operand 1 casted from type ui16 to type i8.
func opUI16ToI8(expr *CXExpression, fp int) {
	outB0 := FromI16(int16(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in i16 function returns operand 1 casted from type ui16 to type i16.
func opUI16ToI16(expr *CXExpression, fp int) {
	outB0 := FromI16(int16(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in i32 function return operand 1 casted from type ui16 to type i32.
func opUI16ToI32(expr *CXExpression, fp int) {
	outB0 := FromI32(int32(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in i64 function returns operand 1 casted from type ui16 to type i64.
func opUI16ToI64(expr *CXExpression, fp int) {
	outB0 := FromI64(int64(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in ui8 function returns operand 1 casted from type ui16 to type ui8.
func opUI16ToUI8(expr *CXExpression, fp int) {
	outB0 := FromUI8(uint8(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in ui32 function returns the operand 1 casted from type ui16 to type ui32.
func opUI16ToUI32(expr *CXExpression, fp int) {
	outB0 := FromUI32(uint32(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in ui64 function returns the operand 1 casted from type ui16 to type ui64.
func opUI16ToUI64(expr *CXExpression, fp int) {
	outB0 := FromUI64(uint64(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in f32 function returns operand 1 casted from type ui16 to type f32.
func opUI16ToF32(expr *CXExpression, fp int) {
	outB0 := FromF32(float32(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in f64 function returns operand 1 casted from type ui16 to type f64.
func opUI16ToF64(expr *CXExpression, fp int) {
	outB0 := FromF64(float64(ReadUI16(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The print built-in function formats its arguments and prints them.
func opUI16Print(expr *CXExpression, fp int) {
	fmt.Println(ReadUI16(fp, expr.Inputs[0]))
}

// The built-in add function returns the sum of two ui16 numbers.
func opUI16Add(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) + ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in sub function returns the difference of two ui16 numbers.
func opUI16Sub(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) - ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in mul function returns the product of two ui16 numbers.
func opUI16Mul(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) * ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in div function returns the quotient of two ui16 numbers.
func opUI16Div(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) / ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in gt function returns true if operand 1 is greater than operand 2.
func opUI16Gt(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) > ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in gteq function returns true if operand 1 is greater than or
// equal to operand 2.
func opUI16Gteq(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) >= ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in lt function returns true if operand 1 is less than operand 2.
func opUI16Lt(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) < ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in lteq function returns true if operand 1 is less than or equal
// to operand 1.
func opUI16Lteq(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) <= ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in eq function returns true if operand 1 is equal to operand 2.
func opUI16Eq(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) == ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in uneq function returns true if operand 1 is different from operand 2.
func opUI16Uneq(expr *CXExpression, fp int) {
	outB0 := FromBool(ReadUI16(fp, expr.Inputs[0]) != ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in mod function returns the remainder of operand 1 / operand 2.
func opUI16Mod(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) % ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in rand function returns a pseudo-random number.
func opUI16Rand(expr *CXExpression, fp int) {
	outB0 := FromUI16(uint16(rand.Int31n(int32(math.MaxUint16))))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitand function returns the bitwise AND of 2 operands.
func opUI16Bitand(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) & ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitor function returns the bitwise OR of 2 operands.
func opUI16Bitor(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) | ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitxor function returns the bitwise XOR of 2 operands.
func opUI16Bitxor(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) ^ ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitclear function returns the bitwise AND NOT of 2 operands.
func opUI16Bitclear(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) &^ ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitshl function returns bits of operand 1 shifted to the left
// by number of positions specified in operand 2.
func opUI16Bitshl(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) << ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in bitshr function returns bits of operand 1 shifted to the right
// by number of positions specified in operand 2.
func opUI16Bitshr(expr *CXExpression, fp int) {
	outB0 := FromUI16(ReadUI16(fp, expr.Inputs[0]) >> ReadUI16(fp, expr.Inputs[1]))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The built-in max function returns the biggest of the two operands..
func opUI16Max(expr *CXExpression, fp int) {
	inp0 := ReadUI16(fp, expr.Inputs[0])
	inp1 := ReadUI16(fp, expr.Inputs[1])
	if inp1 > inp0 {
		inp0 = inp1
	}
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), FromUI16(inp0))
}

// The built-in max function returns the biggest of the two operands.
func opUI16Min(expr *CXExpression, fp int) {
	inp0 := ReadUI16(fp, expr.Inputs[0])
	inp1 := ReadUI16(fp, expr.Inputs[1])
	if inp1 < inp0 {
		inp0 = inp1
	}
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), FromUI16(inp0))
}