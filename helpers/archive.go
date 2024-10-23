package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ensureDir(dirname string) error {
    if !DirExists(dirname) {
        return os.MkdirAll(dirname, os.ModePerm)
    }
    return nil
}


func Unzip(src, dest string) error {
	if !FileExists(src) {
		fmt.Printf("The given file(%s) doesn't exist!\n", src)
		os.Exit(1)
	}
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
		fmt.Println(f.FileInfo())
        fpath := filepath.Join(dest, f.Name)

        rel, err := filepath.Rel(dest, fpath)
        if err != nil || rel == "" || rel[0] == '.' {
            return fmt.Errorf("invalid file path: %s", fpath)
        }

        if f.FileInfo().IsDir() {
            if err = ensureDir(fpath); err != nil {
                return err
            }
            continue
        }

        if err = ensureDir(filepath.Dir(fpath)); err != nil {
            return err
        }

        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return err
        }

        if _, err = io.Copy(outFile, rc); err != nil {
            return err
        }

		outFile.Close()
    }
    return nil
}
