// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Performs a rollback of a transaction. Rolling back a transaction cancels its
// changes.
func (c *Client) RollbackTransaction(ctx context.Context, params *RollbackTransactionInput, optFns ...func(*Options)) (*RollbackTransactionOutput, error) {
	stack := middleware.NewStack("RollbackTransaction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRollbackTransactionMiddlewares(stack)
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
	addOpRollbackTransactionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRollbackTransaction(options.Region), middleware.Before)

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
			OperationName: "RollbackTransaction",
			Err:           err,
		}
	}
	out := result.(*RollbackTransactionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a request to perform a rollback of
// a transaction.
type RollbackTransactionInput struct {
	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	ResourceArn *string
	// The name or ARN of the secret that enables access to the DB cluster.
	SecretArn *string
	// The identifier of the transaction to roll back.
	TransactionId *string
}

// The response elements represent the output of a request to perform a rollback of
// a transaction.
type RollbackTransactionOutput struct {
	// The status of the rollback operation.
	TransactionStatus *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRollbackTransactionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRollbackTransaction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRollbackTransaction{}, middleware.After)
}

func newServiceMetadataMiddleware_opRollbackTransaction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "RollbackTransaction",
	}
}