package lexer

// token kind
const (
	TOKEN_EOF         = iota           // end-of-file
	TOKEN_VARARG                       // ...
	TOKEN_SEP_SEMI                     // ;
	TOKEN_SEP_COMMA                    // ,
	TOKEN_SEP_DOT                      // .
	TOKEN_SEP_COLON                    // :
	TOKEN_SEP_LABEL                    // ::
	TOKEN_SEP_LPAREN                   // (
	TOKEN_SEP_RPAREN                   // )
	TOKEN_SEP_LBRACK                   // [
	TOKEN_SEP_RBRACK                   // ]
	TOKEN_SEP_LCURLY                   // {
	TOKEN_SEP_RCURLY                   // }
	TOKEN_OP_ASSIGN                    // =
	TOKEN_OP_MINUS                     // - (sub or unm)
	TOKEN_OP_WAVE                      // ~ (bnot or bxor)
	TOKEN_OP_ADD                       // +
	TOKEN_OP_MUL                       // *
	TOKEN_OP_DIV                       // /
	TOKEN_OP_IDIV                      // //
	TOKEN_OP_POW                       // ^
	TOKEN_OP_MOD                       // %
	TOKEN_OP_BAND                      // &
	TOKEN_OP_BOR                       // |
	TOKEN_OP_SHR                       // >>
	TOKEN_OP_SHL                       // <<
	TOKEN_OP_CONCAT                    // ..
	TOKEN_OP_LT                        // <
	TOKEN_OP_LE                        // <=
	TOKEN_OP_GT                        // >
	TOKEN_OP_GE                        // >=
	TOKEN_OP_EQ                        // ==
	TOKEN_OP_NE                        // ~=
	TOKEN_OP_LEN                       // #
	TOKEN_KW_NOT					   // not
	TOKEN_KW_IS						   // is
	TOKEN_KW_IF						   // if
	TOKEN_KW_ELSE					   // else
	TOKEN_KW_NULL					   // null
	TOKEN_KW_IMPORT					   // import
	TOKEN_KW_FALSE					   // false
	TOKEN_KW_TRUE					   // true
	TOKEN_KW_BREAK					   // break
	TOKEN_KW_RETURN					   // return
	TOKEN_KW_GOTO					   // goto
	TOKEN_KW_BLOCK					   // block
	TOKEN_KW_DEF					   // def
	TOKEN_KW_OBJ					   // obj
	TOKEN_KW_WHEN					   // when
	TOKEN_KW_LAMBDA					   // lambda
	TOKEN_KW_METHOD					   // method
	TOKEN_KW_SELF					   // self
	TOKEN_KW_PRINT					   // print
	TOKEN_KW_LOOP					   // loop
	TOKEN_IDENTIFIER                   // identifier
	TOKEN_NUMBER                       // number literal
	TOKEN_STRING                       // string literal
	TOKEN_OP_UNM      = TOKEN_OP_MINUS // unary minus
	TOKEN_OP_SUB      = TOKEN_OP_MINUS
	TOKEN_OP_BNOT     = TOKEN_OP_WAVE
	TOKEN_OP_BXOR     = TOKEN_OP_WAVE
)

var keywords = map[string]int{
	"not": TOKEN_KW_NOT,					   // not
	"is": TOKEN_KW_IS,					   // is
	"if": TOKEN_KW_IF,						   // if
	"else": TOKEN_KW_ELSE,					   // else
	"null": TOKEN_KW_NULL,					   // null
	"import": TOKEN_KW_IMPORT,					   // import
	"false": TOKEN_KW_FALSE,					   // false
	"true": TOKEN_KW_TRUE,					   // true
	"break": TOKEN_KW_BREAK,					   // break
	"return": TOKEN_KW_RETURN,					   // return
	"goto": TOKEN_KW_GOTO,					   // goto
	"block": TOKEN_KW_BLOCK,					   // block
	"def": TOKEN_KW_DEF,					   // def
	"obj": TOKEN_KW_OBJ,					   // obj
	"when": TOKEN_KW_WHEN,					   // when
	"lambda": TOKEN_KW_LAMBDA,					   // lambda
	"method": TOKEN_KW_METHOD,					   // method
	"self": TOKEN_KW_SELF,					   // self
	"print": TOKEN_KW_PRINT,					   // print
}