// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes permissions for the specified private image. After you delete
// permissions for an image, AWS accounts to which you previously granted these
// permissions can no longer use the image.
func (c *Client) DeleteImagePermissions(ctx context.Context, params *DeleteImagePermissionsInput, optFns ...func(*Options)) (*DeleteImagePermissionsOutput, error) {
	stack := middleware.NewStack("DeleteImagePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteImagePermissionsMiddlewares(stack)
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
	addOpDeleteImagePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteImagePermissions(options.Region), middleware.Before)

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
			OperationName: "DeleteImagePermissions",
			Err:           err,
		}
	}
	out := result.(*DeleteImagePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteImagePermissionsInput struct {
	// The 12-digit identifier of the AWS account for which to delete image
	// permissions.
	SharedAccountId *string
	// The name of the private image.
	Name *string
}

type DeleteImagePermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteImagePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteImagePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteImagePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteImagePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DeleteImagePermissions",
	}
}