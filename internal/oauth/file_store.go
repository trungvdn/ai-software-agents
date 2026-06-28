package oauth

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

type FileTokenStore struct {
	path string
}

func NewFileTokenStore(path string) *FileTokenStore {
	return &FileTokenStore{path: path}
}

func (s *FileTokenStore) Save(
	ctx context.Context,
	token *oauth2.Token,
) error {
	if s == nil || s.path == "" {
		return nil
	}

	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}

	if err := os.WriteFile(s.path, data, 0o600); err != nil {
		return err
	}

	return nil
}

func (s *FileTokenStore) Load(
	ctx context.Context,
) (*oauth2.Token, error) {
	if s == nil || s.path == "" {
		return nil, nil
	}

	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
