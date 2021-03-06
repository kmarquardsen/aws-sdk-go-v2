// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new schema to your GraphQL API. This operation is asynchronous. Use to
// determine when it has completed.
func (c *Client) StartSchemaCreation(ctx context.Context, params *StartSchemaCreationInput, optFns ...func(*Options)) (*StartSchemaCreationOutput, error) {
	if params == nil {
		params = &StartSchemaCreationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSchemaCreation", params, optFns, addOperationStartSchemaCreationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSchemaCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSchemaCreationInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The schema definition, in GraphQL schema language format.
	//
	// This member is required.
	Definition []byte
}

type StartSchemaCreationOutput struct {

	// The current state of the schema (PROCESSING, FAILED, SUCCESS, or
	// NOT_APPLICABLE). When the schema is in the ACTIVE state, you can add data.
	Status types.SchemaStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartSchemaCreationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSchemaCreation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSchemaCreation{}, middleware.After)
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
	addOpStartSchemaCreationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSchemaCreation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartSchemaCreation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "StartSchemaCreation",
	}
}
