package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/p3t33/interpreter/lexer"
	"github.com/p3t33/interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		monkey_lexer := lexer.CreateNewLexer(line)

		for tok := monkey_lexer.NextToken(); tok.Type != token.EOF; tok = monkey_lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
