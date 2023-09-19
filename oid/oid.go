package main

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "hello\n"
	fmt.Println("raw:")
	fmt.Println(get_sha_str(str))
	fmt.Println("blob:")
	byte_len := fmt.Sprintf("%d", len([]byte(str)))
	blob := "blob " + byte_len + "\x00" + str
	sha_blob := get_sha_str(blob)
	fmt.Println(sha_blob)
	fmt.Println("zipped:")
	deflate_bytes := deflate(blob)
	zlib_bytes := zlib_format(blob)
	fmt.Println("deflate hex: " + hex.EncodeToString(deflate_bytes))
	fmt.Println("zlib hex: " + hex.EncodeToString(zlib_bytes))
	fmt.Println(get_sha_byte(deflate_bytes))
}
func deflate(str string) []byte {
	var buf bytes.Buffer
	w, _ := flate.NewWriter(&buf, 6)
	w.Write([]byte(str))
	w.Flush()

	return buf.Bytes()
}
func zlib_format(str string) []byte {
	var deflate bytes.Buffer
	writer := zlib.NewWriter(&deflate)
	writer.Write([]byte(str))
	writer.Flush()
	return deflate.Bytes()
}
func get_sha_str(str string) string {
	hasher := sha1.New()
	str_bytes := []byte(str)
	hasher.Write(str_bytes)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
func get_sha_byte(str []byte) string {
	hasher := sha1.New()
	hasher.Write(str)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
