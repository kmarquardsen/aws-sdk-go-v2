// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an Amazon S3 presigned URL for an update file associated with a
// specified JobId.
func (c *Client) GetSoftwareUpdates(ctx context.Context, params *GetSoftwareUpdatesInput, optFns ...func(*Options)) (*GetSoftwareUpdatesOutput, error) {
	stack := middleware.NewStack("GetSoftwareUpdates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSoftwareUpdatesMiddlewares(stack)
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
	addOpGetSoftwareUpdatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSoftwareUpdates(options.Region), middleware.Before)

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
			OperationName: "GetSoftwareUpdates",
			Err:           err,
		}
	}
	out := result.(*GetSoftwareUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSoftwareUpdatesInput struct {
	// The ID for a job that you want to get the software update file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string
}

type GetSoftwareUpdatesOutput struct {
	// The Amazon S3 presigned URL for the update file associated with the specified
	// JobId value. The software update will be available for 2 days after this request
	// is made. To access an update after the 2 days have passed, you'll have to make
	// another call to GetSoftwareUpdates.
	UpdatesURI *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSoftwareUpdatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSoftwareUpdates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSoftwareUpdates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSoftwareUpdates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetSoftwareUpdates",
	}
}