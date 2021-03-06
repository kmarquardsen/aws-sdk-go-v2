package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// WriteJSON marshals the given data as JSON and writes it to the specified location.
func WriteJSON(data interface{}, root, dir, name string) error {
	filePath := filepath.Join(root, dir, name+".json")
	changeBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return WriteFile(changeBytes, filePath, false)
}

// WriteFile writes the given bytes to the file at path, creating one if necessary. If appendTo is true, then the data
// will be prepended to the top of the file.
func WriteFile(data []byte, path string, appendTo bool) error {
	if appendTo {
		exists, err := FileExists(path, false)
		if err != nil {
			return err
		}

		if exists {
			existingData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			data = append(data, existingData...)
		}
	}

	return ioutil.WriteFile(path, data, 0644)
}

// FileExists returns whether or not a file exists at the specified path. If dir is true, FileExists checks for the existence
// of the specified directory.
func FileExists(path string, dir bool) (bool, error) {
	if f, err := os.Stat(path); err == nil {
		if f.IsDir() != dir {
			return false, nil
		}

		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// FindFile recursively searches upwards from the current directory to the filesystem root for the specified file.
func FindFile(fileName string, dir bool) (string, error) {
	currPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to find file: %v", err)
	}

	for {
		if currPath == string(os.PathSeparator) || filepath.VolumeName(currPath) == currPath {
			return "", errors.New("failed to find file: reached filesystem root")
		}

		targetFilepath := filepath.Join(currPath, fileName)
		found, err := FileExists(targetFilepath, dir)
		if err != nil {
			return "", fmt.Errorf("failed to find file: %v", err)
		}

		if found {
			return targetFilepath, nil
		}

		// trimming trailing '/' causes filepath.Split to trim the last directory in currPath
		currPath = strings.TrimSuffix(currPath, string(os.PathSeparator))
		currPath, _ = filepath.Split(currPath)
	}
}

// ReplaceLine replaces any line in the file at the given path that begins with linePrefix with the given replacement
// string.
func ReplaceLine(path, linePrefix, replacement string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("couldn't replace file's line: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, linePrefix) {
			lines[i] = replacement
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("failed to write replacement file: %v", err)
	}

	return nil
}

// ExecAt runs the given Cmd with is working directory set to path.
func ExecAt(cmd *exec.Cmd, path string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Dir = path

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("couldn't run cmd %s: %v: %s", cmd.String(), err, stderr.String())
	}

	return stdout.Bytes(), nil
}
