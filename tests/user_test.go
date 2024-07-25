package tests_test

import (
	"mikea/declix/impl"
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var harness *Harness

	BeforeEach(func() {
		harness = NewHarness()
		err := harness.StartTarget()
		Expect(err).NotTo(HaveOccurred())

		DeferCleanup(func() {
			err := harness.CleanupTarget()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	It("users are processed", func() {
		app := harness.App

		err := app.LoadResources("user_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		err = app.ApplyActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// user 0 needs to be created
		Expect(app.Resources[0]).To(impl.HaveStyleStrings(
			"\x1b[32m999 users [users sudo]\x1b[0m",
			"\x1b[31mmissing\x1b[0m",
			"\x1b[32m+user:new_user\x1b[0m"))

		// user 1 needs to be deleted
		Expect(app.Resources[1]).To(impl.HaveStyleStrings(
			"\x1b[31mmissing\x1b[0m",
			"\x1b[32m1001 test_user [test_user users]\x1b[0m",
			"\x1b[31m-user:test_user\x1b[0m"))

		// user 2 needs to be updated
		Expect(app.Resources[2]).To(impl.HaveStyleStrings(
			"\x1b[32m2000 users [users sudo]\x1b[0m",
			"\x1b[32m1002 test_user2 [test_user2 users]\x1b[0m",
			"\x1b[33m~user:test_user2\x1b[0m"))

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// no actions should be expected
		Expect(app.HasActions()).To(BeFalse())
	})
})
