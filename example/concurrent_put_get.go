// Open Source: MIT License
// Author: Leon Ding <ding@ibyte.me>
// Date: 2022/3/10 - 3:36 PM - UTC/GMT+08:00

package main

import (
	"fmt"
	"github.com/auula/bottle"
	"strconv"
	"sync"
)

func init() {
	if err := bottle.Load("./config.yaml"); err != nil {
		fmt.Println(err)
	}
	bottle.SetIndexSize(100000)
}

func main() {

	defer bottle.Close()

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				k := strconv.Itoa(m*1000 + j)
				v := strconv.Itoa(m*1000 + j)
				if err := bottle.Put([]byte(k), []byte(v)); err != nil {
					fmt.Println(err, k, v)
				}
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				k := strconv.Itoa(m*1000 + j)
				v := strconv.Itoa(m*1000 + j)
				d := bottle.Get([]byte(k))
				if d.Err != nil {
					fmt.Println("Get:", d.Err, k, v)
				} else if string(d.Value) != v {
					fmt.Println("Not Equal:", string(d.Value), v)
				} else {
					// fix bug:
					fmt.Println("Get:", d.Err, k, v)
				}
			}
		}(i)
	}
	wg.Wait()
}
