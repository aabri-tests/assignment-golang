package e2e

import (
	"testing"
	"time"

	_ "github.com/aabri-assignments/assignment-golang/e2e/tests/health"
	_ "github.com/aabri-assignments/assignment-golang/e2e/tests/track"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestRunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory.
// This function is called on each Ginkgo node in parallel mode.
func TestRunE2ETests(t *testing.T) {
	RegisterFailHandler(Fail)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	suiteConfig.RandomSeed = time.Now().Unix()
	RunSpecs(t, "E2E Suite", suiteConfig, reporterConfig)
}
