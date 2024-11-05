package models

import (
	"github.com/google/uuid"
)

type Algorithm struct {
	AlgorithmId   uuid.UUID
	AlgorithmName string
}
type AlgorithmVersion struct {
    AlgorithmVersionId  uuid.UUID
    AlgorithmId  uuid.UUID
    AlgorithmVersionName
}
type AlgorithmVersionRun struct {
	AlgorithmVersionRunId  uuid.UUID
    AlgorithVersionId  uuid.UUID
	Keygen string
	Sign string
	Verify string
    PrivateKeySize unsigned int32 = 480
    PublicKeySize unsigned int32 = 288
    SignatureSize unsigned int32 = 96
}

