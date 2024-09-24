package internal

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss/tree"
)

const limitenclosure = 3

func getFiles(startDir string) ([]os.FileInfo, error) {
	dir, err := os.Open(startDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetTree(startDir string, enclosure int, limitFiles int) (*tree.Tree, error) {

	treeFiles := tree.Root(startDir)

	if startDir == "." {
		treeFiles.Root("⁜")
	} else {
		lastIndex := strings.LastIndex(startDir, "/")
		treeFiles.Root(startDir[lastIndex+1:])
	}

	if enclosure > limitenclosure {
		enclosure = limitenclosure
	}

	enclosure--

	files, err := getFiles(startDir)
	if err != nil {
		return nil, err
	}

	filesMap := make(map[string]int)

	for _, file := range files {

		typeFile := strings.ToLower(filepath.Ext(file.Name()))

		filesMap[typeFile]++

		if filesMap[typeFile] > limitFiles && startDir != "." {
			continue
		}

		if file.IsDir() && enclosure != 0 {

			subTree, err := GetTree(startDir+"/"+file.Name(), enclosure, limitFiles)
			if err != nil {
				return nil, err
			}

			treeFiles.Child(subTree)
		} else {
			treeFiles.Child(file.Name())
		}
	}

	if startDir != "." {
		hideFiles(filesMap, limitFiles, treeFiles)
	}

	return treeFiles, nil
}

func hideFiles(filesMap map[string]int, limitFiles int, treeFiles *tree.Tree) {
	for key, value := range filesMap {
		if value < limitFiles {
			continue
		}

		treeFiles.Child("and more " + strconv.Itoa(value-limitFiles) + " " + key)
	}
}
