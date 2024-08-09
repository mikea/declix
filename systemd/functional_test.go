package systemd_test

import (
	"mikea/declix/impl"
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Systemd", func() {
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

	It("services are processed", func() {
		app := harness.App

		err := app.LoadResources("functional_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		err = app.ApplyActions()
		Expect(err).NotTo(HaveOccurred())

		// test1.service file
		Expect(app.Resources[0]).To(impl.HaveStyleStrings(
			"\x1b[32ma25b9053 root:root 644\x1b[0m",
			"\x1b[31mmissing\x1b[0m",
			"\x1b[32m+file:/lib/systemd/system/test1.service\x1b[0m"))

		// test1.service
		Expect(app.Resources[1]).To(impl.HaveStyleStrings(
			"\x1b[32menabled\x1b[0m",
			"\x1b[31mdisabled\x1b[0m \x1b[31minactive\x1b[0m",
			"\x1b[32m+\x1b[0mservice:test1"))

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// no actions should be expected
		Expect(app.Resources[0].Action).To(BeNil())
	})
})
