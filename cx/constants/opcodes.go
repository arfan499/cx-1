package constants

// op codes
// nolint golint
const (
	OP_IDENTITY = iota + 1
	OP_JMP
	OP_ABS_JMP
	OP_JMP_EQ
	OP_JMP_UNEQ
	OP_JMP_GT
	OP_JMP_GTEQ
	OP_JMP_LT
	OP_JMP_LTEQ
	OP_JMP_ZERO
	OP_JMP_NOT_ZERO

	OP_NOP

	OP_GOTO
	OP_BREAK
	OP_CONTINUE
	OP_DEBUG

	OP_SERIALIZE
	OP_DESERIALIZE

	//TODO: remove undeclared type, this is stupid
	START_OF_OPERATORS            //WHAT KIND OF "OPERATORS"
	START_OF_COMPARISON_OPERATORS = iota + START_OF_OPERATORS
	OP_EQUAL
	OP_UNEQUAL
	OP_LT
	OP_GT
	OP_LTEQ
	OP_GTEQ
	END_OF_COMPARISON_OPERATORS
	//TODO: WTF?
	START_OF_ARITHMETIC_OPERATORS = iota + END_OF_COMPARISON_OPERATORS
	OP_BITAND
	OP_BITOR
	OP_BITXOR
	OP_BITCLEAR
	OP_BITSHL
	OP_BITSHR
	OP_MUL
	OP_DIV
	OP_MOD
	OP_ADD
	OP_SUB
	OP_NEG
	END_OF_ARITHMETIC_OPERATORS
	END_OF_OPERATORS = iota + END_OF_ARITHMETIC_OPERATORS
	//TODO: Why this start/end for operators?

	OP_ASSERT
	OP_TEST
	OP_PANIC

	OP_APPEND

	OP_BOOL_OR
	OP_BOOL_AND
	OP_BOOL_NOT

	END_OF_NAMED_OPCODES
)

var (
	START_PARSE_OPS int
	END_PARSE_OPS   int
)
