---
page_title: "hosting_valkey_user Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a Valkey (Redis) user.
---

# hosting_valkey_user (Resource)

Manages a Valkey (Redis-compatible) user with ACL-based access control.

## Example Usage

```hcl
resource "hosting_valkey_user" "app" {
  valkey_instance_id = hosting_valkey.cache.id
  privileges         = ["allcommands"]
  key_pattern        = "*"
}
```

## Schema

### Required

- `valkey_instance_id` (String) Parent Valkey instance ID. Changing this forces a new resource.

### Optional

- `password` (String, Sensitive) User password. Generated if not provided.
- `privileges` (List of String) List of Valkey ACL privileges.
- `key_pattern` (String) Key pattern for ACL (e.g. `*`).

### Read-Only

- `id` (String) Valkey user ID.
- `username` (String) Generated username.
- `status` (String) Current status.

## Import

Import uses the format `valkey_instance_id/user_id`:

```shell
terraform import hosting_valkey_user.app vk_abc123/vku_xyz789
```
