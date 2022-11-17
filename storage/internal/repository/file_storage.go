package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FileClient struct {
	db *sqlx.DB
}

func NewFileClient(db *sqlx.DB) *FileClient {
	return &FileClient{
		db: db,
	}
}

func (c *FileClient) UploadDB(fileName, Key string) error {

	_, err := c.db.Exec(`INSERT INTO files (data, key) VALUES ($1, $2)`, fmt.Sprintf(`{"bucket": "test", "file": "%s"}`, fileName), Key)
	if err != nil {
		return err
	}
	return nil
}
func (c *FileClient) DeleteFile(key string) error {
	_, err := c.db.Exec("DELETE from files where key = $1", key)
	if err != nil {
		return err
	}
	return nil
}

func (c *FileClient) GetKey(key string) (string, error) {
	var fileName string

	err := c.db.Get(&fileName, `select data->>'file' from files where key=$1`, key)
	return fileName, err
}
