package main

import (
	"bytes"
	/* "compress/zlib" */
	"bufio"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	// var b bytes.Buffer
	//	w := zlib.NewWriter(&b)
	//	w.Write([]byte(arg))
	//	w.Close()
	//	fmt.Println(b);
	input := bufio.NewReader(os.Stdin)
	buff := make([]byte, 0)
	for {
		c, err := input.ReadByte()
		if err == io.EOF {
			break
		}
		buff = append(buff, c)
	}
	deflate_reader := flate.NewReader(bytes.NewReader(buff))
	deflate_bytes, _ := io.ReadAll(deflate_reader)
	deflate := string(deflate_bytes)
	if len(deflate) == 0 {
		zlib_reader := flate.NewReader(bytes.NewReader(buff[2:]))
		zlib_bytes, _ := io.ReadAll(zlib_reader)
		zlib := string(zlib_bytes)
		fmt.Println(zlib)
		return
	}
	fmt.Println(deflate)
}
