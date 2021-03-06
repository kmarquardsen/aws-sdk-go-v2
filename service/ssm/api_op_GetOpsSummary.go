// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// View a summary of OpsItems based on specified filters and aggregators.
func (c *Client) GetOpsSummary(ctx context.Context, params *GetOpsSummaryInput, optFns ...func(*Options)) (*GetOpsSummaryOutput, error) {
	if params == nil {
		params = &GetOpsSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpsSummary", params, optFns, addOperationGetOpsSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpsSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpsSummaryInput struct {

	// Optional aggregators that return counts of OpsItems based on one or more
	// expressions.
	Aggregators []*types.OpsAggregator

	// Optional filters used to scope down the returned OpsItems.
	Filters []*types.OpsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// The OpsItem data type to return.
	ResultAttributes []*types.OpsResultAttribute

	// Specify the name of a resource data sync to get.
	SyncName *string
}

type GetOpsSummaryOutput struct {

	// The list of aggregated and filtered OpsItems.
	Entities []*types.OpsEntity

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOpsSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetOpsSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpsSummary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetOpsSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetOpsSummary",
	}
}
