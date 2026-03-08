---
page_title: "hosting_fqdn Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a hostname/FQDN binding.
---

# hosting_fqdn (Resource)

Manages a hostname/FQDN binding. Attach custom domains to webapps with optional SSL.

## Example Usage

```hcl
resource "hosting_fqdn" "main" {
  fqdn        = "myapp.example.com"
  webapp_id   = hosting_webapp.myapp.id
  ssl_enabled = true
}
```

## Schema

### Required

- `fqdn` (String) Fully qualified domain name. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `webapp_id` (String) Webapp ID to attach this FQDN to.
- `ssl_enabled` (Boolean) Enable SSL certificate. Default: `true`.

### Read-Only

- `id` (String) FQDN ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_fqdn.main fqdn_abc123
```
