package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
	"flag"
	"io/ioutil"
	"monkey/lexer"
	"monkey/parser"
	"monkey/evaluator"
	"io"
	"monkey/object"
)

var filePath = flag.String("f", "Fibonacci.monkey", "file.")

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("f = %s\n", *filePath)

	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		// エラー処理
	}

	if *filePath != "" {
		//fmt.Print(string(data))
		env := object.NewEnvironment()

		l := lexer.New(string(data))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			os.Exit(0)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	} else {
		// 解析のために必ずパースを実行する。
		flag.Parse()
		fmt.Printf("Hello %s! this is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n	")
	io.WriteString(out, " parser errors:	")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`
