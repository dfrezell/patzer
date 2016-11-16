package patzer

import (
	"bufio"
	"fmt"
	"os"
)

type UCI struct {
}

func (u *UCI) Loop() {
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		ln := reader.Text()
		fmt.Printf("%d, %s\n", len(ln), ln)
	}

	if err := reader.Err(); err != nil {
		fmt.Printf("%s\n", err)
	}
}
