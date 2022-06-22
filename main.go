package main

import (
	"log"
	"os"
)

var (
	outfile, _ = os.Create("dcscragor.log")
	l          = log.New(outfile, "", log.LstdFlags|log.Lshortfile)
)

func main() {
	ConnectToDC()
}

// TO DO
//
// Logging mechanism will be optimized. It logs every time in while loop.
