// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Confirms a topic rule destination. When you create a rule requiring a
// destination, AWS IoT sends a confirmation message to the endpoint or base
// address you specify. The message includes a token which you pass back when
// calling ConfirmTopicRuleDestination to confirm that you own or have access to
// the endpoint.
func (c *Client) ConfirmTopicRuleDestination(ctx context.Context, params *ConfirmTopicRuleDestinationInput, optFns ...func(*Options)) (*ConfirmTopicRuleDestinationOutput, error) {
	if params == nil {
		params = &ConfirmTopicRuleDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmTopicRuleDestination", params, optFns, addOperationConfirmTopicRuleDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmTopicRuleDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConfirmTopicRuleDestinationInput struct {

	// The token used to confirm ownership or access to the topic rule confirmation
	// URL.
	//
	// This member is required.
	ConfirmationToken *string
}

type ConfirmTopicRuleDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConfirmTopicRuleDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConfirmTopicRuleDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConfirmTopicRuleDestination{}, middleware.After)
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
	addOpConfirmTopicRuleDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmTopicRuleDestination(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opConfirmTopicRuleDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ConfirmTopicRuleDestination",
	}
}
