package main

import (
	"context"
	"time"

	"knative.dev/client-pkg/pkg/output/tui"
)

func main() {
	ctx := context.Background()
	w := tui.NewWidgets(ctx)
	w.NewSpinner("Loading...").With(func(sp tui.Spinner) error {
		<-time.After(2 * time.Second)
		return nil
	})
}
