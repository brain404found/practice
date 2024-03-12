package main

import "sync"

func createWorkerPool(poolSize int, paramCh chan Params) {
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

	wgWorker.Wait()
	close(paramCh)
}
