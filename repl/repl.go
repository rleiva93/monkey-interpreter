// REPL -> Read Eval Print Loop

package repl

import (
	"bufio"
	"fmt"
	"io"

	"../lexer"
	"../token"
)

// PROMPT depcits chars waiting for an instruction
const PROMPT = ">> "

// Start initializes the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
