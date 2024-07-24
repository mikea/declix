package tests_test

import (
	. "mikea/declix/tests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
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

	It("files are processed", func() {
		app := harness.App

		err := app.LoadResources("file_test.pkl")
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())

		err = app.ApplyActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// dir 0 needs to be created
		Expect(app.Expected[0].StyledString(app.Resources[0])).To(Equal("\x1b[32m9f86d081 test_user:users 775\x1b[0m"))
		Expect(app.States[0].StyledString(app.Resources[0])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.Actions[0].StyledString(app.Resources[0])).To(Equal("\x1b[32m+file:/tmp/new_file\x1b[0m"))

		// dir 1 needs to be deleted
		Expect(app.Expected[1].StyledString(app.Resources[1])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.States[1].StyledString(app.Resources[1])).To(Equal("\x1b[32me3b0c442 root:root 644\x1b[0m"))
		Expect(app.Actions[1].StyledString(app.Resources[1])).To(Equal("\x1b[31m-file:/var/test_file\x1b[0m"))

		// dir 2 needs to be updated
		Expect(app.Errors[2]).NotTo(HaveOccurred())
		Expect(app.Expected[2].StyledString(app.Resources[2])).To(Equal("\x1b[32m90a424ea test_user:users 777\x1b[0m"))
		Expect(app.States[2].StyledString(app.Resources[2])).To(Equal("\x1b[32m5891b5b5 root:root 644\x1b[0m"))
		Expect(app.Actions[2].StyledString(app.Resources[2])).To(Equal("\x1b[33m~file:/var/test_file2\x1b[0m"))

		err = app.DetermineStates()
		Expect(err).NotTo(HaveOccurred())

		err = app.DetermineActions()
		Expect(err).NotTo(HaveOccurred())
		Expect(app.HasErrors()).To(BeFalse())

		// no actions should be expected
		Expect(app.Actions[0]).To(BeNil())
		Expect(app.Actions[1]).To(BeNil())
		Expect(app.Actions[2]).To(BeNil())
	})
})
