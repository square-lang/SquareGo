package main

import "fmt"
import "io/ioutil"
import "os"
import "./compiler/lexer"

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		testLexer(string(data), os.Args[1])
	}
}

func testLexer(chunk, chunkName string) {
	Lexer := lexer.NewLexer(chunk, chunkName)
	for {
		line, kind, token := Lexer.NextToken()
		fmt.Printf("[%2d] [%-10s] %s\n",
			line, kindToCategory(kind), token)
		if kind == lexer.TOKEN_EOF {
			break
		}
	}
}

func kindToCategory(kind int) string {
	switch {
	case kind < lexer.TOKEN_SEP_SEMI:
		return "other"
	case kind <= lexer.TOKEN_SEP_RCURLY:
		return "separator"
	case kind <= lexer.TOKEN_OP_LEN:
		return "operator"
	case kind <= lexer.TOKEN_KW_LOOP:
		return "keyword"
	case kind == lexer.TOKEN_IDENTIFIER:
		return "identifier"
	case kind == lexer.TOKEN_NUMBER:
		return "number"
	case kind == lexer.TOKEN_STRING:
		return "string"
	default:
		return "other"
	}
}