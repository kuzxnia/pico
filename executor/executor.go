package executor

import (
	"errors"
	"fmt"
	"io"
	"os"
	"pico/evaluator"
	"pico/lexer"
	"pico/object"
	"pico/parser"
	"time"
)

func ExecuteFile(path string, out io.Writer) {
	println(path)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("file does not exist")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// println(string(file))
	env := object.NewEnvironment()

	l := lexer.New(string(file))
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		panic("parser found errors")
	}

	evalueted := evaluator.Eval(program, env)

	// print only printed values
	if evalueted != nil {
		io.WriteString(out, evalueted.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
