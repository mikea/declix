package tests_test

import (
	"mikea/declix/impl"
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
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

	It("groups are processed", func() {
		app := harness.App

		err := app.LoadResources("group_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		err = app.ApplyActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// group 0 needs to be created
		Expect(app.Resources[0]).To(impl.HaveStyleStrings(
			"\x1b[32m2005\x1b[0m",
			"\x1b[31mmissing\x1b[0m",
			"\x1b[32m+group:new_group\x1b[0m"))

		// group 1 needs to be deleted
		Expect(app.Resources[1]).To(impl.HaveStyleStrings(
			"\x1b[31mmissing\x1b[0m",
			"\x1b[32m1000\x1b[0m",
			"\x1b[31m-group:test_group\x1b[0m"))

		// group 2 needs to be updated
		Expect(app.Resources[2]).To(impl.HaveStyleStrings(
			"\x1b[32m2000\x1b[0m",
			"\x1b[32m1003\x1b[0m",
			"\x1b[33m~group:test_group2\x1b[0m"))

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// no actions should be expected
		Expect(app.HasActions()).To(BeFalse())
	})
})
