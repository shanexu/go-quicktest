package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func main() {

	host := flag.String("host","127.0.0.1", "127.0.0.1")
	port := flag.Int("port", 389, "389")
	username := flag.String("username", "username", "username")
	password := flag.String("password", "password", "password")
	bindusername := flag.String("bindusername", "bindusername", "bind username")
	bindpassword := flag.String("bindpassword", "bindpassword", "bind password")
	baseDN := flag.String("basedn", "basedn", "dc=xusheng,dc=org")
	attributes := flag.String("attributes", "dn,cn,sAMAccountName,userPrincipalName", "dn,cn")
	searchFilter := flag.String("searchfilter", "(cn=%s)", "(cn=%s)")

	flag.Parse()

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// First bind with a read only user
	err = l.Bind(*bindusername, *bindpassword)
	if err != nil {
		log.Fatal(err)
	}

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		*baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(*searchFilter, *username),
		strings.Split(*attributes, ","),
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	bs, _ := json.Marshal(sr.Entries[0].Attributes)
	fmt.Printf("%s", bs)
	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, *password)
	if err != nil {
		log.Fatal(err)
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(*bindusername, *bindpassword)
	if err != nil {
		log.Fatal(err)
	}
}
