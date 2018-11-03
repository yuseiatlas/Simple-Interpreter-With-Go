package repl

import (
	"bufio"
	"fmt"
	"github.com/yuseiatlas/lexer/lexer"
	"github.com/yuseiatlas/lexer/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		PrintResult(line)
	}
}

func PrintResult(input string) {
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
