---
page_title: "hosting_email_account Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an email account.
---

# hosting_email_account (Resource)

Manages an email account associated with a hostname/FQDN.

## Example Usage

```hcl
resource "hosting_email_account" "info" {
  fqdn_id      = hosting_fqdn.main.id
  address      = "info"
  password     = var.email_password
  quota_bytes  = 5368709120  # 5 GB
}
```

## Schema

### Required

- `fqdn_id` (String) Hostname/FQDN ID this email account belongs to. Changing this forces a new resource.
- `address` (String) Email address (local part only, domain from FQDN). Changing this forces a new resource.
- `password` (String, Sensitive) Account password.

### Optional

- `display_name` (String) Display name.
- `quota_bytes` (Number) Mailbox quota in bytes. Default: `1073741824` (1 GB).

### Read-Only

- `id` (String) Email account ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_email_account.info em_abc123
```
