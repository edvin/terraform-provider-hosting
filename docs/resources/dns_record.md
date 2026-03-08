---
page_title: "hosting_dns_record Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a DNS record within a zone.
---

# hosting_dns_record (Resource)

Manages a DNS record within a zone. Supports A, AAAA, CNAME, MX, TXT, SRV, and other record types.

## Example Usage

```hcl
resource "hosting_dns_record" "www" {
  zone_id = hosting_dns_zone.main.id
  type    = "CNAME"
  name    = "www"
  content = "myapp.example.com"
  ttl     = 3600
}

resource "hosting_dns_record" "mail" {
  zone_id  = hosting_dns_zone.main.id
  type     = "MX"
  name     = "@"
  content  = "mail.example.com"
  ttl      = 3600
  priority = 10
}
```

## Schema

### Required

- `zone_id` (String) DNS zone ID. Changing this forces a new resource.
- `type` (String) Record type (A, AAAA, CNAME, MX, TXT, SRV, etc.).
- `name` (String) Record name (e.g. `@`, `www`, `mail`).
- `content` (String) Record content (e.g. IP address, hostname).

### Optional

- `ttl` (Number) TTL in seconds. Default: `3600`.
- `priority` (Number) Priority (for MX, SRV records).

### Read-Only

- `id` (String) Record ID.
- `status` (String) Current status.

## Import

Import uses the format `zone_id/record_id`:

```shell
terraform import hosting_dns_record.www dz_abc123/dr_xyz789
```
