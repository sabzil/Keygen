package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"os"
)

func main() {
	len := 64
	buf := make([]byte, len)
	rand.Read(buf)
	key := base64.StdEncoding.EncodeToString(buf)

	path := "./key.txt"
	f, _ := os.Create(path)
	writer := bufio.NewWriter(f)
	writer.WriteString(key)
	writer.Flush()
}
