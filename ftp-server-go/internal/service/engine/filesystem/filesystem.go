package filesystem

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name        string
	Size        int64
	Permissions string
}

type DirEntry struct {
	Name        string
	Permissions string
}

type Filesystem interface {
	OpenRead(ctx context.Context, path string) (io.ReadCloser, error)
	OpenWrite(ctx context.Context, path string, append bool) (io.WriteCloser, error)
	Stat(ctx context.Context, path string) (FileInfo, error)
	Remove(ctx context.Context, path string) error
	Rename(ctx context.Context, from, to string) error

	Mkdir(ctx context.Context, path string) error
	Rmdir(ctx context.Context, path string) error
	ReadDir(ctx context.Context, path string) ([]DirEntry, error)
}

type filesystem struct {
	root string
}

func New(root string) Filesystem {
	return &filesystem{
		root: root,
	}
}

func (f *filesystem) OpenRead(ctx context.Context, path string) (io.ReadCloser, error) {
	fullPath := filepath.Join(f.root, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *filesystem) OpenWrite(ctx context.Context, path string, append bool) (io.WriteCloser, error) {
	fullPath := filepath.Join(f.root, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *filesystem) Stat(ctx context.Context, path string) (FileInfo, error) {
	fullPath := filepath.Join(f.root, path)
	file, err := os.Stat(fullPath)
	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Name:        file.Name(),
		Permissions: file.Mode().String(),
		Size:        file.Size(),
	}, nil
}

func (f *filesystem) Remove(ctx context.Context, path string) error {
	fullPath := filepath.Join(f.root, path)
	return os.Remove(fullPath)
}

func (f *filesystem) Rename(ctx context.Context, from, to string) error {
	fullPathFrom := filepath.Join(f.root, from)
	fullPathTo := filepath.Join(f.root, to)
	return os.Rename(fullPathFrom, fullPathTo)
}

func (f *filesystem) ReadDir(ctx context.Context, path string) ([]DirEntry, error) {
	fullPath := filepath.Join(f.root, path)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}
	var de []DirEntry
	for _, entry := range entries {
		de = append(de, DirEntry{
			Name:        entry.Name(),
			Permissions: entry.Type().String(),
		})
	}
	return de, nil
}

func (f *filesystem) Mkdir(ctx context.Context, path string) error {
	fullPath := filepath.Join(f.root, path)
	return os.Mkdir(fullPath, 0755)
}

func (f *filesystem) Rmdir(ctx context.Context, path string) error {
	fullPath := filepath.Join(f.root, path)
	return os.RemoveAll(fullPath)
}
