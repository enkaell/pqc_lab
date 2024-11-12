package models

import (
	"github.com/google/uuid"
)

type Algorithm struct {
	AlgorithmId   int32
	AlgorithmName string
}
type AlgorithmVersion struct {
	AlgorithmVersionId   int32
	AlgorithmId          int32
	AlgorithmVersionName string
}
type AlgorithmVersionRun struct {
	AlgorithmVersionRunId uuid.UUID
	AlgorithVersionId     uuid.UUID
	Keygen                string
	Sign                  string
	Verify                string
	PrivateKeySize        int32
	PublicKeySize         int32
	SignatureSize         int32
}
