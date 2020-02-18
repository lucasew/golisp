package main

import (
    "os"
    "bufio"
    "strings"
    parser "github.com/lucasew/golisp/parser/default"
    vm_lib "github.com/lucasew/golisp/vm/default"
)

func main() {
    vm := vm_lib.NewVM(nil)
    in := bufio.NewReader(os.Stdin)
    lines := []string{}
    parenthesis := 0
    for {
        for i := -1; i < parenthesis; i++ {
            print(">")
        }
        print(" ")
        line, err := in.ReadString('\n')
        if err != nil {
            println(err.Error())
            break
        }
        parenthesis += calc_delta(line)
        lines = append(lines, line)
        if parenthesis == 0 {
            stmt := strings.Join(lines, "\n")
            lines = []string{}
            ast, err := parser.Parse(stmt)
            if err != nil {
                println(err.Error())
                break
            }
            ret, err := vm.Eval(ast)
            if err != nil {
                println(err.Error())
                break
            }
            println(ret.Repr())
        }
    }
}

func calc_delta(s string) int {
    delta := 0
    for _, c := range s {
        if c == '(' {
            delta++
        }
        if c == ')' {
            delta--
        }
    }
    return delta
}
