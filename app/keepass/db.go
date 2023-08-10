package keepass

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/tobischo/gokeepasslib/v3"
	"golang.org/x/term"
)

var db *gokeepasslib.Database

func DbIns() *gokeepasslib.Database {
	if db != nil {
		return db
	}
	kpfile := os.Getenv("KEEPASS_DB")
	if kpfile == "" {
		println("KEEPASS_DB is not found")
		os.Exit(2)
	}
	file, _ := os.Open(kpfile)

	db = gokeepasslib.NewDatabase()

	password := getDbPassword()
	db.Credentials = gokeepasslib.NewPasswordCredentials(password)
	if err := gokeepasslib.NewDecoder(file).Decode(db); err != nil {
		log.Fatalf("Fatal decode err: %s", err)
	}

	db.UnlockProtectedEntries()
	// defer db.LockProtectedEntries()
	return db
}

func getDbPassword() string {
	pass := os.Getenv("KEEPASS_PASS")
	if pass != "" {
		return pass
	}

	password := InputPassword("Input keepass master password")
	return password
}

func GetUserPassword(username string) string {
	password := ""
	for _, group := range db.Content.Root.Groups {
		// fmt.Println("group:", group.Name)
		// fmt.Println("group:", group.EnableSearching)
		for _, entry := range group.Entries {
			// fmt.Println(entry.GetTitle())
			val := entry.Get("Username")
			if val == nil {
				continue
			}
			if val.Value.Content == username {
				return entry.GetPassword()
			}
		}
	}
	return password
}

// InputPrompt asks for a string value using the label
func InputPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+": ")
		s, _ = r.ReadString('\n')
		if s != "\n" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func InputPassword(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+": ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		s = string(bytePassword)
		if s != "" {
			break
		}
	}
	println("")
	return s
}
