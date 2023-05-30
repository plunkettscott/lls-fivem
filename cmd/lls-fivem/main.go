package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/responserms/lls-fivem/internal/command"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	defer cancel()

	if err := command.Run(ctx); err != nil {
		panic(err)
	}
}
