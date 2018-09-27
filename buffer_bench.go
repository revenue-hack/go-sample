package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	//noBuffer()
	buffer()
	//bufferOnFile()
}

func noBuffer() {
	// wirteを100回call
	for i := 0; i < 100; i++ {
		fmt.Fprintln(os.Stdout, strings.Repeat("x", 100))
	}
}

func buffer() {
	// bufferによってwriteのcallを少なくできる
	buf := bufio.NewWriter(os.Stdout)
	for i := 0; i < 100; i ++ {
		fmt.Fprintln(buf, strings.Repeat("x", 100))
	}
	buf.Flush()
}

func bufferOnFile() {
	// bufferによってwriteのcallを少なくできる
	fp, _ := os.Open("./testdata/bufferbench_in.txt")
	buf := bufio.NewWriter(fp)
	for i := 0; i < 100; i ++ {
		fmt.Fprintln(buf, strings.Repeat("x", 100))
	}
	buf.Flush()
}

/*
func bufferFileScan() {
	fp, _ := os.Open("./testdata/bufferbench_in.txt")
	fpB, _ := os.Open("./bufferbench_out.txt")
	buf := bufio.NewWriter(fpB)
	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		fmt.Fprintf(buf, strings.Repeat("x", 100))
	}
	buf.Flush()
}
*/
