---
page_title: "hosting_email_forward Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an email forward rule.
---

# hosting_email_forward (Resource)

Manages an email forward rule for an existing email account.

## Example Usage

```hcl
resource "hosting_email_forward" "backup" {
  email_account_id = hosting_email_account.info.id
  destination      = "backup@gmail.com"
  keep_copy        = true
}
```

## Schema

### Required

- `email_account_id` (String) Parent email account ID. Changing this forces a new resource.
- `destination` (String) Destination email address. Changing this forces a new resource.

### Optional

- `keep_copy` (Boolean) Keep a copy in the original mailbox. Default: `true`.

### Read-Only

- `id` (String) Forward ID.
- `status` (String) Current status.

## Import

Import uses the format `email_account_id/forward_id`:

```shell
terraform import hosting_email_forward.backup em_abc123/emf_xyz789
```
