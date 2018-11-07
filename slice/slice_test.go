package main

import (
	"sync"
	"testing"
)
var n = 4

func BenchmarkMutex(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		wg := new(sync.WaitGroup)
		mu := sync.Mutex{}
		arr := make([]int, 0, n)
		for j := 0; j < n; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				mu.Lock()
				arr = append(arr, j)
				mu.Unlock()
			}(j)
			wg.Wait()
		}
	}
}

func BenchmarkIndex(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		wg := new(sync.WaitGroup)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				arr[j] = j
			}(j)
			wg.Wait()
		}
	}
}

func BenchmarkChannel(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		wg := new(sync.WaitGroup)
		arr := make(chan int, n)
		for j := 0; j < n; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				arr <- j
			}(j)
			wg.Wait()
		}
	}
}

func BenchmarkMap(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		arr := make(map[int]struct{}, n)
		for j := 0; j < n; j++ {
			arr[j] = struct{}{}
		}
	}
}

func BenchmarkAppend(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		arr := make([]int, 0, n)
		for j := 0; j < n; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppend2(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		var arr []int
		for j := 0; j < n; j++ {
			arr = append(arr, j)
		}
	}
}