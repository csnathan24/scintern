package folder

import (
	"fmt"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Check if the source folder exists
	var sourceFolder *Folder
	for i, folder := range f.folders {
		if folder.Name == name {
			sourceFolder = &f.folders[i]
			break
		}
	}
	if sourceFolder == nil {
		return nil, fmt.Errorf("source folder '%s' does not exist", name)
	}

	// Check if the destination folder exists
	var destFolder *Folder
	for _, folder := range f.folders {
		if folder.Name == dst {
			destFolder = &folder
			break
		}
	}
	if destFolder == nil {
		return nil, fmt.Errorf("destination folder '%s' does not exist", dst)
	}

	// Check if the source and destination are the same
	if name == dst {
		return nil, fmt.Errorf("cannot move a folder to itself")
	}

	// Check if orgID for source and dest folder match
	if sourceFolder.OrgId != destFolder.OrgId {
		return nil, fmt.Errorf("cannot move a folder to a different organization")
	}

	// Check that the destination folder is not a child of the source folder
	if isChildFolder(destFolder.Paths, sourceFolder.Paths) {
		return nil, fmt.Errorf("cannot move a folder to a child of itself")
	}

	// Hold the original path before changing it
	originalPath := sourceFolder.Paths

	// Create the new path for the source folder
	newPath := destFolder.Paths + "." + sourceFolder.Name

	// Move the source folder
	sourceFolder.Paths = newPath

	// Move all children (if any)
	for i := range f.folders {
		if isChildFolder(f.folders[i].Paths, originalPath) {
			// Create the new relative path for the child folder
			childRelativePath := f.folders[i].Paths[len(originalPath):]

			// Set the new path for the child folder
			f.folders[i].Paths = newPath + childRelativePath
		}
	}

	return f.folders, nil
}
