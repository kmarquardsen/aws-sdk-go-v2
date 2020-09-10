// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the status of a specific version of a message template to active.
func (c *Client) UpdateTemplateActiveVersion(ctx context.Context, params *UpdateTemplateActiveVersionInput, optFns ...func(*Options)) (*UpdateTemplateActiveVersionOutput, error) {
	stack := middleware.NewStack("UpdateTemplateActiveVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateTemplateActiveVersionMiddlewares(stack)
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
	addOpUpdateTemplateActiveVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplateActiveVersion(options.Region), middleware.Before)

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
			OperationName: "UpdateTemplateActiveVersion",
			Err:           err,
		}
	}
	out := result.(*UpdateTemplateActiveVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateActiveVersionInput struct {
	// The type of channel that the message template is designed for. Valid values are:
	// EMAIL, PUSH, SMS, and VOICE.
	TemplateType *string
	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	TemplateName *string
	// Specifies which version of a message template to use as the active version of
	// the template.
	TemplateActiveVersionRequest *types.TemplateActiveVersionRequest
}

type UpdateTemplateActiveVersionOutput struct {
	// Provides information about an API request or response.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateTemplateActiveVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplateActiveVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplateActiveVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTemplateActiveVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateTemplateActiveVersion",
	}
}