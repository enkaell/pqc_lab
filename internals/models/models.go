package models

import (
	"github.com/google/uuid"
)

type Algorithm struct {
	AlgorithmId   uuid.UUID
	AlgorithmName string
}

type AlgorithmRun struct {
	AlgorithmId    uuid.UUID
	AlgorithmInfo1 string
	AlgorithmInfo2 string
	AlgorithmInfo3 string
}
