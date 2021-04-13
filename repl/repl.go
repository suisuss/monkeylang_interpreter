package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/suisuss/monkey-interpreter/lexer"
	"github.com/suisuss/monkey-interpreter/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
			/*
			%v is for formatting string with structs, the plus flag (%+v) adds field names
			*/
		}
	}
}