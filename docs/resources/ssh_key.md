---
page_title: "hosting_ssh_key Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an SSH public key.
---

# hosting_ssh_key (Resource)

Manages an SSH public key for accessing tenant resources.

## Example Usage

```hcl
resource "hosting_ssh_key" "deploy" {
  tenant_id  = var.tenant_id
  name       = "deploy-key"
  public_key = file("~/.ssh/id_ed25519.pub")
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.
- `name` (String) Key name. Changing this forces a new resource.
- `public_key` (String) SSH public key content. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.

### Read-Only

- `id` (String) SSH key ID.
- `fingerprint` (String) Key fingerprint.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_ssh_key.deploy sk_abc123
```
