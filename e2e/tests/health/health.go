package health

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ApiDescribe("Run health endpoint", func() {
	var (
		client    *http.Client
		serverURL string
	)

	BeforeEach(func() {
		client = &http.Client{}
		serverURL = "http://localhost:8080"
	})

	// Test cases:
	It("should return a 200 status code when hitting /_health", func() {
		resp, err := client.Get(serverURL + "/_health")
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})
})
