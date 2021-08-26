package main

import "flag"

func main() {
	fileName := flag.String("i", "", "i")
	flag.Parse()

	file, _ := gameReader(*fileName)
	vc := ViewController{NewGameState()}
	vc.load(file)
	vc.start()
}
