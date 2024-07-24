package tests_test

import (
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dir", func() {
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

	It("dirs are processed", func() {
		app := harness.App

		err := app.LoadResources("dir_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		err = app.ApplyActions()
		Expect(err).NotTo(HaveOccurred())

		// dir 0 needs to be created
		Expect(app.Expected[0].StyledString(app.Resources[0])).To(Equal("\x1b[32mtest_user:users 775\x1b[0m"))
		Expect(app.States[0].StyledString(app.Resources[0])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.Actions[0].StyledString(app.Resources[0])).To(Equal("\x1b[32m+dir:/tmp/new_dir\x1b[0m"))

		// dir 1 needs to be deleted
		Expect(app.Expected[1].StyledString(app.Resources[1])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.States[1].StyledString(app.Resources[1])).To(Equal("\x1b[32mroot:root 755\x1b[0m"))
		Expect(app.Actions[1].StyledString(app.Resources[1])).To(Equal("\x1b[31m-dir:/var/test_dir\x1b[0m"))

		// dir 2 needs to be updated
		Expect(app.Expected[2].StyledString(app.Resources[2])).To(Equal("\x1b[32mtest_user:users 777\x1b[0m"))
		Expect(app.States[2].StyledString(app.Resources[2])).To(Equal("\x1b[32mroot:root 755\x1b[0m"))
		Expect(app.Actions[2].StyledString(app.Resources[2])).To(Equal("\x1b[33m~dir:/var/test_dir2\x1b[0m"))

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		// no actions should be expected
		Expect(app.Actions[0]).To(BeNil())
		Expect(app.Actions[1]).To(BeNil())
		Expect(app.Actions[2]).To(BeNil())
	})
})
