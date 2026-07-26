package store

import "context"

// Store はユーザーの永続化層(デモなのでインメモリ)。
type Store struct {
	dsn   string
	users map[string]string
}

func New(dsn string) *Store {
	return &Store{dsn: dsn, users: map[string]string{}}
}

func (s *Store) Get(ctx context.Context, id string) (string, error) {
	return s.users[id], nil
}

func (s *Store) List(ctx context.Context) ([]string, error) {
	out := make([]string, 0, len(s.users))
	for id := range s.users {
		out = append(out, id)
	}
	return out, nil
}

func (s *Store) Save(ctx context.Context, id, name string) error {
	s.users[id] = name
	return nil
}

func (s *Store) Deactivate(ctx context.Context, id string) error {
	delete(s.users, id)
	return nil
}
