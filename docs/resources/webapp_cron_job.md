---
page_title: "hosting_webapp_cron_job Resource - terraform-provider-hosting"
subcategory: ""
description: |-
  Manages a cron job for a webapp.
---

# hosting_webapp_cron_job (Resource)

Manages a scheduled cron job for a webapp.

## Example Usage

```hcl
resource "hosting_webapp_cron_job" "cleanup" {
  webapp_id       = hosting_webapp.myapp.id
  schedule        = "0 2 * * *"
  command         = "php artisan schedule:run"
  timeout_seconds = 300
  max_memory_mb   = 128
}
```

## Schema

### Required

- `webapp_id` (String) Webapp ID. Changing this forces a new resource.
- `schedule` (String) Cron schedule expression (e.g. `*/5 * * * *`).
- `command` (String) Command to run.

### Optional

- `customer_id` (String) Customer ID. Defaults to provider `customer_id`.
- `working_directory` (String) Working directory for the command.
- `timeout_seconds` (Number) Maximum execution time in seconds (0 = no limit). Default: `0`.
- `max_memory_mb` (Number) Memory limit in MB (0 = unlimited). Default: `0`.

### Read-Only

- `id` (String) Cron job ID.
- `enabled` (Boolean) Whether the cron job is enabled.
- `status` (String) Current status.
- `status_message` (String) Status message (set on failure).

## Import

```shell
terraform import hosting_webapp_cron_job.cleanup cj_abc123
```
