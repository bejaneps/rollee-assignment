package inmemory_test

import (
	"testing"

	internalStorage "github.com/bejaneps/rollee-assignment/test2/internal/storage"
	"github.com/bejaneps/rollee-assignment/test2/internal/storage/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestWord(t *testing.T) {
	testCases := map[string]struct {
		insertWords  []string
		word         string
		expectedWord string
		expectedErr    error
	}{
		"success": {
			insertWords: []string{
				"hello",
				"hello",
				"world",
				"world",
				"hell",
				"hel",
				"hell",
				"hello",
			},
			word:         "hel",
			expectedWord: "hello",
		},
		"word-not-found": {
			insertWords: []string{
				"hello",
				"hell",
			},
			word:         "foo",
			expectedWord: "",
			expectedErr: internalStorage.ErrWordNotFound,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			storage := inmemory.New()

			for _, word := range tt.insertWords {
				err := storage.UpsertWord(word)
				assert.NoError(t, err)
			}

			word, err := storage.GetMostFrequentWord(tt.word)
			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedWord, word)
		})
	}
}
