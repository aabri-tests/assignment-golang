### Run the e2e tests

Make sure you have ginkgo installed on your local machine:
```
go get github.com/onsi/ginkgo/ginkgo
```

#### Run all e2e tests
```
# Install ginkgo and run in this folder
go install github.com/onsi/ginkgo/v2/ginkgo && \
ginkgo ./...
```

#### Run a specific e2e test
```
# Install ginkgo and run in this folder
ginkgo -focus="TEST One"
```
