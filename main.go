package main

import (
	"bufio"
	"fmt"
	"github.com/kr/pretty"
	"github.com/qq906907952/goyacc_cal/parse"
	"os"
	"strconv"
	"strings"
)

func printExpr(e *parse.Expr) {
	pretty.Println(e)
}

func main() {
	//r, _ := parse.Parse("1+2 *(9  + 6) / 12.3 + (2-9+6) * (3+7)+ 69-789/856*56+15")
	//r, _ := parse.Parse("1+2 *(9  + 6) ")
	//printExpr(r)

	r := bufio.NewReader(os.Stdin)
	for {
		os.Stdout.WriteString("calculate> ")
		l, _, err := r.ReadLine()
		if err != nil {
			panic(err)
		}
		line := string(l)
		if strings.TrimSpace(line) == "" {
			continue
		}
		expr, err := parse.Parse(line)
		if err != nil {
			os.Stdout.WriteString(err.Error())
		} else {
			os.Stdout.WriteString(fmt.Sprintln())
			printExpr(expr)
			os.Stdout.WriteString(fmt.Sprintln())
			os.Stdout.WriteString("result: " + strconv.FormatFloat(expr.GetVal(), 'f', 24, 64))
		}
		os.Stdout.WriteString(fmt.Sprintln())
	}

}
