// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
package cloudfrontkeyvaluestore
========
package cleanroomsml
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
	// Access denied.
========
	// You do not have sufficient access to perform this action.
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
	// Resource is not in expected state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Internal server error.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Resource was not found.
========
	// You can't complete this action because another resource depends on this resource.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource you are requesting does not exist.
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
	// Limit exceeded.
========
	// You have exceeded your service quota.
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
	// Validation failed.
========
	// The request parameters for this request are incorrect.
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
<<<<<<<< HEAD:vendor/src/github.com/aws/aws-sdk-go/service/cloudfrontkeyvaluestore/errors.go
	"InternalServerException":       newErrorInternalServerException,
========
>>>>>>>> upstream/mainline:vendor/src/github.com/aws/aws-sdk-go/service/cleanroomsml/errors.go
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ValidationException":           newErrorValidationException,
}
