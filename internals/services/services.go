// Isolated business logic
package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/enkaell/pqc_lab/internals/server"
)

type envError struct {
}

func (e envError) Error() string {
	return "Env vars is not set"
}

func IterateOverAlgsAndRunAlgVers(k int, v *server.Algorithm, output chan<- *server.Algorithm) error {
	BinDir := os.Getenv("BINDIR")
	if BinDir == "" {
		return envError{}
	}
	fmt.Println("INSIDE ITERATOR FUNC")
	for _, element := range v.Versions {
		out, err := exec.Command(filepath.Join(BinDir, v.Name, "pqcsign_"+element.Name)).Output()
		if err != nil {
			log.Fatalf("Command execution error: %s \nOutput:  %s \nAlgorithm: %s \n", err, out, element)
			continue
		}
		outStr := string(out)
		outParsed := &server.Algorithm_Version_Run{Name: outStr}
		formatedOut := &server.Algorithm{
			Id:   1,
			Name: string(element.Name),
			Versions: []*server.Algorithm_Version{
				{Name: element.Name, Runs: []*server.Algorithm_Version_Run{outParsed}}},
		}
		output <- formatedOut
	}
	return nil
}

// O(n^2)
// RunAlgorithms iterates over the algorithms and start goroutines for evaluation of each algorithm
func RunAlgorithms(algs *server.AlgorithmList, output chan *server.Algorithm) {
	wg := sync.WaitGroup{}
	fmt.Println("Performing algorithm")
	defer func() {
		wg.Wait()
		close(output)
	}()

	for k, v := range algs.Algorithms {
		wg.Add(1)
		fmt.Printf("Algorithm %s execution started \n", v.Name)
		go func(k int, v *server.Algorithm) {
			defer wg.Done()
			IterateOverAlgsAndRunAlgVers(k, v, output)
		}(k, v)

	}
}
