package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// Helper function to compare two slices of Folder
func equalsFolders(a, b []folder.Folder) bool {
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

// Use uuid instead of "org1" for integrity
func Test_folder_MoveFolder(t *testing.T) {
	tests := []struct {
		name    string
		folders []folder.Folder
		move    string
		dst     string
		want    []folder.Folder
		errMsg  string
	}{
		{
			name: "Source folder does not exist",
			folders: []folder.Folder{
				{Name: "bravo", Paths: "bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move:   "nonexistent",
			dst:    "bravo",
			want:   nil,
			errMsg: "source folder 'nonexistent' does not exist",
		},
		{
			name: "Destination folder does not exist",
			folders: []folder.Folder{
				{Name: "bravo", Paths: "bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move:   "bravo",
			dst:    "nonexistent",
			want:   nil,
			errMsg: "destination folder 'nonexistent' does not exist",
		},
		{
			name: "Move folder to a different organization",
			folders: []folder.Folder{
				{Name: "bravo", Paths: "bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
			},
			move:   "bravo",
			dst:    "golf",
			want:   nil,
			errMsg: "cannot move a folder to a different organization",
		},
		{
			name: "Move folder to itself",
			folders: []folder.Folder{
				{Name: "bravo", Paths: "bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move:   "bravo",
			dst:    "bravo",
			want:   nil,
			errMsg: "cannot move a folder to itself",
		},
		{
			name: "Move folder to a child of itself",
			folders: []folder.Folder{
				{Name: "bravo", Paths: "bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move:   "bravo",
			dst:    "charlie",
			want:   nil,
			errMsg: "cannot move a folder to a child of itself",
		},
		{
			name: "valid example 1",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move: "bravo",
			dst:  "delta",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			errMsg: "",
		},
		{
			name: "valid example 2",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move: "bravo",
			dst:  "golf",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "golf.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			errMsg: "",
		},
		{
			name: "valid example with nested folder - example 1 + beta in charlie",
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "beta", Paths: "alpha.bravo.charlie.beta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
			move: "bravo",
			dst:  "delta",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "beta", Paths: "alpha.delta.bravo.charlie.beta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.FromStringOrNil("b1234567-b7c0-45a3-a6ae-9546248fb17b")},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil("a1234567-b7c0-45a3-a6ae-9546248fb17a")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.move, tt.dst)

			// Check for errors
			if err != nil {
				if err.Error() != tt.errMsg {
					t.Errorf("expected error %v, got %v", tt.errMsg, err.Error())
				}
				return
			}

			// Check if the results match the expected output
			if !equalsFolders(got, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
