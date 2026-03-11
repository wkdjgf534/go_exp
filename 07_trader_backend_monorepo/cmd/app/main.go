package app

import (
	"context"
	"trader-backend_monorepo/internal/app"
)

func main() {
	ctx := context.Background()

	application, err := app.NewApplication(ctx)
	if err != nil {
		panic(err)
	}

	if err = application.Run(ctx); err != nil {
		panic(err)
	}
}
