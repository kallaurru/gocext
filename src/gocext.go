package src

import (
	"fmt"
	"path/filepath"
)

func MakeCommonInfrastructure(root string) {
	list := makeInternalPathsList()
	for _, path := range *list {
		fullPath := filepath.Join(root, path)
		fmt.Printf("Created path %s\n", fullPath)
	}
}

func makeInternalPathsList() *[]string {
	return &[]string {
		"Downloads", "Projects", "Yandex.Disk", "Storages",
	}
}
