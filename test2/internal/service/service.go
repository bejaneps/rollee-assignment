package service

import (
	"github.com/bejaneps/rollee-assignment/test2/internal/storage"
	"github.com/pkg/errors"
)

// Storage represents a persistent or temporary
// memory where words will be saved.
type Storage interface {
	// UpsertWord updates frequency of word count
	// or inserts word in the storage
	UpsertWord(word string) error

	// GetMostFrequentWord returns most frequest
	// word stored in the storage
	GetMostFrequentWord(beginning string) (string, error)
}

// Service represents a business layer of MFW
type Service struct {
	storage Storage
}

func New(s Storage) *Service {
	service := &Service{
		storage: s,
	}

	return service
}

// UpsertWord updates frequency of word count or
// inserts word in the storage.
func (s *Service) UpsertWord(word string) error {
	err := s.storage.UpsertWord(word)
	if err != nil {
		return errors.Wrap(err, "failed to upsert word")
	}

	return nil
}

func (s *Service) GetMostFrequentWord(
	beginning string,
) (string, error) {
	word, err := s.storage.GetMostFrequentWord(beginning)
	if errors.Is(err, storage.ErrWordNotFound) {
		return "", storage.ErrWordNotFound
	} else if err != nil {
		return "", errors.Wrap(
			err,
			"failed to get most frequent word",
		)
	}

	return word, nil
}
