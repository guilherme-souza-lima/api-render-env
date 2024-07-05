package main

import (
	"api-env-example/cmd"
	"api-env-example/infra"
	"context"
)

func main() {
	env := infra.NewConfig()

	ctx := context.Background()
	cmd.StartHttp(ctx, env)
}
