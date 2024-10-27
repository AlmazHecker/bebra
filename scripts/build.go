package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const appName, outputDir = "bebra", "bin"

var platforms = []struct {
	os, arch, format string
}{{"linux", "amd64", "tar.gz"}, {"windows", "amd64", "zip"}}

func main() {
	os.MkdirAll(outputDir, os.ModePerm)

	for _, p := range platforms {
		ext := ""
		if p.os == "windows" {
			ext = ".exe"
		}
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s-%s-%s%s", appName, p.os, p.arch, ext))

		if err := exec.Command("go", "build", "-o", outputFile).Run(); err != nil {
			fmt.Printf("Build failed for %s/%s: %v\n", p.os, p.arch, err)
			continue
		}

		fmt.Printf("Built %s for %s/%s\n", appName, p.os, p.arch)
		archiveName := filepath.Join(outputDir, fmt.Sprintf("%s-%s-%s.%s", appName, p.os, p.arch, p.format))

		if err := createArchive(outputFile, archiveName, p.format); err != nil {
			fmt.Printf("Error creating archive: %v\n", err)
		}
		os.Remove(outputFile)
	}
	fmt.Println("All builds completed.")
}

func createArchive(file, archivePath, format string) error {
	if format == "zip" {
		return createZip(file, archivePath)
	}
	return createTarGz(file, archivePath)
}

func createZip(file, zipPath string) error {
	zf, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zf.Close()

	w := zip.NewWriter(zf)
	defer w.Close()

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	wr, err := w.Create(filepath.Base(file))
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, f)
	return err
}

func createTarGz(file, tarGzPath string) error {
	tf, err := os.Create(tarGzPath)
	if err != nil {
		return err
	}
	defer tf.Close()

	gz := gzip.NewWriter(tf)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	h, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	h.Name = filepath.Base(file)

	if err := tw.WriteHeader(h); err != nil {
		return err
	}

	_, err = io.Copy(tw, f)
	return err
}
