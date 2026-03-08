---
page_title: "hosting_database Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a MySQL database.
---

# hosting_database (Resource)

Manages a MySQL database.

## Example Usage

```hcl
resource "hosting_database" "mydb" {
  tenant_id = var.tenant_id
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.

### Read-Only

- `id` (String) Database ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_database.mydb db_abc123
```
