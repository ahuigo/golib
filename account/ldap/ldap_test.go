package main

import (
	"testing"
    _ "errors"
    "fmt"
    "github.com/go-ldap/ldap/v3"
)

/**
refer1: py-lib/account/ldap.py
refer2:
https://www.pgadmin.org/docs/pgadmin4/development/config_py.html#config-py
https://www.pgadmin.org/docs/pgadmin4/latest/ldap.html

  LDAP_SERVER_URI: "ldap://192.168.1.40:389"
  -----------------------------------------
3 ways to configure LDAP as follows (Choose anyone):

    # 1. Dedicated User binding

      LDAP_BIND_USER(bindDn): "CN=mydepartment,OU=Account,DC=mycompany,DC=ai"
      LDAP_BIND_PASSWORD: 'bindpass'

    # 2. Anonymous Binding
        LDAP_ANONYMOUS_BIND = False

    # 3. Bind as pgAdmin user
    LDAP_BASE_DN = '<Base-DN>'
        # AD example: (&(objectClass=user)(memberof=CN=MYGROUP,CN=Users,dc=example,dc=com))
        # OpenLDAP example: CN=Users,dc=example,dc=com
  ############################################
    LDAP_SEARCH_BASE_DN = 'OU=All Users,DC=mycompany,DC=ai'
        # Search ldap for further authentication (REQUIRED)
        # It can be optional while bind as pgAdmin user
        # e.g.: "OU=All Users,DC=mycompany,DC=ai"
    LDAP_SEARCH_FILTER = '(objectclass=*)'
        # Filter string for the user search.
        # For OpenLDAP, '(cn=*)' may well be enough.
        # For AD, you might use '(objectClass=user)' (REQUIRED) # e.g.  
            # LDAP_SEARCH_FILTER='(objectclass=person)' # alluser
    LDAP_USERNAME_ATTRIBUTE = '<User-id>'
        # The LDAP attribute containing user names. In OpenLDAP, this may be 'uid'
        # whilst in AD, 'sAMAccountName' might be appropriate. (REQUIRED) ??? 不对！！！
        # 应该是  LDAP_USERNAME_ATTRIBUTE='CN'
            # 最终合成"search_filter": '(&(objectclass=person)(CN=%s))',

    
    LDAP_SEARCH_SCOPE = 'SUBTREE'
        # Search scope for users (one of BASE, LEVEL or SUBTREE)
**/
func TestUserPass(t *testing.T){
    ldapServer := "ldap://192.168.1.40:389"
    bindUser := "CN=department,OU=Account,DC=mycompany,DC=ai" //bindUser
    bindPassword := "开发部password" //bind_pass
    personDN := "OU=All Users,DC=ahuigo,DC=com"
    personAttribute := []string{ "accountName"}
    username := "username"
    password := "user_password"

	l, err := ldap.DialURL(ldapServer)
    if err != nil {
		t.Fatal(err)
    }
    // bind  BaseDN
    err = l.Bind(bindUser, bindPassword)
    if err != nil {
		t.Fatal(err)
    }
    // search person
    searchRequest := ldap.NewSearchRequest(
		personDN,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
        fmt.Sprintf("(&(objectclass=person)(CN=%s))", username),
		personAttribute,
		nil,
    )
	sr, err := l.Search(searchRequest)
	if err != nil {
		t.Error(err)
	}

    if len(sr.Entries) == 0 {
        t.Error("user not exists")
    }

    entry_dn := sr.Entries[0].DN
    t.Logf("Search: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))
    t.Logf("%#v",  sr.Entries[0])
    t.Logf("dn:%s",  sr.Entries[0].DN)

    // verify entry_dn password
    err = l.Bind(entry_dn, password)
    if err != nil {
		t.Fatal("invalid password")
		t.Fatal(err)
    }

}

