package hw05parallelexecution

import (
	"crypto/sha256"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaskSha256(t *testing.T) {
	t.Run("calculation sha256", func(t *testing.T) {
		countTask := 20
		tasks := make([]Task, 0, countTask)
		testPhrase := []string{
			"Flying airplane.",
			"Singing canary.",
			"Burning sunset.",
			"Smell of freshly cut grass.",
			"Cold winter day.",
			"Delicious chocolate cake.",
			"Constant toothache.",
			"Roaring waterfall.",
			"Endless starry sky.",
			"Loud noise of the train.",
			"Spring breeze.",
			"Cry of wild geese.",
			"Warm summer sun.",
			"Joyful childish laughter.",
			"Rustling autumn leaves.",
			"Aromatic coffee.",
			"Cozy home hearth.",
			"Shiny new car.",
			"Sad melody.",
			"Magical fairy-tale world.",
		}
		for i := 0; i < countTask; i++ {
			tasks = append(tasks, func() error {
				sha256.Sum256([]byte(testPhrase[rand.Intn(countTask-1)]))
				return nil
			})
		}

		workersCount := 10
		maxErrorsCount := 2
		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)
	})
}
