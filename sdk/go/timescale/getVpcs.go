// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timescale

import (
	"context"
	"reflect"

	"github.com/itoam/pulumi-timescale/sdk/go/timescale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpcs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupVpcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcsResult
	err := ctx.Invoke("timescale:index/getVpcs:getVpcs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getVpcs.
type LookupVpcsResult struct {
	// The ID of this resource.
	Id   string       `pulumi:"id"`
	Vpcs []GetVpcsVpc `pulumi:"vpcs"`
}

func LookupVpcsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupVpcsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupVpcsResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv LookupVpcsResult
		secret, err := ctx.InvokePackageRaw("timescale:index/getVpcs:getVpcs", nil, &rv, "", opts...)
		if err != nil {
			return LookupVpcsResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(LookupVpcsResultOutput)
		if secret {
			return pulumi.ToSecret(output).(LookupVpcsResultOutput), nil
		}
		return output, nil
	}).(LookupVpcsResultOutput)
}

// A collection of values returned by getVpcs.
type LookupVpcsResultOutput struct{ *pulumi.OutputState }

func (LookupVpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcsResult)(nil)).Elem()
}

func (o LookupVpcsResultOutput) ToLookupVpcsResultOutput() LookupVpcsResultOutput {
	return o
}

func (o LookupVpcsResultOutput) ToLookupVpcsResultOutputWithContext(ctx context.Context) LookupVpcsResultOutput {
	return o
}

// The ID of this resource.
func (o LookupVpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcsResultOutput) Vpcs() GetVpcsVpcArrayOutput {
	return o.ApplyT(func(v LookupVpcsResult) []GetVpcsVpc { return v.Vpcs }).(GetVpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcsResultOutput{})
}
