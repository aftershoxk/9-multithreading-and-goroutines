package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	testData := generateRandomElements(0)
	require.Equal(t, []int{}, testData)

	testData = generateRandomElements(-10)
	require.Equal(t, []int{}, testData)

	testData = generateRandomElements(10)
	require.Equal(t, 10, len(testData))

	data1 := generateRandomElements(10)
	data2 := generateRandomElements(10)
	require.NotEqual(t, data1, data2)
}

func TestMaximum(t *testing.T) {
	require.Equal(t, 0, maximum([]int{}))

	require.Equal(t, 5, maximum([]int{5}))

	require.Equal(t, 10, maximum([]int{5, 1, 3, 1, 10, 2, 8}))
	require.Equal(t, 7, maximum([]int{1, 7, 4, 7, 7, 5, 4}))
	require.Equal(t, 3, maximum([]int{3, 3, 3, 3}))
}

func TestMaxChunks(t *testing.T) {
	data := generateRandomElements(10_000)
	require.Equal(t, maximum(data), maxChunks(data))

	data = generateRandomElements(10_005)
	require.Equal(t, maximum(data), maxChunks(data))

	data = generateRandomElements(5)
	require.Equal(t, maximum(data), maxChunks(data))

	require.Equal(t, 0, maxChunks([]int{}))

	require.Equal(t, 42, maxChunks([]int{42}))

	data = make([]int, 100)
	for i := range data {
		data[i] = 7
	}
	require.Equal(t, 7, maxChunks(data))

	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, 10, maxChunks(data))

}
