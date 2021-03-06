// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets the cache for a project.
func (c *Client) InvalidateProjectCache(ctx context.Context, params *InvalidateProjectCacheInput, optFns ...func(*Options)) (*InvalidateProjectCacheOutput, error) {
	if params == nil {
		params = &InvalidateProjectCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvalidateProjectCache", params, optFns, addOperationInvalidateProjectCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvalidateProjectCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvalidateProjectCacheInput struct {

	// The name of the AWS CodeBuild build project that the cache is reset for.
	//
	// This member is required.
	ProjectName *string
}

type InvalidateProjectCacheOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInvalidateProjectCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInvalidateProjectCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInvalidateProjectCache{}, middleware.After)
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
	addOpInvalidateProjectCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInvalidateProjectCache(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInvalidateProjectCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "InvalidateProjectCache",
	}
}
