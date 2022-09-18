package service_test

import (
	"errors"
	"testing"

	"github.com/bejaneps/rollee-assignment/test2/internal/service"
	"github.com/bejaneps/rollee-assignment/test2/internal/service/mocks"
	internalStorage "github.com/bejaneps/rollee-assignment/test2/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestUpsertWord(t *testing.T) {
	testCases := map[string]struct {
		setup       func() *mocks.Storage
		word        string
		expectedErr error
	}{
		"success": {
			setup: func() *mocks.Storage {
				storage := new(mocks.Storage)

				storage.On(
					"UpsertWord",
					"hello",
				).Return(
					nil,
				)

				return storage
			},
			word: "hello",
		},
		"fail": {
			setup: func() *mocks.Storage {
				storage := new(mocks.Storage)

				storage.On(
					"UpsertWord",
					"hello",
				).Return(
					errors.New("random error"),
				)

				return storage
			},
			word: "hello",
			expectedErr: errors.New(
				"random error: failed to upsert word",
			),
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			storage := tt.setup()

			svc := service.New(storage)

			err := svc.UpsertWord(tt.word)
			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			storage.AssertExpectations(t)
		})
	}
}

func TestGetMostFrequentWord(t *testing.T) {
	testCases := map[string]struct {
		setup        func() *mocks.Storage
		word         string
		expectedWord string
		expectedErr  error
	}{
		"success": {
			setup: func() *mocks.Storage {
				storage := new(mocks.Storage)

				storage.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"hello",
					nil,
				)

				return storage
			},
			word:         "hel",
			expectedWord: "hello",
		},
		"fail": {
			setup: func() *mocks.Storage {
				storage := new(mocks.Storage)

				storage.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"",
					errors.New("random error"),
				)

				return storage
			},
			word:         "hel",
			expectedWord: "",
			expectedErr: errors.New(
				"random error: failed to get most frequent word",
			),
		},
		"fail-word-not-found": {
			setup: func() *mocks.Storage {
				storage := new(mocks.Storage)

				storage.On(
					"GetMostFrequentWord",
					"hel",
				).Return(
					"",
					internalStorage.ErrWordNotFound,
				)

				return storage
			},
			word:         "hel",
			expectedWord: "",
			expectedErr:  internalStorage.ErrWordNotFound,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			storage := tt.setup()

			svc := service.New(storage)

			word, err := svc.GetMostFrequentWord(tt.word)
			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedWord, word)

			storage.AssertExpectations(t)
		})
	}
}
