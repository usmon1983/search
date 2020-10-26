package main

import ( 
	//"io/ioutil"
	"github.com/usmon1983/search/pkg/search"
	"context"
	//"log"
	"fmt"
)


func main(){
	/*
	res, err := ioutil.ReadFile("D:/projectsGo/search/data/1.txt")
	
	if err != nil {
		fmt.Println("err = ", err)
	}
	data := string(res)
	fmt.Println(data)*/
	
	res := search.All(context.Background(), "usmon", []string{"D:/projectsGo/search/data/1.txt"}) 

	v, ok := <- res
	if !ok {
		fmt.Println("error ok = ", ok)
	}
	fmt.Println(v)
}