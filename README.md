### RATIONALE

Because this code:

```golang
	confix.Confix("Ldap", cfg, ldap.LDAPClient)
```

is better than this code:

```golang
	client := ldap.LDAPClient{
		Base:       cfg.LdapBase,
		Host:       cfg.LdapHost,
		Port:       cfg.LdapPort,
		UserFilter: cfg.LdapUserFilter,
	}
	if cfg.LdapGroupFilter != "" {
		client.GroupFilter = cfg.LdapGroupFilter
	}
	if cfg.LdapUseSSL != false {
		client.UseSSL = cfg.LdapUseSSL
		client.ServerName = cfg.LdapServerName
	}
	if cfg.LdapSkipTLS != false {
		client.SkipTLS = cfg.LdapSkipTLS
	}
	if cfg.LdapBindDN != "" {
		client.BindDN = cfg.LdapBindDN
	}
	if cfg.LdapBindPassword != "" {
		client.BindPassword = cfg.LdapBindPassword
	}
	if cfg.LdapAttributes != nil {
		client.Attributes = cfg.LdapAttributes
	}
```



