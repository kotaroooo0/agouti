package selection_test

import (
	"github.com/sclevine/agouti/matchers/internal/mocks"
	. "github.com/sclevine/agouti/matchers/internal/selection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BeEnabledMatcher", func() {
	var (
		matcher   *BeEnabledMatcher
		selection *mocks.Selection
	)

	BeforeEach(func() {
		selection = &mocks.Selection{}
		selection.StringCall.ReturnString = "CSS: #selector"
		matcher = &BeEnabledMatcher{}
	})

	Describe("#Match", func() {
		Context("when the actual object is a selection", func() {
			Context("when the element is enabled", func() {
				It("should successfully return true", func() {
					selection.EnabledCall.ReturnEnabled = true
					success, err := matcher.Match(selection)
					Expect(success).To(BeTrue())
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the element is not enabled", func() {
				It("should successfully return false", func() {
					selection.EnabledCall.ReturnEnabled = false
					success, err := matcher.Match(selection)
					Expect(success).To(BeFalse())
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})

		Context("when the actual object is not a selection", func() {
			It("should return an error", func() {
				_, err := matcher.Match("not a selection")
				Expect(err).To(MatchError("BeEnabled matcher requires a Selection.  Got:\n    <string>: not a selection"))
			})
		})
	})

	Describe("#FailureMessage", func() {
		It("should return a failure message", func() {
			selection.EnabledCall.ReturnEnabled = false
			matcher.Match(selection)
			message := matcher.FailureMessage(selection)
			Expect(message).To(Equal("Expected selection 'CSS: #selector' to be enabled"))
		})
	})

	Describe("#NegatedFailureMessage", func() {
		It("should return a negated failure message", func() {
			selection.EnabledCall.ReturnEnabled = true
			matcher.Match(selection)
			message := matcher.NegatedFailureMessage(selection)
			Expect(message).To(Equal("Expected selection 'CSS: #selector' not to be enabled"))
		})
	})
})
