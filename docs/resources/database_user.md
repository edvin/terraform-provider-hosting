---
page_title: "hosting_database_user Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a MySQL database user.
---

# hosting_database_user (Resource)

Manages a MySQL database user.

## Example Usage

```hcl
resource "hosting_database_user" "myuser" {
  database_id = hosting_database.mydb.id
  privileges  = ["ALL"]
}
```

## Schema

### Required

- `database_id` (String) Parent database ID. Changing this forces a new resource.

### Optional

- `password` (String, Sensitive) User password. Generated if not provided.
- `privileges` (List of String) List of MySQL privileges (e.g. SELECT, INSERT, UPDATE, DELETE, ALL).

### Read-Only

- `id` (String) Database user ID.
- `username` (String) Generated username.
- `status` (String) Current status.

## Import

Import uses the format `database_id/user_id`:

```shell
terraform import hosting_database_user.myuser db_abc123/dbu_xyz789
```
