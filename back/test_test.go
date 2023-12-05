package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	formphaser "github.com/Hana-ame/go-form-phaser"

	"github.com/joho/godotenv"
)

type A struct {
	SB int
}
type AA struct {
	A
}

func TestInherite(t *testing.T) {
	aa := AA{A{1}}
	a, _ := json.Marshal(aa)
	fmt.Printf("%s", a)
}

func TestEnv(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println(s3Bucket, secretKey)
}

func TestFormphaser(t *testing.T) {
	type C struct {
		A int    `formphaser:"aaa"`
		B string `formphaser:"bbb"`
	}
	type O struct {
		A int      `formphaser:"aaa"`
		B string   `formphaser:"bbb"`
		C C        `formphaser:"ccc"`
		D []string `formphaser:"ddd"`
	}
	o := O{}
	formphaser.Unmarshal(func(s string) []string { return []string{s} }, &o)
	fmt.Printf("%+v\n", o)
}
