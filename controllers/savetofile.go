package controllers

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Save(encryptedS string) {
	date := time.Now().String()
	fileName := "./crypted/crypted-" + strings.TrimSpace(date) + ".txt"
	file, _ := os.Create(fileName)
	defer file.Close()
	ln, _ := io.WriteString(file, encryptedS)
	fmt.Println(ln)
}
