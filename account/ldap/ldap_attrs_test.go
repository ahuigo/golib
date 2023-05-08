package main

import (
	_ "errors"
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
)

func GetAllAttributes() (attrs []string, err error) {
	l, err := getLdapConn()
	if err != nil {
		return nil, err
	}
	defer l.Close()

	baseDn := viper.GetString("ldap.baseDn")
	// personAttribute := []string{"accountName"}
	personAttribute := []string(nil) // 获取所有属性，设置为nil
	searchRequest := ldap.NewSearchRequest(
		baseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		10, 0, false,
		// ("(&(objectclass=person))"),
		("(&(objectclass=*))"),
		personAttribute,
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	attrs = make([]string, 0, len(sr.Entries))
	for _, entry := range sr.Entries {
		// attrs = append(attrs, entry.DN)
		fmt.Printf("DN: %s\n", entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf("Attribute: %s; Values: %v\n", attr.Name, attr.Values)
		}
		fmt.Println()
	}

	return attrs, nil
}
