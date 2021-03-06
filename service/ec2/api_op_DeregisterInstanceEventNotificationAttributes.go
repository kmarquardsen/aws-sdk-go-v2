// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters tag keys to prevent tags that have the specified tag keys from being
// included in scheduled event notifications for resources in the Region.
func (c *Client) DeregisterInstanceEventNotificationAttributes(ctx context.Context, params *DeregisterInstanceEventNotificationAttributesInput, optFns ...func(*Options)) (*DeregisterInstanceEventNotificationAttributesOutput, error) {
	if params == nil {
		params = &DeregisterInstanceEventNotificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterInstanceEventNotificationAttributes", params, optFns, addOperationDeregisterInstanceEventNotificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterInstanceEventNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterInstanceEventNotificationAttributesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Information about the tag keys to deregister.
	InstanceTagAttribute *types.DeregisterInstanceTagAttributeRequest
}

type DeregisterInstanceEventNotificationAttributesOutput struct {

	// The resulting set of tag keys.
	InstanceTagAttribute *types.InstanceTagNotificationAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterInstanceEventNotificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeregisterInstanceEventNotificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterInstanceEventNotificationAttributes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterInstanceEventNotificationAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeregisterInstanceEventNotificationAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterInstanceEventNotificationAttributes",
	}
}
