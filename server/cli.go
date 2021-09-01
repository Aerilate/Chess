package main

import (
	"flag"
)

func main() {
	fileName := flag.String("i", "", "i")
	flag.Parse()

	vc := ViewController{}
	_ = vc.loadSavedGame(*fileName)
	vc.start()
}
