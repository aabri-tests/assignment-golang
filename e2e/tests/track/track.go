package track

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aabri-assignments/assignment-golang/e2e/framework"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ApiDescribe("Run track endpoint", func() {
	var (
		client    *http.Client
		serverURL string
	)

	BeforeEach(func() {
		client = &http.Client{}
		serverURL = "http://localhost:8080"
	})

	// Test cases:
	It(`PAYLOAD !: should return ["SFO", "EWR"]`, func() {
		assignment, err := os.ReadFile(filepath.Join("tests/track/testdata/", "payload1.json"))
		Expect(err).NotTo(HaveOccurred())
		resp, err := client.Post(serverURL+"/track", "application/json", bytes.NewBuffer(assignment))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		body, _ := io.ReadAll(resp.Body)
		framework.ExpectEqual(string(body), `{"source":"SFO","destination":"EWR"}`)
	})
	It(`PAYLOAD 2: should return ["SFO", "EWR"]`, func() {
		assignment, err := os.ReadFile(filepath.Join("tests/track/testdata/", "payload2.json"))
		Expect(err).NotTo(HaveOccurred())
		resp, err := client.Post(serverURL+"/track", "application/json", bytes.NewBuffer(assignment))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		body, _ := io.ReadAll(resp.Body)
		framework.ExpectEqual(string(body), `{"source":"SFO","destination":"EWR"}`)
	})
	It(`PAYLOAD 3: should return ["SFO", "EWR"]`, func() {
		assignment, err := os.ReadFile(filepath.Join("tests/track/testdata/", "payload3.json"))
		Expect(err).NotTo(HaveOccurred())
		resp, err := client.Post(serverURL+"/track", "application/json", bytes.NewBuffer(assignment))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		body, _ := io.ReadAll(resp.Body)
		framework.ExpectEqual(string(body), `{"source":"SFO","destination":"EWR"}`)
	})
	It("should return 500 Internal Server Error", func() {
		assignment, err := os.ReadFile(filepath.Join("tests/track/testdata/", "invalid.json"))
		Expect(err).NotTo(HaveOccurred())
		resp, err := client.Post(serverURL+"/track", "application/json", bytes.NewBuffer(assignment))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
	})
	It("should return an empty value", func() {
		assignment, err := os.ReadFile(filepath.Join("tests/track/testdata/", "empty.json"))
		Expect(err).NotTo(HaveOccurred())
		resp, err := client.Post(serverURL+"/track", "application/json", bytes.NewBuffer(assignment))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		body, _ := io.ReadAll(resp.Body)
		framework.ExpectEqual(string(body), `{"source":"","destination":""}`)
	})
})
