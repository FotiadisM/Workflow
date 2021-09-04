package repository

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const FilesFolder = "app_files/"

var FilesPath = filepath.Join("./", FilesFolder)

func (r Repository) PostFile(ctx context.Context, file io.Reader) (id string, err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	id = uid.String()

	b, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = os.WriteFile(filepath.Join(FilesPath, id), b, 0777)

	return
}
