// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update a program in a multiplex.
func (c *Client) UpdateMultiplexProgram(ctx context.Context, params *UpdateMultiplexProgramInput, optFns ...func(*Options)) (*UpdateMultiplexProgramOutput, error) {
	if params == nil {
		params = &UpdateMultiplexProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMultiplexProgram", params, optFns, addOperationUpdateMultiplexProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a program in a multiplex.
type UpdateMultiplexProgramInput struct {

	// The ID of the multiplex of the program to update.
	//
	// This member is required.
	MultiplexId *string

	// The name of the program to update.
	//
	// This member is required.
	ProgramName *string

	// The new settings for a multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings
}

// Placeholder documentation for UpdateMultiplexProgramResponse
type UpdateMultiplexProgramOutput struct {

	// The updated multiplex program.
	MultiplexProgram *types.MultiplexProgram

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMultiplexProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMultiplexProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMultiplexProgram{}, middleware.After)
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
	addOpUpdateMultiplexProgramValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMultiplexProgram(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMultiplexProgram(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateMultiplexProgram",
	}
}
