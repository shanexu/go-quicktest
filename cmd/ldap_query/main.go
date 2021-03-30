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
	host := flag.String("host", "127.0.0.1", "127.0.0.1")
	port := flag.Int("port", 389, "389")
	bindusername := flag.String("bindusername", "bindusername", "bind username")
	bindpassword := flag.String("bindpassword", "bindpassword", "bind password")
	baseDN := flag.String("basedn", "basedn", "dc=xusheng,dc=org")
	attributes := flag.String("attributes", "dn,cn,sAMAccountName,userPrincipalName", "dn,cn")

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

	searchRequest := ldap.NewSearchRequest(
		*baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectCategory=person)(objectClass=user)(!(userAccountControl:1.2.840.113556.1.4.803:=2))(userPrincipalName=*@sumscope.com))",
		strings.Split(*attributes, ","),
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	us := make([]User, 0, len(sr.Entries))
	for _, entry := range sr.Entries {

		if strings.Contains(entry.DN, "离职人员") ||
			strings.Contains(entry.DN, "功能性用户") ||
			strings.Contains(entry.DN, "公司公共邮箱") {
			continue
		}
		u := User{
			Username: entry.GetAttributeValue("sAMAccountName"),
			Mail:     entry.GetAttributeValue("userPrincipalName"),
			Comment:  entry.DN,
		}
		us = append(us, u)
		b, _ := json.Marshal(u)
		fmt.Println(string(b))
	}
}

type User struct {
	Username string
	Mail     string
	Comment  string
}
