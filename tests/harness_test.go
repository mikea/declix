package tests_test

import (
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Harness", func() {
	var harness *Harness

	BeforeEach(func() {
		var err error

		harness = NewHarness()
		err = harness.StartTarget()
		Expect(err).NotTo(HaveOccurred())

		DeferCleanup(func() {
			err := harness.CleanupTarget()
			Expect(err).NotTo(HaveOccurred())
		})

	})

	It("can load empty resources", func() {
		app := harness.App

		err := app.LoadResources("empty_resources_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())
	})
})
