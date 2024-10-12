package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ensureDir(dirname string) error {
    if !DirExists(dirname) {
        return os.MkdirAll(dirname, os.ModePerm)
    }
    return nil
}

func copyFile(srcFilePath string, dest io.Writer) error {
    file, err := os.Open(srcFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.Copy(dest, file)
    return err
}

func Zip(src, dest string) error {
	if !DirExists(src) {
		fmt.Printf("The given directory(%s) doesn't exist!\n", src)
		os.Exit(1)
	}

    if !strings.HasSuffix(dest, ".xapk") {
        dest += ".xapk"
    }

    zipFile := CreateFile(dest)
    defer zipFile.Close()

    writer := zip.NewWriter(zipFile)
    defer writer.Close()

    // Обход файлов директории
    err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Получаем относительный путь
        relPath, err := filepath.Rel(filepath.Dir(src), path)
        if err != nil {
            return err
        }

        if relPath == "." {
            return nil
        }

        if info.IsDir() {
            _, err := writer.Create(relPath + "/") // Для папок
            return err
        }

        zipEntry, err := writer.Create(relPath)
        if err != nil {
            return err
        }

        return copyFile(path, zipEntry)
    })

    return err
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
