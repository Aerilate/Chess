package main

import "flag"

func main() {
	fileName := flag.String("i", "", "i")
	flag.Parse()

	movesQueue, _ := gameReader(*fileName)

	v := ViewController{}
	v.loadQueue(movesQueue)
	v.start()
}
