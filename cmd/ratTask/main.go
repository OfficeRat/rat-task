package main

import (
	"officerat/ratTask/internal/args"
	"os"
)

func main()  {
	args := os.Args[1:]
	Args.ArgHandler(args)
}



