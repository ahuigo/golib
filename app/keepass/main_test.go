package keepass

import (
	"fmt"
	"testing"
	// . "keepass"
)

func TestMain(t *testing.T) {
	db := DbIns()
	defer db.LockProtectedEntries()

	// Note: This is a simplified example and the groups and entries will depend on the specific file.
	// bound checking for the slices is recommended to avoid panics.
	for _, group := range db.Content.Root.Groups {
		fmt.Println("group:", group.Name)
		fmt.Println("group:", group.EnableSearching)
		entry := group.Entries[0]
		fmt.Println(entry.GetTitle())
		fmt.Println(entry.GetPassword())
	}

	fmt.Println("user1 passwd:", GetUserPassword("user1"))
}
