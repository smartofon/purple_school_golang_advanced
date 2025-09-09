package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Storage представляет глобальное хранилище
type VerifyStorage struct {
	Data map[string]string
	File string
	mu   sync.RWMutex
}

var GlobalStorage *VerifyStorage

// Set добавляет значение по ключу
func (s *VerifyStorage) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Data[key] = value
}

// Get возвращает значение по ключу
func (s *VerifyStorage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.Data[key]
	return val, ok
}

func (s *VerifyStorage) Load() error {
	fileBytes, err := os.ReadFile(s.File)
	if err != nil {
		return fmt.Errorf("Error: Ошибка файла json для тестового письма: %s", err)
	}

	err = json.Unmarshal(fileBytes, &s)

	if err != nil {
		return fmt.Errorf("Ошибка чтения json для тестового письма: %s", err)
	}

	return nil
}

func (s *VerifyStorage) Save() error {

	s.mu.Lock()
	defer s.mu.Unlock()

	fo, err := os.Create(s.File)
	if err != nil {
		return fmt.Errorf("Ошибка создания json хранилища %s", err)
	}
	defer fo.Close()
	e := json.NewEncoder(fo)
	if err := e.Encode(s); err != nil {
		return fmt.Errorf("Ошибка сохранения json хранилища %s", err)
	}

	return nil
}
