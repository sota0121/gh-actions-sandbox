package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	// Test FizzBuzz(1)
	result, err := FizzBuzz(1)
	require.NoError(t, err)
	require.Equal(t, "1", result)

	// Test FizzBuzz(3)
	result, err = FizzBuzz(3)
	require.NoError(t, err)
	require.Equal(t, "Fizz", result)

	// Test FizzBuzz(5)
	result, err = FizzBuzz(5)
	require.NoError(t, err)
	require.Equal(t, "Buzz", result)

	// Test FizzBuzz(15)
	result, err = FizzBuzz(15)
	require.NoError(t, err)
	require.Equal(t, "FizzBuzz", result)

	// Test FizzBuzz(0)
	result, err = FizzBuzz(0)
	require.Error(t, err)
	require.Equal(t, "", result)

	// Test FizzBuzz(-1)
	result, err = FizzBuzz(-1)
	require.Error(t, err)
	require.Equal(t, "", result)
}
