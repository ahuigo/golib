package main
import (
	"strings"
	"testing"

	"github.com/pkg/errors"

	"fmt"
	"regexp"

	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
)

const (
	ldapServer       = "ldap://127.0.0.1:389"
	departmentDN     = "CN=department,OU=Account,DC=mycompany"
	departmentPasswd = "password"
	personDN         = "OU=All Users,DC=mycompany"
)

func TestListUserNames(t *testing.T) {
	usernames, err := GetAllUserNames()
	if err != nil {
		t.Fatal(err)
	}
	if len(usernames) == 0 {
		t.Fatalf("empty usernames!")
	}
	fmt.Println(usernames)
}

func getLdapConn() (l *ldap.Conn, err error) {
	l, err = ldap.DialURL(ldapServer)
	if err != nil {
		return
	}
	// bind  BaseDN
	err = l.Bind(departmentDN, departmentPasswd)

	return
}

func searchPerson(l *ldap.Conn, username string) (sr *ldap.SearchResult, err error) {
	personAttribute := []string{"sAMAccountName"}
	searchRequest := ldap.NewSearchRequest(
		personDN,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		fmt.Sprintf("(&(objectclass=person)(CN=%s))", username),
		personAttribute,
		nil,
	)
	sr, err = l.Search(searchRequest)
	return
}

var patternPersonDN = regexp.MustCompile(`CN=([a-z][\w\-\.]+),`)

func GetAllUserNames() (usernames []string, err error) {
	l, err := getLdapConn()
	if err != nil {
		return nil, err
	}
	defer l.Close()

	personDN := viper.GetString("ldap.baseDn")
	personAttribute := []string{"sAMAccountName"}
	searchRequest := ldap.NewSearchRequest(
		personDN,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		("(&(objectclass=person))"),
		personAttribute,
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	usernames = make([]string, 0, len(sr.Entries))
	for _, entry := range sr.Entries {

		res := patternPersonDN.FindStringSubmatch(entry.DN)
		if len(res) > 0 {
			usernames = append(usernames, res[1])
		}
	}

	return usernames, nil
}

// IsLdapUsername check on if the user name is belonging to the email user
func IsLdapUsername(username string) (isExisted bool, err error) {
	l, err := getLdapConn()
	if err != nil {
		return
	}
	defer l.Close()
	isExisted = true
	sr, err := searchPerson(l, username)
	if sr == nil || len(sr.Entries) == 0 {
		isExisted = false
	}
	return
}

//CheckUserPassword Check LDAP USER PASSWORD
func CheckUserPassword(username string, password string) (err error) {

	l, err := getLdapConn()
	if err != nil {
		// logger.Error("Get Ldap Connection failed: ", err)
		return
	}

	sr, err := searchPerson(l, username)
	if err != nil {
		// logger.Error("Search Person failed:", err)
		return
	}

	if sr.Entries == nil || len(sr.Entries) == 0 {
		return errors.New("invaid username and password")
	}

	entryDn := sr.Entries[0].DN

	// verify password
	err = l.Bind(entryDn, password)
	if err != nil {
		if strings.Contains(err.Error(), "LDAP Result Code 49") {
			return errors.New("invaid username and password")
		}
		return errors.Wrapf(err, "invaid username and password")
	}
	return
}
