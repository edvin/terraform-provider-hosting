---
page_title: "hosting_dns_zone Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a DNS zone via domain verification.
---

# hosting_dns_zone (Resource)

Manages a DNS zone via domain verification. On creation, starts a TXT-based verification. Once the TXT record is in place, the zone is auto-created. Works like AWS ACM certificate validation -- output the verification records and create them via your DNS provider.

## Example Usage

```hcl
resource "hosting_dns_zone" "main" {
  domain = "example.com"
}

# Create the verification TXT record at your external DNS provider
# Then the zone will activate automatically
output "verification_record" {
  value = {
    name  = hosting_dns_zone.main.verification_txt_name
    value = hosting_dns_zone.main.verification_txt_value
  }
}
```

## Schema

### Required

- `domain` (String) Domain name (e.g. `example.com`). Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.

### Read-Only

- `id` (String) Zone ID (available after verification completes).
- `managed` (Boolean) Whether the zone is managed by the platform's nameservers.
- `status` (String) Current status (`pending_verification` or `active`).
- `verification_txt_name` (String) TXT record name to create at your DNS provider for verification. Empty after verification completes.
- `verification_txt_value` (String) TXT record value for verification. Empty after verification completes.

## Import

```shell
terraform import hosting_dns_zone.main dz_abc123
```
