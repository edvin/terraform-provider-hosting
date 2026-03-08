---
page_title: "hosting_egress_rule Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a tenant egress firewall rule.
---

# hosting_egress_rule (Resource)

Manages a tenant egress firewall rule. Allows outbound traffic to specific CIDR ranges.

## Example Usage

```hcl
resource "hosting_egress_rule" "api" {
  tenant_id   = var.tenant_id
  cidr        = "203.0.113.0/24"
  description = "Allow access to external API"
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.
- `cidr` (String) CIDR block to allow egress to. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.
- `description` (String) Description of the rule.

### Read-Only

- `id` (String) Egress rule ID.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_egress_rule.api er_abc123
```
