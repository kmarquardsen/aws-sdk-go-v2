// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the comment for a specified traffic policy version.
func (c *Client) UpdateTrafficPolicyComment(ctx context.Context, params *UpdateTrafficPolicyCommentInput, optFns ...func(*Options)) (*UpdateTrafficPolicyCommentOutput, error) {
	if params == nil {
		params = &UpdateTrafficPolicyCommentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrafficPolicyComment", params, optFns, addOperationUpdateTrafficPolicyCommentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrafficPolicyCommentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the traffic policy that you want
// to update the comment for.
type UpdateTrafficPolicyCommentInput struct {

	// The new comment for the specified traffic policy and version.
	//
	// This member is required.
	Comment *string

	// The value of Id for the traffic policy that you want to update the comment for.
	//
	// This member is required.
	Id *string

	// The value of Version for the traffic policy that you want to update the comment
	// for.
	//
	// This member is required.
	Version *int32
}

// A complex type that contains the response information for the traffic policy.
type UpdateTrafficPolicyCommentOutput struct {

	// A complex type that contains settings for the specified traffic policy.
	//
	// This member is required.
	TrafficPolicy *types.TrafficPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTrafficPolicyCommentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateTrafficPolicyComment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateTrafficPolicyComment{}, middleware.After)
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
	addOpUpdateTrafficPolicyCommentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrafficPolicyComment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateTrafficPolicyComment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "UpdateTrafficPolicyComment",
	}
}
