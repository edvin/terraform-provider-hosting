package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/massive-hosting/go-hosting"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &customerUserResource{}
	_ resource.ResourceWithImportState = &customerUserResource{}
)

type customerUserResource struct {
	data *ProviderData
}

type customerUserModel struct {
	CustomerID types.String `tfsdk:"customer_id"`
	UserID     types.String `tfsdk:"user_id"`
	Role       types.String `tfsdk:"role"`
	Email      types.String `tfsdk:"email"`
	DisplayName types.String `tfsdk:"display_name"`
}

type customerUserAPI struct {
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func NewCustomerUser() resource.Resource {
	return &customerUserResource{}
}

func (r *customerUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer_user"
}

func (r *customerUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a team member on a customer account.",
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Optional: true, Computed: true, Description: "Customer ID. Defaults to provider customer_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"user_id": schema.StringAttribute{
				Required: true, Description: "User ID to add as a team member. Changing this forces a new resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"role": schema.StringAttribute{
				Optional: true, Computed: true, Description: "Role: owner, admin, developer, or viewer.",
				Default: stringdefault.StaticString("developer"),
			},
			"email": schema.StringAttribute{
				Computed: true, Description: "User's email address.",
			},
			"display_name": schema.StringAttribute{
				Computed: true, Description: "User's display name.",
			},
		},
	}
}

func (r *customerUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *customerUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan customerUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerID := resolveCustomerID(plan.CustomerID, r.data)
	if customerID == "" {
		resp.Diagnostics.AddError("Missing customer_id", "Set customer_id on the resource or in the provider config.")
		return
	}

	_, err := hosting.Post[map[string]any](ctx, r.data.Client, fmt.Sprintf("/api/v1/customers/%s/team", customerID), map[string]any{
		"user_id": plan.UserID.ValueString(),
		"role":    plan.Role.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Create Customer User Failed", err.Error())
		return
	}

	// Read back the full user details from the team list.
	found, diag := r.findUser(ctx, customerID, plan.UserID.ValueString())
	if diag != "" {
		resp.Diagnostics.AddError("Read Customer User Failed", diag)
		return
	}
	if found == nil {
		resp.Diagnostics.AddError("Read Customer User Failed", "User not found after creation")
		return
	}

	mapCustomerUser(found, &plan, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *customerUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state customerUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerID := state.CustomerID.ValueString()
	if customerID == "" {
		customerID = r.data.CustomerID
	}

	found, diag := r.findUser(ctx, customerID, state.UserID.ValueString())
	if diag != "" {
		resp.Diagnostics.AddError("Read Customer User Failed", diag)
		return
	}
	if found == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	mapCustomerUser(found, &state, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *customerUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan customerUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerID := resolveCustomerID(plan.CustomerID, r.data)
	if customerID == "" {
		resp.Diagnostics.AddError("Missing customer_id", "Set customer_id on the resource or in the provider config.")
		return
	}

	_, err := hosting.Put[map[string]any](ctx, r.data.Client, fmt.Sprintf("/api/v1/customers/%s/team/%s/role", customerID, plan.UserID.ValueString()), map[string]any{
		"role": plan.Role.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Update Customer User Role Failed", err.Error())
		return
	}

	// Read back the full user details.
	found, diag := r.findUser(ctx, customerID, plan.UserID.ValueString())
	if diag != "" {
		resp.Diagnostics.AddError("Read Customer User Failed", diag)
		return
	}
	if found == nil {
		resp.Diagnostics.AddError("Read Customer User Failed", "User not found after update")
		return
	}

	mapCustomerUser(found, &plan, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *customerUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state customerUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerID := state.CustomerID.ValueString()
	if customerID == "" {
		customerID = r.data.CustomerID
	}

	if err := r.data.Client.Delete(ctx, fmt.Sprintf("/api/v1/customers/%s/team/%s", customerID, state.UserID.ValueString())); err != nil {
		resp.Diagnostics.AddError("Delete Customer User Failed", err.Error())
	}
}

func (r *customerUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: customer_id/user_id
	parts := strings.SplitN(req.ID, "/", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: customer_id/user_id")
		return
	}

	customerID := parts[0]
	userID := parts[1]

	found, diag := r.findUser(ctx, customerID, userID)
	if diag != "" {
		resp.Diagnostics.AddError("Import Customer User Failed", diag)
		return
	}
	if found == nil {
		resp.Diagnostics.AddError("Import Customer User Failed", "User not found")
		return
	}

	var state customerUserModel
	mapCustomerUser(found, &state, customerID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *customerUserResource) findUser(ctx context.Context, customerID, userID string) (*customerUserAPI, string) {
	users, err := hosting.List[customerUserAPI](ctx, r.data.Client, fmt.Sprintf("/api/v1/customers/%s/team", customerID))
	if err != nil {
		return nil, err.Error()
	}
	for i := range users {
		if users[i].UserID == userID {
			return &users[i], ""
		}
	}
	return nil, ""
}

func mapCustomerUser(api *customerUserAPI, state *customerUserModel, customerID string) {
	state.UserID = types.StringValue(api.UserID)
	state.Role = types.StringValue(api.Role)
	state.Email = types.StringValue(api.Email)
	state.DisplayName = types.StringValue(api.DisplayName)
	if customerID != "" {
		state.CustomerID = types.StringValue(customerID)
	}
}
