---
page_title: "hosting_customer_user Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a team member on a customer account.
---

# hosting_customer_user (Resource)

Manages a team member on a customer account with role-based access control. Roles control what actions the user can perform: owners have full access, admins can manage team members and resources, developers can create and modify resources, and viewers have read-only access.

## Example Usage

```hcl
resource "hosting_customer_user" "dev" {
  user_id = "usr_abc123"
  role    = "developer"
}

resource "hosting_customer_user" "readonly" {
  user_id = "usr_xyz789"
  role    = "viewer"
}
```

## Schema

### Required

- `user_id` (String) User ID to add as a team member. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `role` (String) Role to assign. One of: `owner`, `admin`, `developer`, `viewer`. Defaults to `developer`.

### Read-Only

- `email` (String) User's email address.
- `display_name` (String) User's display name.

## Role Permissions

| Role | View | Create/Edit/Delete | Manage Team | Billing/Account |
|------|------|--------------------|-------------|-----------------|
| owner | yes | yes | yes (all roles) | yes |
| admin | yes | yes | yes (dev/viewer only) | no |
| developer | yes | yes | no | no |
| viewer | yes | no | no | no |

## Import

Import uses the format `customer_id/user_id`:

```shell
terraform import hosting_customer_user.dev cust_abc123/usr_xyz789
```
