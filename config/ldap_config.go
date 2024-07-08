package config

import (
	"fmt"
	"thegraduate-server/interfaces"
)

type LDAPConfig struct {
	LdapDomain                   string
	UseTLS                       bool
	TCPPort                      map[string]int
	BaseDN                       string
	GroupsDN                     string
	LdapUserAttribute            string
	LdapGroupAttribute           string
	LdapSearchAttribute          []string
	LdapMemberOfGroupsIdentifier string
	Groups                       []string
	GroupByUsers                 map[string][]string
	UserNameAndGroup             map[string][]string
}

func NewLDAPConfig() interfaces.LDAPConfigInterface {
	return &LDAPConfig{
		LdapDomain:                   "pdc.yarsi.ac.id",
		UseTLS:                       false,
		TCPPort:                      map[string]int{"default": 389, "tls": 636},
		BaseDN:                       "dc=yarsi,dc=ac,dc=id",
		GroupsDN:                     "ou=Groups,dc=yarsi,dc=ac,dc=id",
		LdapUserAttribute:            "uid",
		LdapGroupAttribute:           "cn",
		LdapSearchAttribute:          []string{"dn", "cn", "displayName", "telephoneNumber", "description", "homePhone", "street", "homePostalAddress", "title"},
		LdapMemberOfGroupsIdentifier: "cn",
		Groups:                       []string{},
		GroupByUsers:                 map[string][]string{},
		UserNameAndGroup:             map[string][]string{},
	}
}

func (c *LDAPConfig) GetLdapURL() string {
	port := c.TCPPort["default"]
	if c.UseTLS {
		port = c.TCPPort["tls"]
	}
	return fmt.Sprintf("%s:%d", c.LdapDomain, port)
}
func (c *LDAPConfig) IsTLS() bool {
	return c.UseTLS
}

func (c *LDAPConfig) GetBaseDN() string {
	return c.BaseDN
}
func (c *LDAPConfig) GetGroupsDN() string {
	return c.GroupsDN
}

func (c *LDAPConfig) GetLdapGroupAttribute() string {
	return c.LdapGroupAttribute
}

func (c *LDAPConfig) GetLdapUserAttribute() string {
	return c.LdapUserAttribute
}

func (c *LDAPConfig) GetLdapSearchAttribute() []string {
	return c.LdapSearchAttribute
}

func (c *LDAPConfig) GetLdapMemberOfGroupsIdentifier() string {
	return c.LdapMemberOfGroupsIdentifier
}

func (c *LDAPConfig) SetGroup(groupName string, users []string) {
	if !contains(c.Groups, groupName) {
		c.Groups = append(c.Groups, groupName)
		c.GroupByUsers[groupName] = users
	}
}

func (c *LDAPConfig) SetUserAndGroup(userName string, groupName string) {
	c.UserNameAndGroup[userName] = append(c.UserNameAndGroup[userName], groupName)
}

func (c *LDAPConfig) GetGroups() []string {
	return c.Groups
}

func (c *LDAPConfig) GetGroupByUsers() map[string][]string {
	return c.GroupByUsers
}

func (c *LDAPConfig) GetRoleByUserName(userName string) []string {
	return c.UserNameAndGroup[userName]
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
