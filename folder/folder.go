package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error)

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// example: feel free to change the data structure, if slice is not what you want
	folders   []Folder
	folderMap map[string]Folder
}

func NewDriver(folders []Folder) IDriver {
	folderMap := make(map[string]Folder)
	for _, folder := range folders {
		folderMap[folder.Name+folder.OrgId.String()] = folder
	}

	return &driver{
		folders:   folders,
		folderMap: folderMap, // Initialize the map
	}
}
