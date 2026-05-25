# Changelog

vnext generates and overwrites this file at each release-tag CI run.
A placeholder is committed for the first release so the
`workflow-simple-release` reusable workflow's `bodyFile: CHANGELOG.md`
parameter has something to read from.

## v0.0.1

Initial release — upjet-generated Crossplane provider for Listmonk.

Resources (cluster-scoped + namespaced variants of each):

- `settings.listmonk.crossplane.io.SecuritySettings`
- `settings.listmonk.crossplane.io.AppSettings`
- `identity.listmonk.crossplane.io.UserRole`
- `identity.listmonk.crossplane.io.User`

Generated from `hops-ops/terraform-provider-listmonk v0.2.0` via upjet v2.
