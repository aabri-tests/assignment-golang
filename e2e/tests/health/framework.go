package health

import "github.com/onsi/ginkgo/v2"

// ApiDescribe annotates the test with the label.
func ApiDescribe(text string, body func()) bool {
	return ginkgo.Describe("[health] "+text, body)
}
