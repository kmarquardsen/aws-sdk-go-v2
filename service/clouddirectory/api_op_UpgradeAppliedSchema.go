// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Upgrades a single directory in-place using the PublishedSchemaArn with schema
// updates found in MinorVersion. Backwards-compatible minor version upgrades are
// instantaneously available for readers on all objects in the directory. Note:
// This is a synchronous API call and upgrades only one schema on a given directory
// per call. To upgrade multiple directories from one schema, you would need to
// call this API on each directory.
func (c *Client) UpgradeAppliedSchema(ctx context.Context, params *UpgradeAppliedSchemaInput, optFns ...func(*Options)) (*UpgradeAppliedSchemaOutput, error) {
	stack := middleware.NewStack("UpgradeAppliedSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpgradeAppliedSchemaMiddlewares(stack)
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
	addOpUpgradeAppliedSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradeAppliedSchema(options.Region), middleware.Before)

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
			OperationName: "UpgradeAppliedSchema",
			Err:           err,
		}
	}
	out := result.(*UpgradeAppliedSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpgradeAppliedSchemaInput struct {
	// The ARN for the directory to which the upgraded schema will be applied.
	DirectoryArn *string
	// The revision of the published schema to upgrade the directory to.
	PublishedSchemaArn *string
	// Used for testing whether the major version schemas are backward compatible or
	// not. If schema compatibility fails, an exception would be thrown else the call
	// would succeed but no changes will be saved. This parameter is optional.
	DryRun *bool
}

type UpgradeAppliedSchemaOutput struct {
	// The ARN of the directory that is returned as part of the response.
	DirectoryArn *string
	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpgradeAppliedSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpgradeAppliedSchema{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradeAppliedSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpgradeAppliedSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpgradeAppliedSchema",
	}
}