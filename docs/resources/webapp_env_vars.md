---
page_title: "hosting_webapp_env_vars Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages environment variables for a webapp.
---

# hosting_webapp_env_vars (Resource)

Manages environment variables for a webapp. This resource manages ALL env vars for the webapp -- any vars not included will be removed.

~> **Note:** This resource replaces all environment variables on the webapp. Variables not specified in `vars` or `secret_vars` will be deleted.

## Example Usage

```hcl
resource "hosting_webapp_env_vars" "myapp" {
  webapp_id = hosting_webapp.myapp.id

  vars = {
    APP_ENV   = "production"
    APP_DEBUG = "false"
    APP_URL   = "https://myapp.example.com"
  }

  secret_vars = {
    APP_KEY    = var.app_key
    DB_PASSWORD = var.db_password
  }
}
```

## Schema

### Required

- `webapp_id` (String) Webapp ID. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`.
- `vars` (Map of String) Non-secret environment variables (name -> value).
- `secret_vars` (Map of String, Sensitive) Secret environment variables (name -> value). Values are encrypted server-side and cannot be read back.

### Read-Only

- `id` (String) Resource ID (same as webapp_id).
