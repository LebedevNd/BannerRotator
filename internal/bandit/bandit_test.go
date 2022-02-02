package bandit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateRating(t *testing.T) {

	rating, err := calculateRating(1, 10, 100)
	require.Nil(t, err)
	require.Equal(t, 1.06, rating)

	_, err = calculateRating(10, 1, 100)
	require.Error(t, err)

	_, err = calculateRating(10, 100, 10)
	require.Error(t, err)
}
