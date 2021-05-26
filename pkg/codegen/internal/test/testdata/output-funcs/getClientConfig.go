// *** WARNING: this file was generated by tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codegentest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this function to access the current configuration of the native Azure provider.
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	var rv GetClientConfigResult
	err := ctx.Invoke("azure-native:codegentest:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// Configuration values returned by getClientConfig.
type GetClientConfigResult struct {
	// Azure Client ID (Application Object ID).
	ClientId string `pulumi:"clientId"`
	// Azure Object ID of the current user or service principal.
	ObjectId string `pulumi:"objectId"`
	// Azure Subscription ID
	SubscriptionId string `pulumi:"subscriptionId"`
	// Azure Tenant ID
	TenantId string `pulumi:"tenantId"`
}


func GetClientConfigApply(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetClientConfigResultOutput {
	return pulumi.Any(opts).ApplyT(func (v interface{}) (GetClientConfigResult, error) {
	r, err := GetClientConfig(ctx, opts...)
	return *r, err

}).(GetClientConfigResultOutput)}
// Configuration values returned by getClientConfig.
type GetClientConfigResultOutput struct { *pulumi.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutput() GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutputWithContext(ctx context.Context) GetClientConfigResultOutput {
	return o
}

// Azure Client ID (Application Object ID).
func (o GetClientConfigResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func (v GetClientConfigResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Azure Object ID of the current user or service principal.
func (o GetClientConfigResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func (v GetClientConfigResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// Azure Subscription ID
func (o GetClientConfigResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func (v GetClientConfigResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Azure Tenant ID
func (o GetClientConfigResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func (v GetClientConfigResult) string { return v.TenantId }).(pulumi.StringOutput)
}


func init() {
	pulumi.RegisterOutputType(GetClientConfigResultOutput{})
}