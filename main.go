package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/dgrijalva/jwt-go"
	// _ "github.com/gogo/protobuf/proto"
	// _ "github.com/hashicorp/golang-lru"
	// _ "github.com/owncast/owncast/logging"
)

func main() {
	catchSAST()
	fmt.Println("Hello world!")

	fmt.Println("Hello world")
	pat := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.120xrlxxx"
	fmt.Printf("I am randomly printing a PAT %s", pat)
	newsecret := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.120yyyxxx"
	fmt.Printf("I am randomly printing a new secret %s", newsecret)

	newsecret = "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.120yyyzxx"
	fmt.Printf("I am randomly printing a new secret %s", newsecret)
	fmt.Println("ASANA_CLIENT_SECRET=0f9a3a5d2b8e4f3a95a2ef1234567890")
	fmt.Println("AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE")
}

func catchSAST() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}

	userInput := os.Args[1]

	// Intentionally vulnerable: user-controlled input is executed by a shell.
	cmd := exec.Command("sh", "-c", userInput)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command failed:", err)
		return
	}

	fmt.Println(string(output))
}
