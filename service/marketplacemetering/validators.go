// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpBatchMeterUsage struct {
}

func (*validateOpBatchMeterUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchMeterUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchMeterUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchMeterUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMeterUsage struct {
}

func (*validateOpMeterUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMeterUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MeterUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMeterUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterUsage struct {
}

func (*validateOpRegisterUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResolveCustomer struct {
}

func (*validateOpResolveCustomer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResolveCustomer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResolveCustomerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResolveCustomerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchMeterUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchMeterUsage{}, middleware.After)
}

func addOpMeterUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMeterUsage{}, middleware.After)
}

func addOpRegisterUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterUsage{}, middleware.After)
}

func addOpResolveCustomerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResolveCustomer{}, middleware.After)
}

func validateUsageRecord(v *types.UsageRecord) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UsageRecord"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if v.Dimension == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Dimension"))
	}
	if v.CustomerIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CustomerIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUsageRecordList(v []*types.UsageRecord) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UsageRecordList"}
	for i := range v {
		if err := validateUsageRecord(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchMeterUsageInput(v *BatchMeterUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchMeterUsageInput"}
	if v.ProductCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductCode"))
	}
	if v.UsageRecords == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UsageRecords"))
	} else if v.UsageRecords != nil {
		if err := validateUsageRecordList(v.UsageRecords); err != nil {
			invalidParams.AddNested("UsageRecords", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMeterUsageInput(v *MeterUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MeterUsageInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if v.ProductCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductCode"))
	}
	if v.UsageDimension == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UsageDimension"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterUsageInput(v *RegisterUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterUsageInput"}
	if v.ProductCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductCode"))
	}
	if v.PublicKeyVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PublicKeyVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResolveCustomerInput(v *ResolveCustomerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResolveCustomerInput"}
	if v.RegistrationToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RegistrationToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
