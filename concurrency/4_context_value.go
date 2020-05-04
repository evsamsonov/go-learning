package main

import (
	"context"
	"fmt"
)

func main() {
	processRequest(1, "123456789")
}

func processRequest(userID int, token string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "token", token)
	doProcessRequest(ctx)
}

func doProcessRequest(ctx context.Context) {
	fmt.Println(ctx.Value("userID"))
	fmt.Println(ctx.Value("token"))
}
