package main

import (
	"explorer"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("explorer", func(w *unison.Window) {
		w.Content().AddChild(explorer.New(".").Layout())
	})
}
