// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists operations that match the criteria that you specify.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if params == nil {
		params = &ListOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperations", params, optFns, addOperationListOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOperationsInput struct {

	// A complex type that contains specifications for the operations that you want to
	// list, for example, operations that you started between a specified start date
	// and end date. If you specify more than one filter, an operation must match all
	// filters to be returned by ListOperations.
	Filters []*types.OperationFilter

	// The maximum number of items that you want AWS Cloud Map to return in the
	// response to a ListOperations request. If you don't specify a value for
	// MaxResults, AWS Cloud Map returns up to 100 operations.
	MaxResults *int32

	// For the first ListOperations request, omit this value. If the response contains
	// NextToken, submit another ListOperations request to get the next group of
	// results. Specify the value of NextToken from the previous response in the next
	// request. AWS Cloud Map gets MaxResults operations and then filters them based on
	// the specified criteria. It's possible that no operations in the first MaxResults
	// operations matched the specified criteria but that subsequent groups of
	// MaxResults operations do contain operations that match the criteria.
	NextToken *string
}

type ListOperationsOutput struct {

	// If the response contains NextToken, submit another ListOperations request to get
	// the next group of results. Specify the value of NextToken from the previous
	// response in the next request. AWS Cloud Map gets MaxResults operations and then
	// filters them based on the specified criteria. It's possible that no operations
	// in the first MaxResults operations matched the specified criteria but that
	// subsequent groups of MaxResults operations do contain operations that match the
	// criteria.
	NextToken *string

	// Summary information about the operations that match the specified criteria.
	Operations []*types.OperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOperations{}, middleware.After)
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
	addOpListOperationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListOperations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "ListOperations",
	}
}
