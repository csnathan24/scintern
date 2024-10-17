package folder

import (
	"fmt"
)

// A method to move a subtree from one parent node to another, while maintaining the order of the children.
// The method should return the new folder structure once the move has occurred.
// Implement any necessary error handling (e.g. invalid paths, moving a node to a child of itself, moving folders to a different orgID, etc).
// There is no need to persist state, we can assume each method call will be independent of the previous one
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Create a map for folder lookups
	folderMap := make(map[string]*Folder)
	for i := range f.folders {
		folderMap[f.folders[i].Name] = &f.folders[i]
	}

	// Check if the source/ dest folder exists and if source = destination
	sourceFolder, exists := folderMap[name]
	if !exists {
		return nil, fmt.Errorf("source folder '%s' does not exist", name)
	}

	destFolder, exists := folderMap[dst]
	if !exists {
		return nil, fmt.Errorf("destination folder '%s' does not exist", dst)
	}

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

	// Create a new slice to hold the updated folder structure
	newFolders := make([]Folder, len(f.folders))
	copy(newFolders, f.folders) // Copy the original folder structure

	// Create the new path for the source folder
	newPath := destFolder.Paths + "." + sourceFolder.Name

	// Move the source folder in the new structure
	for i := range newFolders {
		if newFolders[i].Name == sourceFolder.Name {
			newFolders[i].Paths = newPath
			break
		}
	}

	// Move all children (if any) in the new structure
	for i := range newFolders {
		if isChildFolder(newFolders[i].Paths, sourceFolder.Paths) {
			// Create the new relative path for the child folder
			childRelativePath := newFolders[i].Paths[len(sourceFolder.Paths):]
			// Set the new path for the child folder
			newFolders[i].Paths = newPath + childRelativePath
		}
	}
	return newFolders, nil
}
