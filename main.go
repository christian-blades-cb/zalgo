package main // import "github.com/christian-blades-cb/zalgo"

import (
	"bufio"
	"fmt"
	"github.com/kortschak/zalgo"
	"os"
)

func main() {
	z := zalgo.NewCorrupter(os.Stdout)

	z.Zalgo = func(n int, r rune, z *zalgo.Corrupter) bool {
		z.Up += 0.1
		z.Middle += complex(0.01, 0.01)
		z.Down += complex(real(z.Down)*0.1, 0)
		return false
	}

	z.Up = complex(0, 0.2)
	z.Middle = complex(0, 0.2)
	z.Down = complex(0.001, 0.3)

	rd := bufio.NewReader(os.Stdin)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		fmt.Fprintln(z, string(line))
	}

}
