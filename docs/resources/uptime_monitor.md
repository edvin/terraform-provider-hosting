# hosting_uptime_monitor

Manages an uptime monitor that performs periodic HTTP health checks against a URL. When the check fails (non-matching status code or timeout), an incident is automatically created. When it recovers, the incident is auto-resolved.

## Example Usage

```hcl
resource "hosting_uptime_monitor" "main" {
  tenant_id        = var.tenant_id
  url              = "https://myapp.example.com/healthz"
  interval_seconds = 60
  timeout_seconds  = 10
  expected_status  = 200
}

resource "hosting_uptime_monitor" "api" {
  tenant_id        = var.tenant_id
  url              = "https://api.example.com/health"
  interval_seconds = 30
  timeout_seconds  = 5
  expected_status  = 200
}
```

## Argument Reference

- `tenant_id` - (Required, Forces new resource) The tenant ID.
- `url` - (Required) The URL to check.
- `interval_seconds` - (Optional) Check interval in seconds. Min 30, max 3600. Default `300`.
- `timeout_seconds` - (Optional) HTTP request timeout in seconds. Min 1, max 60. Default `10`.
- `expected_status` - (Optional) Expected HTTP status code. Min 100, max 599. Default `200`.
- `enabled` - (Optional) Whether the monitor is enabled. Default `true`.
- `customer_id` - (Optional) Customer ID. Defaults to provider-level `customer_id`.

## Attribute Reference

- `id` - The uptime monitor ID.
- `status` - Current status (always `active`).
- `last_check_at` - Timestamp of the last check.
- `last_status_code` - HTTP status code from the last check.

## Import

```bash
terraform import hosting_uptime_monitor.main <monitor-id>
```
