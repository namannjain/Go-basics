package main

import (
	"fmt"
	"sync"
)

//To save such unused memory whcih is not freed up - sync.Pool helps us

type SomeObject struct {
	Data []byte
}

func createObject() *SomeObject {
	//simulate creation of some object
	return &SomeObject{
		Data: make([]byte, 1024*1024), //1MB of data
	}
}

func main() {
	var memoryPiece int //0

	//sync pool
	objectPool := sync.Pool{
		//by default new function any type ka leta hai - so interface{} is any so thats why
		New: func() interface{} {
			memoryPiece += 1
			return createObject()
		},
	}

	//create a large number of objects without reusing
	const workers = 1024 * 1024
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			//get from local pool
			obj := objectPool.Get()
			//some code operatrions which u can perfomr here

			//place back in pool - if u remove this line adn then run it will take hell lot of time beacause memory wont be reusing
			objectPool.Put(obj)
			wg.Done()
		}()
	}

	//the objects are no longer used but they still occupy space
	fmt.Println("Done: ", memoryPiece)
}
