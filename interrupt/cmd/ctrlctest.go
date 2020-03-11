package main

import (
	log "github.com/p9c/logi"
	"github.com/p9c/util/interrupt"
)

func main() {
	interrupt.AddHandler(func() {
		log.Println("IT'S THE END OF THE WORLD!")
	})
	<-interrupt.HandlersDone
}
