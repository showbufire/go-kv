package main

import (
	"log"
	"os"
	"strconv"

	"github.com/showbufire/kv"
	"github.com/showbufire/kv/lock"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("not enough arguments")
	}
	nthreads, _ := strconv.Atoi(os.Args[1])
	nsize, _ := strconv.Atoi(os.Args[2])
	m := lock.NewMemo()
	kv.Run(m, nthreads, nsize)
}
