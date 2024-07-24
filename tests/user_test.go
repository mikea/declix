package tests_test

import (
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

		// user 0 needs to be created
		Expect(app.Expected[0].StyledString(app.Resources[0])).To(Equal("\x1b[32mnew_user 999 users [users sudo]\x1b[0m"))
		Expect(app.States[0].StyledString(app.Resources[0])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.Actions[0].StyledString(app.Resources[0])).To(Equal("\x1b[32m+user:new_user\x1b[0m"))

		// user 1 needs to be deleted
		Expect(app.Expected[1].StyledString(app.Resources[1])).To(Equal("\x1b[31mmissing\x1b[0m"))
		Expect(app.States[1].StyledString(app.Resources[1])).To(Equal("\x1b[32mtest_user 1001 test_user [test_user users]\x1b[0m"))
		Expect(app.Actions[1].StyledString(app.Resources[1])).To(Equal("\x1b[31m-user:test_user\x1b[0m"))

		// user 2 needs to be updated
		Expect(app.Expected[2].StyledString(app.Resources[2])).To(Equal("\x1b[32mtest_user2 2000 users [users sudo]\x1b[0m"))
		Expect(app.States[2].StyledString(app.Resources[2])).To(Equal("\x1b[32mtest_user2 1002 test_user2 [test_user2 users]\x1b[0m"))
		Expect(app.Actions[2].StyledString(app.Resources[2])).To(Equal("\x1b[33m~user:test_user2\x1b[0m"))

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
