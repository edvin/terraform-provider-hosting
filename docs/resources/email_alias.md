---
page_title: "hosting_email_alias Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an email alias.
---

# hosting_email_alias (Resource)

Manages an email alias for an existing email account.

## Example Usage

```hcl
resource "hosting_email_alias" "contact" {
  email_account_id = hosting_email_account.info.id
  address          = "contact"
}
```

## Schema

### Required

- `email_account_id` (String) Parent email account ID. Changing this forces a new resource.
- `address` (String) Alias address. Changing this forces a new resource.

### Read-Only

- `id` (String) Alias ID.
- `status` (String) Current status.

## Import

Import uses the format `email_account_id/alias_id`:

```shell
terraform import hosting_email_alias.contact em_abc123/ema_xyz789
```
