package inmemory

import (
	"sort"
	"strings"
	"sync"

	"github.com/bejaneps/rollee-assignment/test2/internal/storage"
)

// Storage represents in-memory storage.
type Storage struct {
	// to protect map from concurrent read-writes
	mtx *sync.RWMutex

	// count of words and their frequency
	wordsCount map[string]int
}

// New instantiates an instance of in-memory storage.
func New() *Storage {
	return &Storage{
		mtx:        &sync.RWMutex{},
		wordsCount: make(map[string]int),
	}
}

// UpsertWord updates a frequency of word in the map
// or inserts a word if it doesn't exist in map.
func (s *Storage) UpsertWord(word string) error {
	s.mtx.RLock()
	_, ok := s.wordsCount[word]
	s.mtx.RUnlock()
	if ok {
		s.mtx.Lock()
		s.wordsCount[strings.ToLower(word)]++
		s.mtx.Unlock()

		return nil
	}

	s.mtx.Lock()
	s.wordsCount[strings.ToLower(word)] = 1
	s.mtx.Unlock()

	return nil
}

// GetMostFrequentWord returns most frequest
// word stored in the storage.
func (s *Storage) GetMostFrequentWord(
	beginning string,
) (string, error) {
	keys := make([]string, 0, len(s.wordsCount))

	s.mtx.RLock()
	for key := range s.wordsCount {
		keys = append(keys, key)
	}
	s.mtx.RUnlock()

	s.mtx.RLock()
	sort.Slice(keys, func(i, j int) bool {
		return s.wordsCount[keys[i]] > s.wordsCount[keys[j]]
	})
	s.mtx.RUnlock()

	s.mtx.RLock()
	defer s.mtx.RUnlock()
	for _, key := range keys {
		if strings.HasPrefix(key, beginning) {
			return key, nil
		}
	}

	return "", storage.ErrWordNotFound
}
