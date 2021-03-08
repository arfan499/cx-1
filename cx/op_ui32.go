package cxcore

import (
	"fmt"
	"math/rand"
	"strconv"
)

// The built-in str function returns the base 10 string representation of operand 1.
func opUI32ToStr(expr *CXExpression, fp int) {
	outB0 := FromStr(strconv.FormatUint(uint64(ReadUI32(fp, expr.Inputs[0])), 10))
	WriteObject(GetOffsetUI32(fp, expr.Outputs[0]), outB0)
}

// The built-in i8 function returns operand 1 casted from type ui32 to type i8.
func opUI32ToI8(expr *CXExpression, fp int) {
	outV0 := int8(ReadUI32(fp, expr.Inputs[0]))
	WriteI8(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in i16 function returns operand 1 casted from type ui32 to type i16.
func opUI32ToI16(expr *CXExpression, fp int) {
	outV0 := int16(ReadUI32(fp, expr.Inputs[0]))
	WriteI16(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in i32 function return operand 1 casted from type ui32 to type i32.
func opUI32ToI32(expr *CXExpression, fp int) {
	outV0 := int32(ReadUI32(fp, expr.Inputs[0]))
	WriteI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in i64 function returns operand 1 casted from type ui32 to type i64.
func opUI32ToI64(expr *CXExpression, fp int) {
	outV0 := int64(ReadUI32(fp, expr.Inputs[0]))
	WriteI64(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in ui8 function returns operand 1 casted from type ui32 to type ui8.
func opUI32ToUI8(expr *CXExpression, fp int) {
	outV0 := uint8(ReadUI32(fp, expr.Inputs[0]))
	WriteUI8(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in ui16 function returns the operand 1 casted from type ui32 to type ui16.
func opUI32ToUI16(expr *CXExpression, fp int) {
	outV0 := uint16(ReadUI32(fp, expr.Inputs[0]))
	WriteUI16(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in ui64 function returns the operand 1 casted from type ui32 to type ui64.
func opUI32ToUI64(expr *CXExpression, fp int) {
	outV0 := uint64(ReadUI32(fp, expr.Inputs[0]))
	WriteUI64(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in f32 function returns operand 1 casted from type ui32 to type f32.
func opUI32ToF32(expr *CXExpression, fp int) {
	outV0 := float32(ReadUI32(fp, expr.Inputs[0]))
	WriteF32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in f64 function returns operand 1 casted from type ui32 to type f64.
func opUI32ToF64(expr *CXExpression, fp int) {
	outV0 := float64(ReadUI32(fp, expr.Inputs[0]))
	WriteF64(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The print built-in function formats its arguments and prints them.
func opUI32Print(expr *CXExpression, fp int) {
	fmt.Println(ReadUI32(fp, expr.Inputs[0]))
}

// The built-in add function returns the sum of two ui32 numbers.
func opUI32Add(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) + ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in sub function returns the difference of two ui32 numbers.
func opUI32Sub(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) - ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in mul function returns the product of two ui32 numbers.
func opUI32Mul(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) * ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in div function returns the quotient of two ui32 numbers.
func opUI32Div(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) / ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in gt function returns true if operand 1 is greater than operand 2.
func opUI32Gt(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) > ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in gteq function returns true if operand 1 is greater than or
// equal to operand 2.
func opUI32Gteq(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) >= ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in lt function returns true if operand 1 is less than operand 2.
func opUI32Lt(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) < ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in lteq function returns true if operand 1 is less than or equal
// to operand 1.
func opUI32Lteq(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) <= ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in eq function returns true if operand 1 is equal to operand 2.
func opUI32Eq(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) == ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in uneq function returns true if operand 1 is different from operand 2.
func opUI32Uneq(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) != ReadUI32(fp, expr.Inputs[1])
	WriteBool(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in mod function returns the remainder of operand 1 / operand 2.
func opUI32Mod(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) % ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in rand function returns a pseudo-random number.
func opUI32Rand(expr *CXExpression, fp int) {
	outV0 := rand.Uint32()
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitand function returns the bitwise AND of 2 operands.
func opUI32Bitand(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) & ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitor function returns the bitwise OR of 2 operands.
func opUI32Bitor(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) | ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitxor function returns the bitwise XOR of 2 operands.
func opUI32Bitxor(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) ^ ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitclear function returns the bitwise AND NOT of 2 operands.
func opUI32Bitclear(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) &^ ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitshl function returns bits of operand 1 shifted to the left
// by number of positions specified in operand 2.
func opUI32Bitshl(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) << ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in bitshr function returns bits of operand 1 shifted to the right
// by number of positions specified in operand 2.
func opUI32Bitshr(expr *CXExpression, fp int) {
	outV0 := ReadUI32(fp, expr.Inputs[0]) >> ReadUI32(fp, expr.Inputs[1])
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), outV0)
}

// The built-in max function returns the biggest of the two operands.
func opUI32Max(expr *CXExpression, fp int) {
	inpV0 := ReadUI32(fp, expr.Inputs[0])
	inpV1 := ReadUI32(fp, expr.Inputs[1])
	if inpV1 > inpV0 {
		inpV0 = inpV1
	}
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), inpV0)
}

// The built-in min function returns the smallest of the two operands.
func opUI32Min(expr *CXExpression, fp int) {
	inpV0 := ReadUI32(fp, expr.Inputs[0])
	inpV1 := ReadUI32(fp, expr.Inputs[1])
	if inpV1 < inpV0 {
		inpV0 = inpV1
	}
	WriteUI32(GetOffsetUI32(fp, expr.Outputs[0]), inpV0)
}
