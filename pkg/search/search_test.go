package search

import (
	"fmt"
	"context"
	"testing"
	
)


func TestAll_user(t *testing.T)  {
	res := All(context.Background(), "Alif", []string{"D:/projectsGo/search/data/1.txt"}) 
	result, ok := <- res
	if !ok {
		fmt.Println("error ok = ", ok)
	}
	fmt.Println("result => ", result)
}
