// Isolated business logic
package services

import (
	"github.com/enkaell/pqc_lab/internals/server"
)

type AlgorithmPreProcessor interface {
	Preprocess() error
}

func getAllAlgVersionByAlgName(name string) (*server.AlgorithmVersionList, error) {

	return nil, nil
}
