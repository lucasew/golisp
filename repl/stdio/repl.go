package main

import (
	"bufio"
	"context"
	"github.com/lucasew/golisp/stdlib"
	"github.com/lucasew/golisp/stdlib/dump"
	"github.com/lucasew/golisp/stdlib/portal"
	"github.com/lucasew/golisp/stdlib/rand"
	"github.com/lucasew/golisp/toolchain/default"
	"os"
	"strings"
)

const banner = `
 ▄████ ▒█████  ██▓    ██▓ ██████ ██▓███  
 ██▒ ▀█▒██▒  ██▓██▒   ▓██▒██    ▒▓██░  ██▒
▒██░▄▄▄▒██░  ██▒██░   ▒██░ ▓██▄  ▓██░ ██▓▒
░▓█  ██▒██   ██▒██░   ░██░ ▒   ██▒██▄█▓▒ ▒
░▒▓███▀░ ████▓▒░██████░██▒██████▒▒██▒ ░  ░
 ░▒   ▒░ ▒░▒░▒░░ ▒░▓  ░▓ ▒ ▒▓▒ ▒ ▒▓▒░ ░  ░
  ░   ░  ░ ▒ ▒░░ ░ ▒  ░▒ ░ ░▒  ░ ░▒ ░     
░ ░   ░░ ░ ░ ▒   ░ ░   ▒ ░  ░  ░ ░░       
      ░    ░ ░     ░  ░░       ░          
`

func main() {
	tc := tdefault.NewDefaultToolchain(nil)
	ctx := context.Background()
	repo := stdlib.NewImporter(
		libdump.ELEMENTS,
		libportal.ELEMENTS,
		rand.ELEMENTS,
	)
	tc.Import(
		map[string]interface{}{
			"import": stdlib.AsMacro(repo),
			"repo":   repo,
		},
	)
	in := bufio.NewReader(os.Stdin)
	lines := []string{}
	println(banner)
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
			ret, err := tc.EvalString(ctx, stmt)
			if err != nil {
				println(err.Error())
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
