// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns compliance details for the conformance pack based on the cumulative
// compliance results of all the rules in that conformance pack.
func (c *Client) GetConformancePackComplianceSummary(ctx context.Context, params *GetConformancePackComplianceSummaryInput, optFns ...func(*Options)) (*GetConformancePackComplianceSummaryOutput, error) {
	stack := middleware.NewStack("GetConformancePackComplianceSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetConformancePackComplianceSummaryMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetConformancePackComplianceSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConformancePackComplianceSummary(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetConformancePackComplianceSummary",
			Err:           err,
		}
	}
	out := result.(*GetConformancePackComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConformancePackComplianceSummaryInput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of conformance packs returned on each page.
	Limit *int32
	// Names of conformance packs.
	ConformancePackNames []*string
}

type GetConformancePackComplianceSummaryOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// A list of ConformancePackComplianceSummary objects.
	ConformancePackComplianceSummaryList []*types.ConformancePackComplianceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetConformancePackComplianceSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetConformancePackComplianceSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConformancePackComplianceSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetConformancePackComplianceSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetConformancePackComplianceSummary",
	}
}