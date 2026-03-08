---
page_title: "hosting_valkey Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a Valkey (Redis) instance.
---

# hosting_valkey (Resource)

Manages a Valkey (Redis-compatible) instance.

## Example Usage

```hcl
resource "hosting_valkey" "cache" {
  tenant_id     = var.tenant_id
  max_memory_mb = 128
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `max_memory_mb` (Number) Maximum memory in MB. Default: `64`.

### Read-Only

- `id` (String) Valkey instance ID.
- `port` (Number) Assigned port number.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_valkey.cache vk_abc123
```
