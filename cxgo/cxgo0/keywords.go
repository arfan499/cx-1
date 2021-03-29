package cxgo0

var KeywordMap map[string]int = map[string]int{
	"func":      FUNC,
	"var":       VAR,
	"package":   PACKAGE,
	"if":        IF,
	"else":      ELSE,
	"for":       FOR,
	"struct":    STRUCT,
	"import":    IMPORT,
	"return":    RETURN,
	"goto":      GOTO,
	"new":       NEW,
	"bool":      BOOL,
	"i8":        I8,
	"ui8":       UI8,
	"i16":       I16,
	"ui16":      UI16,
	"i32":       I32,
	"ui32":      UI32,
	"f32":       F32,
	"i64":       I64,
	"ui64":      UI64,
	"f64":       F64,
	"str":       STR,
	"aff":       AFF,
	"union":     UNION,
	"enum":      ENUM,
	"const":     CONST,
	"case":      CASE,
	"default":   DEFAULT,
	"switch":    SWITCH,
	"break":     BREAK,
	"continue":  CONTINUE,
	"type":      TYPE,
	":dl":       DSTATE,
	":dLocals":  DSTATE,
	":ds":       DSTACK,
	":dStack":   DSTACK,
	":dp":       DPROGRAM,
	":dProgram": DPROGRAM,
	//":package":  SPACKAGE,
	//":struct":   SSTRUCT,
	//":func":     SFUNC,
	":rem":      REM,
	//":step":     STEP,
	//":tStep":    TSTEP,
	//":tstep":    TSTEP,
	//":pStep":    PSTEP,
	//":pstep":    PSTEP,
	":aff":      CAFF,
	"def":       DEF,
	"clauses":   CLAUSES,
	"field":     FIELD,
	"true":      BOOLEAN_LITERAL,
	"false":     BOOLEAN_LITERAL,
}