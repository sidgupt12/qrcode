package main

import (
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	//This is used to read the content of the file

	b, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	abc := string(b)

	newQR := "abcd-qr.png"

	//This is used to create a new QR code

	err = qrcode.WriteFile(abc, qrcode.Medium, 512, newQR)

	if err != nil {
		log.Fatal(err)
	}

}
