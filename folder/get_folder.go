package folder

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	return res
}

// A method to get all child folders of a given folder.
// The method should return a list of all child folders.
// Implement any necessary error handling (e.g. invalid orgID, invalid paths, etc).
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	// Safe practice input validation
	if orgID == uuid.Nil {
		return []Folder{}, fmt.Errorf("invalid orgID: orgID cannot be nil")
	}
	if name == "" {
		return []Folder{}, fmt.Errorf("invalid name: folder name cannot be empty")
	}

	// Finding parent folder using the precomputed map in folder.go
	parentKey := name + orgID.String()
	parentFolder, exists := f.folderMap[parentKey]
	if !exists {
		return []Folder{}, fmt.Errorf("folder '%s' does not exist in the specified organization", name)
	}

	// Retrieve child folders
	var childFolders []Folder
	for _, folder := range f.folders {
		// Check if the folder's path is a child of the parent folder's path
		if isChildFolder(folder.Paths, parentFolder.Paths) && folder.OrgId == orgID {
			childFolders = append(childFolders, folder)
		}
	}

	// Return empty folder if there are no child folders
	if len(childFolders) == 0 {
		return []Folder{}, nil
	}

	return childFolders, nil
}

// Helper function to check if one path is a child of another
func isChildFolder(childPath, parentPath string) bool {
	if parentPath == childPath {
		return false
	}
	return len(childPath) > len(parentPath) && childPath[len(parentPath)] == '.' && childPath[:len(parentPath)] == parentPath
}
