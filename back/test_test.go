package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

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
