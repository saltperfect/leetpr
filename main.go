package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format(time.RFC1123))
	fmt.Println(time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30).Format(time.RFC3339))
}
