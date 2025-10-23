package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("Длина слайса должна быть положительным числом")
		return []int{}
	}
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		value := rand.Int()
		data = append(data, value)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	switch {
	case len(data) == 0:
		fmt.Println("Длина слайса должна быть больше 0")
		return 0
	case len(data) == 1:
		return data[0]
	}

	max := data[0]
	for _, value := range data {
		if max < value {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	ch := make(chan int)

	count := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		first := i * count
		last := first + count
		if i == CHUNKS-1 {
			last = len(data)
		}
		part := data[first:last]

		wg.Add(1)
		go func(part []int) {
			defer wg.Done()
			ch <- maximum(part)
		}(part)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	maxValues := []int{}
	for v := range ch {
		maxValues = append(maxValues, v)
	}
	absoluteNumber := maximum(maxValues)
	return absoluteNumber
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	t := time.Now()
	max := maximum(data)
	elapsed := time.Since(t).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	t = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(t).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
