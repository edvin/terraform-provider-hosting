package resources

import (
	"context"
	"fmt"

	"github.com/massive-hosting/go-hosting"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &uptimeMonitorResource{}
	_ resource.ResourceWithImportState = &uptimeMonitorResource{}
)

type uptimeMonitorResource struct {
	data *ProviderData
}

type uptimeMonitorModel struct {
	ID              types.String `tfsdk:"id"`
	CustomerID      types.String `tfsdk:"customer_id"`
	TenantID        types.String `tfsdk:"tenant_id"`
	URL             types.String `tfsdk:"url"`
	IntervalSeconds types.Int64  `tfsdk:"interval_seconds"`
	TimeoutSeconds  types.Int64  `tfsdk:"timeout_seconds"`
	ExpectedStatus  types.Int64  `tfsdk:"expected_status"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	Status          types.String `tfsdk:"status"`
	LastCheckAt     types.String `tfsdk:"last_check_at"`
	LastStatusCode  types.Int64  `tfsdk:"last_status_code"`
}

type uptimeMonitorAPI struct {
	ID              string  `json:"id"`
	TenantID        string  `json:"tenant_id"`
	URL             string  `json:"url"`
	IntervalSeconds int     `json:"interval_seconds"`
	TimeoutSeconds  int     `json:"timeout_seconds"`
	ExpectedStatus  int     `json:"expected_status"`
	Enabled         bool    `json:"enabled"`
	Status          string  `json:"status"`
	LastCheckAt     *string `json:"last_check_at"`
	LastStatusCode  *int    `json:"last_status_code"`
}

func NewUptimeMonitor() resource.Resource {
	return &uptimeMonitorResource{}
}

func (r *uptimeMonitorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_uptime_monitor"
}

func (r *uptimeMonitorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an uptime monitor that performs periodic HTTP health checks.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, Description: "Uptime monitor ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_id": schema.StringAttribute{
				Optional: true, Computed: true, Description: "Customer ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"tenant_id": schema.StringAttribute{
				Required: true, Description: "Tenant ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"url": schema.StringAttribute{
				Required: true, Description: "URL to check.",
			},
			"interval_seconds": schema.Int64Attribute{
				Optional: true, Computed: true, Description: "Check interval in seconds (30-3600).",
				Default: int64default.StaticInt64(300),
			},
			"timeout_seconds": schema.Int64Attribute{
				Optional: true, Computed: true, Description: "HTTP request timeout in seconds (1-60).",
				Default: int64default.StaticInt64(10),
			},
			"expected_status": schema.Int64Attribute{
				Optional: true, Computed: true, Description: "Expected HTTP status code (100-599).",
				Default: int64default.StaticInt64(200),
			},
			"enabled": schema.BoolAttribute{
				Optional: true, Computed: true, Description: "Whether the monitor is enabled.",
				Default: booldefault.StaticBool(true),
			},
			"status": schema.StringAttribute{
				Computed: true, Description: "Current status.",
			},
			"last_check_at": schema.StringAttribute{
				Computed: true, Description: "Timestamp of last check.",
			},
			"last_status_code": schema.Int64Attribute{
				Computed: true, Description: "Status code from last check.",
			},
		},
	}
}

func (r *uptimeMonitorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(*ProviderData)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", fmt.Sprintf("Expected *ProviderData, got %T", req.ProviderData))
		return
	}
	r.data = data
}

func (r *uptimeMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan uptimeMonitorModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerID := resolveCustomerID(plan.CustomerID, r.data)
	if customerID == "" {
		resp.Diagnostics.AddError("Missing customer_id", "Set customer_id on the resource or in the provider config.")
		return
	}

	body := map[string]any{
		"tenant_id":        plan.TenantID.ValueString(),
		"url":              plan.URL.ValueString(),
		"interval_seconds": plan.IntervalSeconds.ValueInt64(),
		"timeout_seconds":  plan.TimeoutSeconds.ValueInt64(),
		"expected_status":  plan.ExpectedStatus.ValueInt64(),
	}

	result, err := hosting.Post[uptimeMonitorAPI](ctx, r.data.Client, fmt.Sprintf("/api/v1/customers/%s/uptime-monitors", customerID), body)
	if err != nil {
		resp.Diagnostics.AddError("Create Uptime Monitor Failed", err.Error())
		return
	}

	mapUptimeMonitor(result, &plan, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *uptimeMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state uptimeMonitorModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := hosting.Get[uptimeMonitorAPI](ctx, r.data.Client, fmt.Sprintf("/api/v1/uptime-monitors/%s", state.ID.ValueString()))
	if err != nil {
		resp.State.RemoveResource(ctx)
		return
	}

	customerID := state.CustomerID.ValueString()
	if customerID == "" {
		customerID = r.data.CustomerID
	}
	mapUptimeMonitor(result, &state, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *uptimeMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan uptimeMonitorModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := map[string]any{
		"url":              plan.URL.ValueString(),
		"interval_seconds": plan.IntervalSeconds.ValueInt64(),
		"timeout_seconds":  plan.TimeoutSeconds.ValueInt64(),
		"expected_status":  plan.ExpectedStatus.ValueInt64(),
		"enabled":          plan.Enabled.ValueBool(),
	}

	result, err := hosting.Patch[uptimeMonitorAPI](ctx, r.data.Client, fmt.Sprintf("/api/v1/uptime-monitors/%s", plan.ID.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Update Uptime Monitor Failed", err.Error())
		return
	}

	customerID := plan.CustomerID.ValueString()
	if customerID == "" {
		customerID = r.data.CustomerID
	}
	mapUptimeMonitor(result, &plan, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *uptimeMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state uptimeMonitorModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.data.Client.Delete(ctx, "/api/v1/uptime-monitors/"+state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete Uptime Monitor Failed", err.Error())
	}
}

func (r *uptimeMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	result, err := hosting.Get[uptimeMonitorAPI](ctx, r.data.Client, fmt.Sprintf("/api/v1/uptime-monitors/%s", req.ID))
	if err != nil {
		resp.Diagnostics.AddError("Import Uptime Monitor Failed", err.Error())
		return
	}

	customerID := r.data.CustomerID
	var state uptimeMonitorModel
	mapUptimeMonitor(result, &state, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func mapUptimeMonitor(api *uptimeMonitorAPI, state *uptimeMonitorModel, customerID string) {
	state.ID = types.StringValue(api.ID)
	state.TenantID = types.StringValue(api.TenantID)
	state.URL = types.StringValue(api.URL)
	state.IntervalSeconds = types.Int64Value(int64(api.IntervalSeconds))
	state.TimeoutSeconds = types.Int64Value(int64(api.TimeoutSeconds))
	state.ExpectedStatus = types.Int64Value(int64(api.ExpectedStatus))
	state.Enabled = types.BoolValue(api.Enabled)
	state.Status = types.StringValue(api.Status)
	if api.LastCheckAt != nil {
		state.LastCheckAt = types.StringValue(*api.LastCheckAt)
	} else {
		state.LastCheckAt = types.StringNull()
	}
	if api.LastStatusCode != nil {
		state.LastStatusCode = types.Int64Value(int64(*api.LastStatusCode))
	} else {
		state.LastStatusCode = types.Int64Null()
	}
	if customerID != "" {
		state.CustomerID = types.StringValue(customerID)
	}
}
