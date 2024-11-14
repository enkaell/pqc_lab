// Isolated business logic
package services

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/enkaell/pqc_lab/internals/repository"
	"github.com/enkaell/pqc_lab/internals/server"
)

type envError struct {
}

func (e envError) Error() string {
	return "Env vars is not set"
}

// Helper function to extract data using regular expressions
func extractField(input string, pattern string) string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return ""
}

func iterateOverAlgVersAndRunAlgs(v *server.Algorithm, output chan<- *server.Algorithm) {
	var runs []*server.Algorithm_Version_Run
	var versions []*server.Algorithm_Version
	BinDir := os.Getenv("BINDIR")
	if BinDir == "" {
		panic(envError{})
	}
	for _, element := range v.Versions {
		log.Printf("Algorithm %s version %s execution has been started \n", v.Name, element.Name)
		out, err := exec.Command(filepath.Join(BinDir, v.Name, "pqcsign_"+element.Name)).Output()
		if err != nil {
			// Skipping algorithm naming error or other inconsistences
			log.Printf("Command execution error: %s \nOutput:  %s \nAlgorithm: %s \n", err, out, element)
			continue
		}
		input := string(out)
		privateKeySize, _ := strconv.ParseUint(extractField(input, `Private key size:\s*(\d+)`), 10, 32)
		publicKeySize, _ := strconv.ParseUint(extractField(input, `Public key size:\s*(\d+)`), 10, 32)
		signatureSize, _ := strconv.ParseUint(extractField(input, `Signature size:\s*(\d+)`), 10, 32)
		keygenCycles := extractField(input, `Keygen:\s*(\d+) cycles`)
		signCycles := extractField(input, `Sign:\s*(\d+) cycles`)
		verifyCycles := extractField(input, `Verify:\s*(\d+) cycles`)
		// Populate the protobuf struct
		outParsed := &server.Algorithm_Version_Run{
			Name:           element.Name,
			Keygen:         keygenCycles + " cycles",
			Sign:           signCycles + " cycles",
			Verify:         verifyCycles + " cycles",
			PrivateKeySize: uint32(privateKeySize),
			PublicKeySize:  uint32(publicKeySize),
			SignatureSize:  uint32(signatureSize),
		}
		runs := append(runs, outParsed)
		versions := append(versions, &server.Algorithm_Version{Name: element.Name, Runs: runs})
		log.Printf("Sending data to the channel %v, channel buf size: %v \n", versions, len(out))
		output <- &server.Algorithm{
			Name:     string(v.Name),
			Versions: versions,
		}
	}
	err := repository.DBPrepareAndRunStatementTransaction(runs)
	if err != nil {
		log.Fatal("Database transaction error")
	}
}

// O(n^2)
// RunAlgorithms iterates over the algorithms and start goroutines for evaluation of each algorithm
func RunAlgorithms(algs *server.AlgorithmList, output chan *server.Algorithm) {
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		close(output)
	}()

	for _, v := range algs.Algorithms {
		wg.Add(1)
		log.Printf("Algorithm %s execution started \n", v.Name)
		go func(v *server.Algorithm) {
			defer wg.Done()
			iterateOverAlgVersAndRunAlgs(v, output)
		}(v)

	}
}
