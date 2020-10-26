package main

import ( 
	"github.com/usmon1983/search/pkg/search"
	"context"
	//"log"
	
	"fmt"
)


func main(){
	//ctx context.Context, phrase string, files []string
	//ch := make(chan []Result)
	res := search.All(context.Background(), "usmon", []string{"D:/projectsGo/search/data/1.txt"}) 

	v, ok := <- res
	if !ok {
		fmt.Println("error ok = ", ok)
	}
	fmt.Println(v)
	//return ch
}