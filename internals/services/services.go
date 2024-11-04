// Isolated business logic
package services

type AlgorithmPreProcessor interface{
	Preprocess() error
}

func (...) Preprocess() error {
	} 

type AlgorithmPostProcessor interface{
	Postprocess() error 
}

func (...) Postprocess() error {

}

