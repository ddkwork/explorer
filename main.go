package main

import (
	"explorer"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("explorer", func(w *unison.Window) {
		w.Content().AddChild(explorer.New(".").Layout())
	})
}
