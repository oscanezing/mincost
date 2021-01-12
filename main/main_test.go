package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllEquals(t *testing.T) {
	costs := [][]int{
		{10, 10},
		{10, 10},
		{10, 10},
		{10, 10},
	}

	assert.Equal(t, 40, findMinCost(costs))
}

func TestOrdered(t *testing.T) {
	costs := [][]int{
		{10, 20},
		{10, 50},
		{29, 200},
		{30, 100},
	}

	assert.Equal(t, 129, findMinCost(costs))
}

func TestUnOrdered(t *testing.T) {
	costs := [][]int{
		{30, 100},
		{10, 20},
		{29, 200},
		{10, 50},
	}

	assert.Equal(t, 129, findMinCost(costs))
}