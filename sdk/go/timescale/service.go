// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timescale

import (
	"context"
	"reflect"

	"github.com/itoam/pulumi-timescale/sdk/go/timescale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/itoam/pulumi-timescale/sdk/go/timescale"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := timescale.NewService(ctx, "test", nil)
//			if err != nil {
//				return err
//			}
//			// Read replica
//			_, err = timescale.NewService(ctx, "read_replica", &timescale.ServiceArgs{
//				ReadReplicaSource: test.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Service struct {
	pulumi.CustomResourceState

	// Set connection pooler status for this service.
	ConnectionPoolerEnabled pulumi.BoolOutput `pulumi:"connectionPoolerEnabled"`
	// Enable HA Replica
	EnableHaReplica pulumi.BoolOutput `pulumi:"enableHaReplica"`
	// Set environment tag for this service.
	EnvironmentTag pulumi.StringOutput `pulumi:"environmentTag"`
	// The hostname for this service
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Memory GB
	MemoryGb pulumi.IntOutput `pulumi:"memoryGb"`
	// Milli CPU
	MilliCpu pulumi.IntOutput `pulumi:"milliCpu"`
	// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Postgres password for this service
	Password pulumi.StringOutput `pulumi:"password"`
	// Paused status of the service.
	Paused pulumi.BoolOutput `pulumi:"paused"`
	// Hostname of the pooler of this service.
	PoolerHostname pulumi.StringOutput `pulumi:"poolerHostname"`
	// Port of the pooler of this service.
	PoolerPort pulumi.IntOutput `pulumi:"poolerPort"`
	// The port for this service
	Port pulumi.IntOutput `pulumi:"port"`
	// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
	ReadReplicaSource pulumi.StringPtrOutput `pulumi:"readReplicaSource"`
	// The region for this service.
	RegionCode pulumi.StringOutput `pulumi:"regionCode"`
	// Hostname of the HA-Replica of this service.
	ReplicaHostname pulumi.StringOutput `pulumi:"replicaHostname"`
	// Port of the HA-Replica of this service.
	ReplicaPort pulumi.IntOutput `pulumi:"replicaPort"`
	// Deprecated: Storage GB
	//
	// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
	StorageGb pulumi.IntPtrOutput      `pulumi:"storageGb"`
	Timeouts  ServiceTimeoutsPtrOutput `pulumi:"timeouts"`
	// The Postgres user for this service
	Username pulumi.StringOutput `pulumi:"username"`
	// The VpcID this service is tied to.
	VpcId pulumi.IntPtrOutput `pulumi:"vpcId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("timescale:index/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("timescale:index/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Set connection pooler status for this service.
	ConnectionPoolerEnabled *bool `pulumi:"connectionPoolerEnabled"`
	// Enable HA Replica
	EnableHaReplica *bool `pulumi:"enableHaReplica"`
	// Set environment tag for this service.
	EnvironmentTag *string `pulumi:"environmentTag"`
	// The hostname for this service
	Hostname *string `pulumi:"hostname"`
	// Memory GB
	MemoryGb *int `pulumi:"memoryGb"`
	// Milli CPU
	MilliCpu *int `pulumi:"milliCpu"`
	// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
	Name *string `pulumi:"name"`
	// The Postgres password for this service
	Password *string `pulumi:"password"`
	// Paused status of the service.
	Paused *bool `pulumi:"paused"`
	// Hostname of the pooler of this service.
	PoolerHostname *string `pulumi:"poolerHostname"`
	// Port of the pooler of this service.
	PoolerPort *int `pulumi:"poolerPort"`
	// The port for this service
	Port *int `pulumi:"port"`
	// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
	ReadReplicaSource *string `pulumi:"readReplicaSource"`
	// The region for this service.
	RegionCode *string `pulumi:"regionCode"`
	// Hostname of the HA-Replica of this service.
	ReplicaHostname *string `pulumi:"replicaHostname"`
	// Port of the HA-Replica of this service.
	ReplicaPort *int `pulumi:"replicaPort"`
	// Deprecated: Storage GB
	//
	// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
	StorageGb *int             `pulumi:"storageGb"`
	Timeouts  *ServiceTimeouts `pulumi:"timeouts"`
	// The Postgres user for this service
	Username *string `pulumi:"username"`
	// The VpcID this service is tied to.
	VpcId *int `pulumi:"vpcId"`
}

type ServiceState struct {
	// Set connection pooler status for this service.
	ConnectionPoolerEnabled pulumi.BoolPtrInput
	// Enable HA Replica
	EnableHaReplica pulumi.BoolPtrInput
	// Set environment tag for this service.
	EnvironmentTag pulumi.StringPtrInput
	// The hostname for this service
	Hostname pulumi.StringPtrInput
	// Memory GB
	MemoryGb pulumi.IntPtrInput
	// Milli CPU
	MilliCpu pulumi.IntPtrInput
	// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
	Name pulumi.StringPtrInput
	// The Postgres password for this service
	Password pulumi.StringPtrInput
	// Paused status of the service.
	Paused pulumi.BoolPtrInput
	// Hostname of the pooler of this service.
	PoolerHostname pulumi.StringPtrInput
	// Port of the pooler of this service.
	PoolerPort pulumi.IntPtrInput
	// The port for this service
	Port pulumi.IntPtrInput
	// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
	ReadReplicaSource pulumi.StringPtrInput
	// The region for this service.
	RegionCode pulumi.StringPtrInput
	// Hostname of the HA-Replica of this service.
	ReplicaHostname pulumi.StringPtrInput
	// Port of the HA-Replica of this service.
	ReplicaPort pulumi.IntPtrInput
	// Deprecated: Storage GB
	//
	// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
	StorageGb pulumi.IntPtrInput
	Timeouts  ServiceTimeoutsPtrInput
	// The Postgres user for this service
	Username pulumi.StringPtrInput
	// The VpcID this service is tied to.
	VpcId pulumi.IntPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Set connection pooler status for this service.
	ConnectionPoolerEnabled *bool `pulumi:"connectionPoolerEnabled"`
	// Enable HA Replica
	EnableHaReplica *bool `pulumi:"enableHaReplica"`
	// Set environment tag for this service.
	EnvironmentTag *string `pulumi:"environmentTag"`
	// Memory GB
	MemoryGb *int `pulumi:"memoryGb"`
	// Milli CPU
	MilliCpu *int `pulumi:"milliCpu"`
	// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
	Name *string `pulumi:"name"`
	// The Postgres password for this service
	Password *string `pulumi:"password"`
	// Paused status of the service.
	Paused *bool `pulumi:"paused"`
	// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
	ReadReplicaSource *string `pulumi:"readReplicaSource"`
	// The region for this service.
	RegionCode *string `pulumi:"regionCode"`
	// Deprecated: Storage GB
	//
	// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
	StorageGb *int             `pulumi:"storageGb"`
	Timeouts  *ServiceTimeouts `pulumi:"timeouts"`
	// The VpcID this service is tied to.
	VpcId *int `pulumi:"vpcId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Set connection pooler status for this service.
	ConnectionPoolerEnabled pulumi.BoolPtrInput
	// Enable HA Replica
	EnableHaReplica pulumi.BoolPtrInput
	// Set environment tag for this service.
	EnvironmentTag pulumi.StringPtrInput
	// Memory GB
	MemoryGb pulumi.IntPtrInput
	// Milli CPU
	MilliCpu pulumi.IntPtrInput
	// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
	Name pulumi.StringPtrInput
	// The Postgres password for this service
	Password pulumi.StringPtrInput
	// Paused status of the service.
	Paused pulumi.BoolPtrInput
	// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
	ReadReplicaSource pulumi.StringPtrInput
	// The region for this service.
	RegionCode pulumi.StringPtrInput
	// Deprecated: Storage GB
	//
	// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
	StorageGb pulumi.IntPtrInput
	Timeouts  ServiceTimeoutsPtrInput
	// The VpcID this service is tied to.
	VpcId pulumi.IntPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Set connection pooler status for this service.
func (o ServiceOutput) ConnectionPoolerEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.ConnectionPoolerEnabled }).(pulumi.BoolOutput)
}

// Enable HA Replica
func (o ServiceOutput) EnableHaReplica() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.EnableHaReplica }).(pulumi.BoolOutput)
}

// Set environment tag for this service.
func (o ServiceOutput) EnvironmentTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.EnvironmentTag }).(pulumi.StringOutput)
}

// The hostname for this service
func (o ServiceOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Memory GB
func (o ServiceOutput) MemoryGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.MemoryGb }).(pulumi.IntOutput)
}

// Milli CPU
func (o ServiceOutput) MilliCpu() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.MilliCpu }).(pulumi.IntOutput)
}

// Service Name is the configurable name assigned to this resource. If none is provided, a default will be generated by the provider.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Postgres password for this service
func (o ServiceOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Paused status of the service.
func (o ServiceOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.Paused }).(pulumi.BoolOutput)
}

// Hostname of the pooler of this service.
func (o ServiceOutput) PoolerHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.PoolerHostname }).(pulumi.StringOutput)
}

// Port of the pooler of this service.
func (o ServiceOutput) PoolerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.PoolerPort }).(pulumi.IntOutput)
}

// The port for this service
func (o ServiceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// If set, this database will be a read replica of the provided source database. The region must be the same as the source, or if omitted will be handled by the provider
func (o ServiceOutput) ReadReplicaSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.ReadReplicaSource }).(pulumi.StringPtrOutput)
}

// The region for this service.
func (o ServiceOutput) RegionCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.RegionCode }).(pulumi.StringOutput)
}

// Hostname of the HA-Replica of this service.
func (o ServiceOutput) ReplicaHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ReplicaHostname }).(pulumi.StringOutput)
}

// Port of the HA-Replica of this service.
func (o ServiceOutput) ReplicaPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.ReplicaPort }).(pulumi.IntOutput)
}

// Deprecated: Storage GB
//
// Deprecated: This field is ignored. With the new usage-based storage Timescale automatically allocates the disk space needed by your service and you only pay for the disk space you use.
func (o ServiceOutput) StorageGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.StorageGb }).(pulumi.IntPtrOutput)
}

func (o ServiceOutput) Timeouts() ServiceTimeoutsPtrOutput {
	return o.ApplyT(func(v *Service) ServiceTimeoutsPtrOutput { return v.Timeouts }).(ServiceTimeoutsPtrOutput)
}

// The Postgres user for this service
func (o ServiceOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// The VpcID this service is tied to.
func (o ServiceOutput) VpcId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.VpcId }).(pulumi.IntPtrOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
