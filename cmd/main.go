package main

import ( 
	"context"
	"log"
	
)
type Result struct {
	Phrase string
	Line string
	LineNum int64
	ColNum int64
}

func All(ctx context.Context, phrase string, files []string) <- chan []Result {
	ch := make(chan []Result)
	log.Print(ch)
	return ch
}