package consultantTests_test

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		StartMyApp(3000)

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("firefox"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:3000/login"))
		})

		By("allowing the user to fill out the login form and submit it", func() {
			Eventually(page.FindByLabel("E-mail")).Should(BeFound())
			Expect(page.FindByLabel("E-mail").Fill("spud@example.com")).To(Succeed())
			Expect(page.FindByLabel("Password").Fill("secret-password")).To(Succeed())
			Expect(page.Find("#remember_me").Check()).To(Succeed())
			Expect(page.Find("#login_form").Submit()).To(Succeed())
		})

		By("allowing the user to view their profile", func() {
			Eventually(page.FindByLink("Profile Page")).Should(BeFound())
			Expect(page.FindByLink("Profile Page").Click()).To(Succeed())
			profile := page.Find("section.profile")
			Eventually(profile.Find(".greeting")).Should(HaveText("Hello Spud!"))
			Expect(profile.Find("img#profile_pic")).To(BeVisible())
		})

		By("allowing the user to log out", func() {
			Expect(page.Find("#logout").Click()).To(Succeed())
			Expect(page).To(HavePopupText("Are you sure?"))
			Expect(page.ConfirmPopup()).To(Succeed())
			Eventually(page).Should(HaveTitle("Login"))
		})
	})
})
