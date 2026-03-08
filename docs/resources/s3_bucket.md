---
page_title: "hosting_s3_bucket Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages an S3 storage bucket.
---

# hosting_s3_bucket (Resource)

Manages an S3-compatible storage bucket.

## Example Usage

```hcl
resource "hosting_s3_bucket" "assets" {
  tenant_id   = var.tenant_id
  public      = true
  quota_bytes = 10737418240  # 10 GB
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `public` (Boolean) Whether the bucket is publicly accessible. Default: `false`.
- `quota_bytes` (Number) Storage quota in bytes (0 = unlimited). Default: `0`.

### Read-Only

- `id` (String) Bucket ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_s3_bucket.assets s3_abc123
```
