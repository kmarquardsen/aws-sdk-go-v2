// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Recovers the saved resource identified by an Amazon Resource Name (ARN). If the
// resource ARN is included in the request, then the last complete backup of that
// resource is recovered. If the ARN of a recovery point is supplied, then that
// recovery point is restored.
func (c *Client) StartRestoreJob(ctx context.Context, params *StartRestoreJobInput, optFns ...func(*Options)) (*StartRestoreJobOutput, error) {
	if params == nil {
		params = &StartRestoreJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRestoreJob", params, optFns, addOperationStartRestoreJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRestoreJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRestoreJobInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that AWS Backup uses to create
	// the target recovery point; for example, arn:aws:iam::123456789012:role/S3Access.
	//
	// This member is required.
	IamRoleArn *string

	// A set of metadata key-value pairs. Contains information, such as a resource
	// name, required to restore a recovery point. You can get configuration metadata
	// about a resource at the time it was backed up by calling
	// GetRecoveryPointRestoreMetadata. However, values in addition to those provided
	// by GetRecoveryPointRestoreMetadata might be required to restore a resource. For
	// example, you might need to provide a new resource name if the original already
	// exists. You need to specify specific metadata to restore an Amazon Elastic File
	// System (Amazon EFS) instance:
	//
	//     * file-system-id: ID of the Amazon EFS file
	// system that is backed up by AWS Backup. Returned in
	// GetRecoveryPointRestoreMetadata.
	//
	//     * Encrypted: A Boolean value that, if
	// true, specifies that the file system is encrypted. If KmsKeyId is specified,
	// Encrypted must be set to true.
	//
	//     * KmsKeyId: Specifies the AWS KMS key that
	// is used to encrypt the restored file system.
	//
	//     * PerformanceMode: Specifies
	// the throughput mode of the file system.
	//
	//     * CreationToken: A user-supplied
	// value that ensures the uniqueness (idempotency) of the request.
	//
	//     *
	// newFileSystem: A Boolean value that, if true, specifies that the recovery point
	// is restored to a new Amazon EFS file system.
	//
	// This member is required.
	Metadata map[string]*string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	//
	// This member is required.
	RecoveryPointArn *string

	// A customer chosen string that can be used to distinguish between calls to
	// StartRestoreJob.
	IdempotencyToken *string

	// Starts a job to restore a recovery point for one of the following resources:
	//
	//
	// * DynamoDB for Amazon DynamoDB
	//
	//     * EBS for Amazon Elastic Block Store
	//
	//     *
	// EC2 for Amazon Elastic Compute Cloud
	//
	//     * EFS for Amazon Elastic File System
	//
	//
	// * RDS for Amazon Relational Database Service
	//
	//     * Storage Gateway for AWS
	// Storage Gateway
	ResourceType *string
}

type StartRestoreJobOutput struct {

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartRestoreJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartRestoreJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartRestoreJob{}, middleware.After)
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
	addOpStartRestoreJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartRestoreJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartRestoreJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartRestoreJob",
	}
}
