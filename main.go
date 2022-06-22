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
