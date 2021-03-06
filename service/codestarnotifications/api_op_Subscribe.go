// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an association between a notification rule and an SNS topic so that the
// associated target can receive notifications when the events described in the
// rule are triggered.
func (c *Client) Subscribe(ctx context.Context, params *SubscribeInput, optFns ...func(*Options)) (*SubscribeOutput, error) {
	if params == nil {
		params = &SubscribeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Subscribe", params, optFns, addOperationSubscribeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubscribeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubscribeInput struct {

	// The Amazon Resource Name (ARN) of the notification rule for which you want to
	// create the association.
	//
	// This member is required.
	Arn *string

	// Information about the SNS topics associated with a notification rule.
	//
	// This member is required.
	Target *types.Target

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	ClientRequestToken *string
}

type SubscribeOutput struct {

	// The Amazon Resource Name (ARN) of the notification rule for which you have
	// created assocations.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSubscribeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSubscribe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSubscribe{}, middleware.After)
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
	addOpSubscribeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubscribe(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSubscribe(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "Subscribe",
	}
}
