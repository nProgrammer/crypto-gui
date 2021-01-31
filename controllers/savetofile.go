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
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ln, _ := io.WriteString(file, encryptedS)
	fmt.Println(ln)
}
