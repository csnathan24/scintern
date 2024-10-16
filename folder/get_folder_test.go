package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// Tests are structured based off name inputs
// Use uuid instead of "org1" for integrity

// For empty string as argument to GetAllChildFolders
func Test_folder_GetAllChildFolders_InvalidNames(t *testing.T) {
	testCases := []struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		errMsg  string
	}{
		{
			name:  "Missing name but valid orgID",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{Name: "", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want:   []folder.Folder{},
			errMsg: "invalid name: folder name cannot be empty",
		},
	}

	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			driver := folder.NewDriver(tests.folders)
			result, error := driver.GetAllChildFolders(tests.orgID, "")

			if error != nil {
				if error.Error() != tests.errMsg {
					t.Errorf("expected error %v, got %v", tests.errMsg, error.Error())
				}
				return
			}

			if !equal(result, tests.want) {
				t.Errorf("GetAllChildFolders() = %v, want %v", result, tests.want)
			}
		})
	}
}

// invalid name input
func Test_folder_GetAllChildFolders_ExistingNames(t *testing.T) {
	testsCases := []struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		errMsg  string
	}{
		{
			name:  "Non-existent Name with valid orgID",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want:   []folder.Folder{},
			errMsg: "folder 'non_existent_folder' does not exist in the specified organization",
		},
	}

	for _, tests := range testsCases {
		t.Run(tests.name, func(t *testing.T) {
			driver := folder.NewDriver(tests.folders)
			result, error := driver.GetAllChildFolders(tests.orgID, "non_existent_folder")

			if error != nil {
				if error.Error() != tests.errMsg {
					t.Errorf("expected error %v, got %v", tests.errMsg, error.Error())
				}
				return
			}

			if !equal(result, tests.want) {
				t.Errorf("GetAllChildFolders() = %v, want %v", result, tests.want)
			}
		})
	}
}

// Test suite where alpha is the name input for GetAllChildFolders
func Test_folder_GetAllChildFolders_Valid(t *testing.T) {
	testCases := []struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		errMsg  string
	}{
		{
			name:  "Missing orgID but valid name",
			orgID: uuid.Nil,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want:   []folder.Folder{},
			errMsg: "invalid orgID: orgID cannot be nil",
		},
		{
			name:  "Valid name but orgID doesn't exist",
			orgID: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b"),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want:   []folder.Folder{},
			errMsg: "folder 'alpha' does not exist in the specified organization",
		},
		{
			name:  "No child folders",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want:   []folder.Folder{},
			errMsg: "",
		},
		{
			name:  "Valid child folder - single",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"), // Valid orgID
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			errMsg: "",
		},
		{
			name:  "Valid child folders - multiple",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},       // Unrelated folder but same orgID
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")}, // Different organisation
			},
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			errMsg: "",
		},
		{
			name:  "Multiple folders with same name in different organizations",
			orgID: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},  // org1
				{Name: "alpha", Paths: "alpha2", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")}, // org2
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			errMsg: "",
		},
	}

	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			driver := folder.NewDriver(tests.folders)
			result, error := driver.GetAllChildFolders(tests.orgID, "alpha")

			if error != nil {
				if error.Error() != tests.errMsg {
					t.Errorf("expected error %v, got %v", tests.errMsg, error.Error())
				}
				return
			}

			if !equal(result, tests.want) {
				t.Errorf("GetAllChildFolders() = %v, want %v", result, tests.want)
			}
		})
	}
}

// Helper function to compare two slices of Folder
func equal(a, b []folder.Folder) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
