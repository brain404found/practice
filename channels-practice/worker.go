package main

import (
	"sync"
)

func createWorkerPool(poolSize int) chan Params {
	paramCh := make(chan Params)
	var wgWorker sync.WaitGroup

	for i := 0; i < poolSize; i++ {
		wgWorker.Add(1)
		go func() {
			defer wgWorker.Done()
			for params := range paramCh {
				for i := 0; i < len(params.MathParams); i += 100 {
					end := i + 100
					if end > len(params.MathParams) {
						end = len(params.MathParams)
					}
					subParams := params.MathParams[i:end]
					subParamsStruct := Params{UserID: params.UserID, MathParams: subParams}
					processParams(subParamsStruct)
				}
			}
		}()
	}

	go func() {
		wgWorker.Wait()
		close(paramCh)
	}()

	return paramCh
}
