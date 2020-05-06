package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(ctx context.Context) int {
	return ctx.Value(ctxUserID).(int)
}

func AuthToken(ctx context.Context) string {
	return ctx.Value(ctxAuthToken).(string)
}

func main() {
	processRequest(1, "123456789")
}

func processRequest(userID int, token string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, token)
	doProcessRequest(ctx)
}

func doProcessRequest(ctx context.Context) {
	fmt.Println(AuthToken(ctx))
	fmt.Println(UserID(ctx))
}
