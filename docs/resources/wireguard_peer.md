---
page_title: "hosting_wireguard_peer Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a WireGuard VPN peer.
---

# hosting_wireguard_peer (Resource)

Manages a WireGuard VPN peer. The private key and client configuration are only available on creation.

## Example Usage

```hcl
resource "hosting_wireguard_peer" "dev" {
  tenant_id = var.tenant_id
  name      = "dev-laptop"
}

output "wireguard_config" {
  value     = hosting_wireguard_peer.dev.client_config
  sensitive = true
}
```

## Schema

### Required

- `tenant_id` (String) Tenant ID. Changing this forces a new resource.
- `name` (String) Peer name. Changing this forces a new resource.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`. Changing this forces a new resource.

### Read-Only

- `id` (String) Peer ID.
- `public_key` (String) WireGuard public key.
- `assigned_ip` (String) Assigned VPN IP address.
- `endpoint` (String) WireGuard endpoint address.
- `private_key` (String, Sensitive) WireGuard private key. Only available on creation.
- `client_config` (String, Sensitive) Complete WireGuard client configuration. Only available on creation.
- `status` (String) Current status.

## Import

```shell
terraform import hosting_wireguard_peer.dev wg_abc123
```
