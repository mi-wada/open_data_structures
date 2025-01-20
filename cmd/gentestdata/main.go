package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("correct: %s n", os.Args[0]))
	}
	nStr := os.Args[1]
	n, err := strconv.Atoi(removeUnderscores(nStr))
	if err != nil {
		panic(fmt.Sprintf("invalid number: %s", nStr))
	}

	gen(os.Stdout, n)
}

func removeUnderscores(s string) string {
	return strings.ReplaceAll(s, "_", "")
}

func gen(w io.Writer, n int) {
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	for i := 0; i < n; i++ {
		fmt.Fprintf(bw, "%d\n", rand.Int())
	}
}
