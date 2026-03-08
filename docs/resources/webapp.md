---
page_title: "hosting_webapp Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a web application.
---

# hosting_webapp (Resource)

Manages a web application. Supports PHP, Node.js, Python, Ruby, Go, Java, and static sites.

## Example Usage

```hcl
resource "hosting_webapp" "myapp" {
  tenant_id       = var.tenant_id
  runtime         = "php"
  runtime_version = "8.4"
  public_folder   = "public"
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.
- `runtime` (String) Runtime type (e.g. php, nodejs, python, ruby, static).
- `runtime_version` (String) Runtime version (e.g. 8.4, 22, 3.13).

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `public_folder` (String) Public folder relative to app root (e.g. `public`). Default: `""`.
- `service_hostname_enabled` (Boolean) Whether the built-in service hostname is enabled. Default: `true`.

### Read-Only

- `id` (String) Webapp ID.
- `env_file_name` (String) Environment file name.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_webapp.myapp wp_abc123
```
