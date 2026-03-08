---
page_title: "hosting_s3_access_key Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an S3 access key for a bucket.
---

# hosting_s3_access_key (Resource)

Manages an S3 access key for a bucket. The secret access key is only available on creation.

## Example Usage

```hcl
resource "hosting_s3_access_key" "app" {
  s3_bucket_id = hosting_s3_bucket.assets.id
  permissions  = ["read", "write"]
}

output "s3_credentials" {
  value = {
    access_key = hosting_s3_access_key.app.access_key_id
    secret_key = hosting_s3_access_key.app.secret_access_key
  }
  sensitive = true
}
```

## Schema

### Required

- `s3_bucket_id` (String) Parent S3 bucket ID. Changing this forces a new resource.

### Optional

- `permissions` (List of String) List of permissions (e.g. `read`, `write`).

### Read-Only

- `id` (String) Access key resource ID.
- `access_key_id` (String) AWS-style access key ID.
- `secret_access_key` (String, Sensitive) Secret access key. Only available on creation.
- `status` (String) Current status.

## Import

Import uses the format `s3_bucket_id/access_key_id`:

```shell
terraform import hosting_s3_access_key.app s3_abc123/s3k_xyz789
```
