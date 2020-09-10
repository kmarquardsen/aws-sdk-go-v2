// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an email address to the suppression list for your account.
func (c *Client) PutSuppressedDestination(ctx context.Context, params *PutSuppressedDestinationInput, optFns ...func(*Options)) (*PutSuppressedDestinationOutput, error) {
	stack := middleware.NewStack("PutSuppressedDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutSuppressedDestinationMiddlewares(stack)
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
	addOpPutSuppressedDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutSuppressedDestination(options.Region), middleware.Before)

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
			OperationName: "PutSuppressedDestination",
			Err:           err,
		}
	}
	out := result.(*PutSuppressedDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add an email destination to the suppression list for your account.
type PutSuppressedDestinationInput struct {
	// The email address that should be added to the suppression list for your account.
	EmailAddress *string
	// The factors that should cause the email address to be added to the suppression
	// list for your account.
	Reason types.SuppressionListReason
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutSuppressedDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutSuppressedDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutSuppressedDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSuppressedDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutSuppressedDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutSuppressedDestination",
	}
}