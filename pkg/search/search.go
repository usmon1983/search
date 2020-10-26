package search

import ( 
	"io/ioutil"
	"context"
	"fmt"
	"sync"
	"strings"
)

type Result struct {
	Phrase string
	Line string
	LineNum int64
	ColNum int64
}

func All(ctx context.Context, phrase string, files []string) <- chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func (ctx context.Context, fileName string, i int, ch chan <- []Result)  {
			defer wg.Done()
			resultRead := FindAll(phrase, fileName)
			if len(resultRead) > 0 {
				ch <- resultRead
			}
		}(ctx, files[i], i, ch)
	}
	go func ()  {
		defer close(ch)
		wg.Wait()
	}()
	
	cancel()
	return ch
}

func FindAll(phrase, filename string) (results []Result){
	result, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error =>", err)
	}

	temp := strings.Split(string(result), "\n")
	for l, line := range temp {
		if strings.Contains(line, phrase) {
			resultText := Result {
				Phrase: phrase,
				Line: line,
				LineNum: int64(l + 1),
				ColNum: int64(strings.Index(line, phrase)) + 1,
			}
			results = append(results, resultText)
		}
	}
	return results
}

func Any(ctx context.Context, phrase string, files []string) <- chan Result {
	ch := make(chan Result)
	wg := sync.WaitGroup{}
	result := Result{}

	for i := 0; i < len(files); i++ {
		data, err := ioutil.ReadFile(files[i])
		if err != nil {
			fmt.Println("error =>", err)
		}

		if strings.Contains(string(data), phrase) {
			tempRes := FindAny(phrase, string(data))
			if (Result{}) != tempRes {
				result = tempRes
				break
			}
		}
	}
	wg.Add(1)
	go func(ctx context.Context, ch chan <- Result) {
		defer wg.Done()
		if (Result{}) != result {
			ch <- result
		}
	}(ctx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}

func FindAny(phrase, search string) (result Result)  {
	for l, line := range strings.Split(search, "\n") {
		if strings.Contains(line, phrase) {
			return Result{
				Phrase: phrase,
				Line: line,
				LineNum: int64(l + 1),
				ColNum: int64(strings.Index(line, phrase)) + 1,
			}
		}
	}
	return result
}