// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: Dataset
//
// Creates a dataset.
//
// > **Note** If this dataset already exists, creating this resource is a no-op.
//
// > **Note** Destroying or replacing this resource will not delete the created dataset. It's not possible to delete a dataset using the API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := honeycomb.NewDataset(ctx, "myDataset", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Dataset struct {
	pulumi.CustomResourceState

	// The name of the dataset.
	Name pulumi.StringOutput `pulumi:"name"`
	// The slug of the dataset.
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		args = &DatasetArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Dataset
	err := ctx.RegisterResource("honeycomb:index/dataset:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("honeycomb:index/dataset:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
	// The name of the dataset.
	Name *string `pulumi:"name"`
	// The slug of the dataset.
	Slug *string `pulumi:"slug"`
}

type DatasetState struct {
	// The name of the dataset.
	Name pulumi.StringPtrInput
	// The slug of the dataset.
	Slug pulumi.StringPtrInput
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// The name of the dataset.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// The name of the dataset.
	Name pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

// DatasetArrayInput is an input type that accepts DatasetArray and DatasetArrayOutput values.
// You can construct a concrete instance of `DatasetArrayInput` via:
//
//          DatasetArray{ DatasetArgs{...} }
type DatasetArrayInput interface {
	pulumi.Input

	ToDatasetArrayOutput() DatasetArrayOutput
	ToDatasetArrayOutputWithContext(context.Context) DatasetArrayOutput
}

type DatasetArray []DatasetInput

func (DatasetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataset)(nil)).Elem()
}

func (i DatasetArray) ToDatasetArrayOutput() DatasetArrayOutput {
	return i.ToDatasetArrayOutputWithContext(context.Background())
}

func (i DatasetArray) ToDatasetArrayOutputWithContext(ctx context.Context) DatasetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetArrayOutput)
}

// DatasetMapInput is an input type that accepts DatasetMap and DatasetMapOutput values.
// You can construct a concrete instance of `DatasetMapInput` via:
//
//          DatasetMap{ "key": DatasetArgs{...} }
type DatasetMapInput interface {
	pulumi.Input

	ToDatasetMapOutput() DatasetMapOutput
	ToDatasetMapOutputWithContext(context.Context) DatasetMapOutput
}

type DatasetMap map[string]DatasetInput

func (DatasetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataset)(nil)).Elem()
}

func (i DatasetMap) ToDatasetMapOutput() DatasetMapOutput {
	return i.ToDatasetMapOutputWithContext(context.Background())
}

func (i DatasetMap) ToDatasetMapOutputWithContext(ctx context.Context) DatasetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetMapOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

// The name of the dataset.
func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The slug of the dataset.
func (o DatasetOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type DatasetArrayOutput struct{ *pulumi.OutputState }

func (DatasetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataset)(nil)).Elem()
}

func (o DatasetArrayOutput) ToDatasetArrayOutput() DatasetArrayOutput {
	return o
}

func (o DatasetArrayOutput) ToDatasetArrayOutputWithContext(ctx context.Context) DatasetArrayOutput {
	return o
}

func (o DatasetArrayOutput) Index(i pulumi.IntInput) DatasetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dataset {
		return vs[0].([]*Dataset)[vs[1].(int)]
	}).(DatasetOutput)
}

type DatasetMapOutput struct{ *pulumi.OutputState }

func (DatasetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataset)(nil)).Elem()
}

func (o DatasetMapOutput) ToDatasetMapOutput() DatasetMapOutput {
	return o
}

func (o DatasetMapOutput) ToDatasetMapOutputWithContext(ctx context.Context) DatasetMapOutput {
	return o
}

func (o DatasetMapOutput) MapIndex(k pulumi.StringInput) DatasetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dataset {
		return vs[0].(map[string]*Dataset)[vs[1].(string)]
	}).(DatasetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetInput)(nil)).Elem(), &Dataset{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetArrayInput)(nil)).Elem(), DatasetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetMapInput)(nil)).Elem(), DatasetMap{})
	pulumi.RegisterOutputType(DatasetOutput{})
	pulumi.RegisterOutputType(DatasetArrayOutput{})
	pulumi.RegisterOutputType(DatasetMapOutput{})
}