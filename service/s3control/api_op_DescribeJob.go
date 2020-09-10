// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the configuration parameters and status for a Batch Operations job.
// For more information, see Amazon S3 Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
// Amazon Simple Storage Service Developer Guide. Related actions include:
//
//     *
// CreateJob ()
//
//     * ListJobs ()
//
//     * UpdateJobPriority ()
//
//     *
// UpdateJobStatus ()
func (c *Client) DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) {
	stack := middleware.NewStack("DescribeJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDescribeJobMiddlewares(stack)
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
	addOpDescribeJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJob(options.Region), middleware.Before)

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
			OperationName: "DescribeJob",
			Err:           err,
		}
	}
	out := result.(*DescribeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobInput struct {
	// The ID for the job whose information you want to retrieve.
	JobId *string
	//
	AccountId *string
}

type DescribeJobOutput struct {
	// Contains the configuration parameters and status for the job specified in the
	// Describe Job request.
	Job *types.JobDescriptor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDescribeJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDescribeJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDescribeJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DescribeJob",
	}
}