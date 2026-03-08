---
page_title: "hosting_container Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an OCI container.
---

# hosting_container (Resource)

Manages an OCI container.

## Example Usage

```hcl
resource "hosting_container" "redis" {
  tenant_id      = var.tenant_id
  name           = "redis"
  image          = "redis:7-alpine"
  max_memory_mb  = 128
  restart_policy = "always"
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.
- `name` (String) Container name.
- `image` (String) Container image (e.g. `nginx:latest`).

### Optional

- `command` (String) Override command.
- `restart_policy` (String) Restart policy (`always`, `on-failure`, `no`). Default: `"always"`.
- `max_memory_mb` (Number) Memory limit in MB. Default: `256`.
- `max_cpu_cores` (Number) CPU cores limit. Default: `1.0`.
- `enabled` (Boolean) Whether the container is enabled. Default: `true`.

### Read-Only

- `id` (String) Container ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_container.redis ct_abc123
```
