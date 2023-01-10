# URIs, URNs, URLs

- https://www.rfc-editor.org/rfc/rfc3986
- https://github.com/coopnorge/coop-playbook/issues/352
- schemes: https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml

```
URI = scheme ":" ["//" authority] path ["?" query] ["#" fragment]
authority = [userinfo "@"] host [":" port]
```

Example:

- http://user@example.com:80/some/path/to/resource.txt

Well known:

- http, https
- ftp
- file
- urn


Others:

https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml

- [`example:`](https://www.rfc-editor.org/rfc/rfc7595.html)
- [`tag:`](https://www.rfc-editor.org/rfc/rfc4151.html)

## URLs

Specifies a location of a resource

- file:///some/file.csv
- gcs://bucket/path/file.xls
- abfs://a/b/c/d.txt

- postgres://user@localhost/dbname
- mysql://user@localhost/dbname

- https://localhost:5001

Bad examples:

- https://playbook.internal.coop/knowledge_base/systems/idp.html?h=coop_id#new-idp-coop-tokens

```
{
  "https://id.coop.no/coop_id": "817621741",
  "iss": "https://login.test.coop/",
  "sub": "auth0|ed9139fd-12e8-4341-99dc-8a25b92d296c",
  "aud": "https://login.coop.no/user-management",
  "iat": 1634903544,
  "exp": 1634989944,
  "azp": "CKFAONqrsiK48oXEkrLu7rbA6fQHopZB",
  "scope": "delete:user create:user read:user update:user authenticate:user offline_access",
  "gty": "password"
}
```

- https://github.com/coopnorge/idp-auth0-config/blob/main/locals.production.tf#L485

```
      "https://file-pseudonymiser.coop.no" = {
        name   = "File Pseudonymiser"
        scopes = [{ value = "pseudonymise:entity", description = "Pseudonymise entity" }, { value = "depseudonymise:entity", description = "Depseudoynmise entity" }, { value = "create:context", description = "Create context" }, { value = "update:context", description = "Update context" }, { value = "withdraw_consent:entity", description = "Withdraw consent entity" }]
        authorized_clients = {
          "PseudonymiserApiM2MClient" = {
            requested_scopes = ["pseudonymise:entity", "depseudonymise:entity", "create:context", "update:context", "withdraw_consent:entity"]
          }
        }
      }
```

## URNs

Names a resource:

- `urn:...`
- `urn:fdc`
- `urn:uuid`
