package interfaces

type LDAPConfigInterface interface {
	GetLdapURL() string
	IsTLS() bool
	GetBaseDN() string
	GetGroupsDN() string
	GetLdapGroupAttribute() string
	GetLdapUserAttribute() string
	GetLdapSearchAttribute() []string
	GetLdapMemberOfGroupsIdentifier() string
	SetGroup(groupName string, users []string)
	SetUserAndGroup(userName string, groupName string)
	GetGroups() []string
	GetGroupByUsers() map[string][]string
	GetRoleByUserName(userName string) []string
}
