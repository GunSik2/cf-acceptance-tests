package apps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"

	. "github.com/cloudfoundry/cf-acceptance-tests/helpers"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/cf"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/generator"
)

var _ = Describe("Application", func() {
	BeforeEach(func() {
		AppName = RandomName()

		Expect(Cf("push", AppName, "-p", TestAssets.Dora)).To(Say("App started"))
	})

	AfterEach(func() {
		Expect(Cf("delete", AppName, "-f")).To(Say("OK"))
	})

	Describe("pushing", func() {
		It("makes the app reachable via its bound route", func() {
			Eventually(Curling(AppName, "/", config.AppsDomain)).Should(Say("Hi, I'm Dora!"))
		})
	})

	Describe("stopping", func() {
		BeforeEach(func() {
			Expect(Cf("stop", AppName)).To(Say("OK"))
		})

		It("makes the app unreachable", func() {
			Eventually(Curling(AppName, "/", config.AppsDomain), 5.0).Should(Say("404"))
		})

		Describe("and then starting", func() {
			BeforeEach(func() {
				Expect(Cf("start", AppName)).To(Say("App started"))
			})

			It("makes the app reachable again", func() {
				Eventually(Curling(AppName, "/", config.AppsDomain)).Should(Say("Hi, I'm Dora!"))
			})
		})
	})

	Describe("updating", func() {
		It("is reflected through another push", func() {
			Eventually(Curling(AppName, "/", config.AppsDomain)).Should(Say("Hi, I'm Dora!"))

			Expect(Cf("push", AppName, "-p", TestAssets.HelloWorld)).To(Say("App started"))

			Eventually(Curling(AppName, "/", config.AppsDomain)).Should(Say("Hello, world!"))
		})
	})

	Describe("deleting", func() {
		BeforeEach(func() {
			Expect(Cf("delete", AppName, "-f")).To(Say("OK"))
		})

		It("removes the application", func() {
			Expect(Cf("app", AppName)).To(Say("not found"))
		})

		It("makes the app unreachable", func() {
			Eventually(Curling(AppName, "/", config.AppsDomain)).Should(Say("404"))
		})
	})
})
