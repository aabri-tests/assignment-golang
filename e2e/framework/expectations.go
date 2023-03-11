package framework

import "github.com/onsi/gomega"

// ExpectNoError checks if "err" is set, and if so, fails assertion while logging the error.
func ExpectNoError(err error, explain ...interface{}) {
	ExpectNoErrorWithOffset(1, err, explain...)
}

// ExpectNoErrorWithOffset checks if "err" is set, and if so, fails assertion while logging the error at "offset" levels above its caller
// (for example, for call chain f -> g -> ExpectNoErrorWithOffset(1, ...) error would be logged for "f").
func ExpectNoErrorWithOffset(offset int, err error, explain ...interface{}) {
	gomega.ExpectWithOffset(1+offset, err).NotTo(gomega.HaveOccurred(), explain...)
}

// ExpectEqual expects the specified two are the same, otherwise an exception raises
func ExpectEqual(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.Equal(extra), explain...)
}

// ExpectNotEqual expects the specified two are not the same, otherwise an exception raises
func ExpectNotEqual(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).NotTo(gomega.Equal(extra), explain...)
}

// ExpectErrorMatch ExpectMatchError expects an error happens and has a message matching the given string, otherwise an exception raises
func ExpectErrorMatch(err error, msg string, explain ...interface{}) {
	gomega.ExpectWithOffset(1, err).To(gomega.HaveOccurred(), explain...)
	gomega.ExpectWithOffset(1, err, explain...).To(gomega.MatchError(msg), explain...)
}

// ExpectConsistOf expects actual contains precisely the extra elements.  The ordering of the elements does not matter.
func ExpectConsistOf(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.ConsistOf(extra), explain...)
}

// ExpectHaveKey expects the actual map has the key in the keyset
func ExpectHaveKey(actual interface{}, key interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.HaveKey(key), explain...)
}

// ExpectEmpty expects actual is empty
func ExpectEmpty(actual interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.BeEmpty(), explain...)
}
