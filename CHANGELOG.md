### What's changed in v0.0.2

* fix: bump TERRAFORM_PROVIDER_VERSION 0.2.0 → 0.2.1 (by @patrickleet)

  Picks up the upstream provider's empty-state.ID guard fix in Read
  (terraform-provider-listmonk v0.2.1) — without that, fresh
  UserRole/User MRs fail their first observe with 'Invalid id in
  state: strconv.ParseInt parsing empty string'.

  Provider schema is unchanged between v0.2.0 and v0.2.1, so
  config/schema.json doesn't need to regenerate.


See full diff: [v0.0.1...v0.0.2](https://github.com/hops-ops/provider-listmonk/compare/v0.0.1...v0.0.2)
