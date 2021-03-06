// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a suite.
func (c *Client) GetSuite(ctx context.Context, params *GetSuiteInput, optFns ...func(*Options)) (*GetSuiteOutput, error) {
	if params == nil {
		params = &GetSuiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSuite", params, optFns, addOperationGetSuiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSuiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the get suite operation.
type GetSuiteInput struct {

	// The suite's ARN.
	//
	// This member is required.
	Arn *string
}

// Represents the result of a get suite request.
type GetSuiteOutput struct {

	// A collection of one or more tests.
	Suite *types.Suite

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSuiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSuite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSuite{}, middleware.After)
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
	addOpGetSuiteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSuite(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSuite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetSuite",
	}
}
