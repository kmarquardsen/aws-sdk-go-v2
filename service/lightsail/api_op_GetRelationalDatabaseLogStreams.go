// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of available log streams for a specific database in Amazon
// Lightsail.
func (c *Client) GetRelationalDatabaseLogStreams(ctx context.Context, params *GetRelationalDatabaseLogStreamsInput, optFns ...func(*Options)) (*GetRelationalDatabaseLogStreamsOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseLogStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseLogStreams", params, optFns, addOperationGetRelationalDatabaseLogStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseLogStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseLogStreamsInput struct {

	// The name of your database for which to get log streams.
	//
	// This member is required.
	RelationalDatabaseName *string
}

type GetRelationalDatabaseLogStreamsOutput struct {

	// An object describing the result of your get relational database log streams
	// request.
	LogStreams []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseLogStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseLogStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseLogStreams{}, middleware.After)
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
	addOpGetRelationalDatabaseLogStreamsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseLogStreams(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRelationalDatabaseLogStreams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseLogStreams",
	}
}
