// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a traffic policy.
func (c *Client) DeleteTrafficPolicy(ctx context.Context, params *DeleteTrafficPolicyInput, optFns ...func(*Options)) (*DeleteTrafficPolicyOutput, error) {
	if params == nil {
		params = &DeleteTrafficPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrafficPolicy", params, optFns, addOperationDeleteTrafficPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a specified traffic policy version.
type DeleteTrafficPolicyInput struct {

	// The ID of the traffic policy that you want to delete.
	//
	// This member is required.
	Id *string

	// The version number of the traffic policy that you want to delete.
	//
	// This member is required.
	Version *int32
}

// An empty element.
type DeleteTrafficPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrafficPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteTrafficPolicy{}, middleware.After)
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
	addOpDeleteTrafficPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrafficPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTrafficPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteTrafficPolicy",
	}
}
