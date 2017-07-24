### RATIONALE

Because this code:

```golang
   client := ldap.LDAPClient{}
   if err := confix.Confix("Ldap", &cfg, &client); err != nil {
      slog.P("Confix failed with `%v'", err)
   }
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


DETAILS

Confix will return an error on: 

1. Type mis-match between your struct and their struct.
1. Unsettable members in their struct that your struct tries to set.
