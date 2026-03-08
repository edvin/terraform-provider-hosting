---
page_title: "hosting_container_env_vars Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages environment variables for a container.
---

# hosting_container_env_vars (Resource)

Manages environment variables for a container. This resource manages ALL env vars for the container -- any vars not included will be removed.

~> **Note:** This resource replaces all environment variables on the container. Variables not specified in `vars` or `secret_vars` will be deleted.

## Example Usage

```hcl
resource "hosting_container_env_vars" "redis" {
  container_id = hosting_container.redis.id

  vars = {
    REDIS_MAXMEMORY = "128mb"
  }

  secret_vars = {
    REDIS_PASSWORD = var.redis_password
  }
}
```

## Schema

### Required

- `container_id` (String) Container ID. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`.
- `vars` (Map of String) Non-secret environment variables (name -> value).
- `secret_vars` (Map of String, Sensitive) Secret environment variables (name -> value). Values are encrypted server-side and cannot be read back.

### Read-Only

- `id` (String) Resource ID (same as container_id).
