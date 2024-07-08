package helper

import (
	"crypto/tls"
	"fmt"
	"log"
	"thegraduate-server/interfaces"

	"github.com/go-ldap/ldap/v3"
)

type AuthLdap struct {
	config       interfaces.LDAPConfigInterface
	ldapResource *ldap.Conn
}

func NewLdapConnection(config interfaces.LDAPConfigInterface) *AuthLdap {

	l, err := ldap.Dial("tcp", "pdc.yarsi.ac.id:389")

	if err != nil {
		log.Fatal("error connecting: ", err)
	}

	if config.IsTLS() {
		err = l.StartTLS(&tls.Config{})
		if err != nil {
			log.Fatalf("failed start TLS: ", err.Error())
		}
	}
	return &AuthLdap{
		config:       config,
		ldapResource: l,
	}
}

func (a *AuthLdap) Authenticate(userName, password string) map[string]interface{} {
	filterCriteria := fmt.Sprintf("(%s=%s)", a.config.GetLdapUserAttribute(), userName)
	searchAttributes := append(a.config.GetLdapSearchAttribute(), a.config.GetLdapUserAttribute())
	ldapSearchRequest := ldap.NewSearchRequest(
		a.config.GetBaseDN(),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filterCriteria,
		searchAttributes,
		nil,
	)

	ldapSearchResponse, err := a.ldapResource.Search(ldapSearchRequest)
	if err != nil {
		return nil
	}
	if len(ldapSearchResponse.Entries) == 0 {
		return nil
	}

	ldapBindRDN := ldapSearchResponse.Entries[0].DN

	err = a.ldapResource.Bind(ldapBindRDN, password)
	if err != nil {
		return nil
	}

	cn := ldapSearchResponse.Entries[0].GetAttributeValue("cn")
	dn := ldapSearchResponse.Entries[0].DN
	displayName := ldapSearchResponse.Entries[0].GetAttributeValue("displayname")
	idNik := ldapSearchResponse.Entries[0].GetAttributeValue("description")
	indexGroups := ldapSearchResponse.Entries[0].GetAttributeValue("title")

	arrGroup := map[string]string{"M": "Mahasiswa", "D": "Dosen", "S": "Staff"}
	memberOfGroups := arrGroup[indexGroups]

	return map[string]interface{}{
		"cn":          cn,
		"dn":          dn,
		"username":    userName,
		"role":        memberOfGroups,
		"displayname": displayName,
		"id_nik":      idNik,
		"logged_in":   true,
	}
}
