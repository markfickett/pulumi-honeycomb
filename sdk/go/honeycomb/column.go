// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		dataset := cfg.Require("dataset")
// 		_, err := honeycomb.NewColumn(ctx, "durationMs", &honeycomb.ColumnArgs{
// 			KeyName:     pulumi.String("duration_ms_log10"),
// 			Type:        pulumi.String("float"),
// 			Description: pulumi.String("Duration of the trace"),
// 			Dataset:     pulumi.String(dataset),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Columns can be imported using a combination of the dataset name and their key name, e.g.
//
// ```sh
//  $ pulumi import honeycomb:index/column:Column my_column my-dataset/duration_ms
// ```
type Column struct {
	pulumi.CustomResourceState

	// The dataset this column is added to.
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this column should be hidden in the query builder and sample data. Defaults to false.
	Hidden pulumi.BoolPtrOutput `pulumi:"hidden"`
	// The name of the column. Must be unique per dataset.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewColumn registers a new resource with the given unique name, arguments, and options.
func NewColumn(ctx *pulumi.Context,
	name string, args *ColumnArgs, opts ...pulumi.ResourceOption) (*Column, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Column
	err := ctx.RegisterResource("honeycomb:index/column:Column", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetColumn gets an existing Column resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetColumn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ColumnState, opts ...pulumi.ResourceOption) (*Column, error) {
	var resource Column
	err := ctx.ReadResource("honeycomb:index/column:Column", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Column resources.
type columnState struct {
	// The dataset this column is added to.
	Dataset *string `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description *string `pulumi:"description"`
	// Whether this column should be hidden in the query builder and sample data. Defaults to false.
	Hidden *bool `pulumi:"hidden"`
	// The name of the column. Must be unique per dataset.
	KeyName *string `pulumi:"keyName"`
	// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
	Type *string `pulumi:"type"`
}

type ColumnState struct {
	// The dataset this column is added to.
	Dataset pulumi.StringPtrInput
	// A description that is shown in the UI.
	Description pulumi.StringPtrInput
	// Whether this column should be hidden in the query builder and sample data. Defaults to false.
	Hidden pulumi.BoolPtrInput
	// The name of the column. Must be unique per dataset.
	KeyName pulumi.StringPtrInput
	// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
	Type pulumi.StringPtrInput
}

func (ColumnState) ElementType() reflect.Type {
	return reflect.TypeOf((*columnState)(nil)).Elem()
}

type columnArgs struct {
	// The dataset this column is added to.
	Dataset string `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description *string `pulumi:"description"`
	// Whether this column should be hidden in the query builder and sample data. Defaults to false.
	Hidden *bool `pulumi:"hidden"`
	// The name of the column. Must be unique per dataset.
	KeyName string `pulumi:"keyName"`
	// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Column resource.
type ColumnArgs struct {
	// The dataset this column is added to.
	Dataset pulumi.StringInput
	// A description that is shown in the UI.
	Description pulumi.StringPtrInput
	// Whether this column should be hidden in the query builder and sample data. Defaults to false.
	Hidden pulumi.BoolPtrInput
	// The name of the column. Must be unique per dataset.
	KeyName pulumi.StringInput
	// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
	Type pulumi.StringPtrInput
}

func (ColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*columnArgs)(nil)).Elem()
}

type ColumnInput interface {
	pulumi.Input

	ToColumnOutput() ColumnOutput
	ToColumnOutputWithContext(ctx context.Context) ColumnOutput
}

func (*Column) ElementType() reflect.Type {
	return reflect.TypeOf((**Column)(nil)).Elem()
}

func (i *Column) ToColumnOutput() ColumnOutput {
	return i.ToColumnOutputWithContext(context.Background())
}

func (i *Column) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnOutput)
}

// ColumnArrayInput is an input type that accepts ColumnArray and ColumnArrayOutput values.
// You can construct a concrete instance of `ColumnArrayInput` via:
//
//          ColumnArray{ ColumnArgs{...} }
type ColumnArrayInput interface {
	pulumi.Input

	ToColumnArrayOutput() ColumnArrayOutput
	ToColumnArrayOutputWithContext(context.Context) ColumnArrayOutput
}

type ColumnArray []ColumnInput

func (ColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Column)(nil)).Elem()
}

func (i ColumnArray) ToColumnArrayOutput() ColumnArrayOutput {
	return i.ToColumnArrayOutputWithContext(context.Background())
}

func (i ColumnArray) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnArrayOutput)
}

// ColumnMapInput is an input type that accepts ColumnMap and ColumnMapOutput values.
// You can construct a concrete instance of `ColumnMapInput` via:
//
//          ColumnMap{ "key": ColumnArgs{...} }
type ColumnMapInput interface {
	pulumi.Input

	ToColumnMapOutput() ColumnMapOutput
	ToColumnMapOutputWithContext(context.Context) ColumnMapOutput
}

type ColumnMap map[string]ColumnInput

func (ColumnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Column)(nil)).Elem()
}

func (i ColumnMap) ToColumnMapOutput() ColumnMapOutput {
	return i.ToColumnMapOutputWithContext(context.Background())
}

func (i ColumnMap) ToColumnMapOutputWithContext(ctx context.Context) ColumnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnMapOutput)
}

type ColumnOutput struct{ *pulumi.OutputState }

func (ColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Column)(nil)).Elem()
}

func (o ColumnOutput) ToColumnOutput() ColumnOutput {
	return o
}

func (o ColumnOutput) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return o
}

// The dataset this column is added to.
func (o ColumnOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *Column) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// A description that is shown in the UI.
func (o ColumnOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Column) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this column should be hidden in the query builder and sample data. Defaults to false.
func (o ColumnOutput) Hidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Column) pulumi.BoolPtrOutput { return v.Hidden }).(pulumi.BoolPtrOutput)
}

// The name of the column. Must be unique per dataset.
func (o ColumnOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Column) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// The type of the column, allowed values are `string`, `float`, `integer` and `boolean`. Defaults to `string`.
func (o ColumnOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Column) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnArrayOutput struct{ *pulumi.OutputState }

func (ColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Column)(nil)).Elem()
}

func (o ColumnArrayOutput) ToColumnArrayOutput() ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) Index(i pulumi.IntInput) ColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Column {
		return vs[0].([]*Column)[vs[1].(int)]
	}).(ColumnOutput)
}

type ColumnMapOutput struct{ *pulumi.OutputState }

func (ColumnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Column)(nil)).Elem()
}

func (o ColumnMapOutput) ToColumnMapOutput() ColumnMapOutput {
	return o
}

func (o ColumnMapOutput) ToColumnMapOutputWithContext(ctx context.Context) ColumnMapOutput {
	return o
}

func (o ColumnMapOutput) MapIndex(k pulumi.StringInput) ColumnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Column {
		return vs[0].(map[string]*Column)[vs[1].(string)]
	}).(ColumnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ColumnInput)(nil)).Elem(), &Column{})
	pulumi.RegisterInputType(reflect.TypeOf((*ColumnArrayInput)(nil)).Elem(), ColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ColumnMapInput)(nil)).Elem(), ColumnMap{})
	pulumi.RegisterOutputType(ColumnOutput{})
	pulumi.RegisterOutputType(ColumnArrayOutput{})
	pulumi.RegisterOutputType(ColumnMapOutput{})
}
