package msgutils

import (
. "github.com/onsi/gomega"
"testing"
)

func TestReverse(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Reverse("abcd")).Should(Equal("dcba"))
	g.Expect(Reverse("")).Should(Equal(""))
}

func TestIsPalindrome(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(IsPalindrome("")).To(BeFalse())
	g.Expect(IsPalindrome("ABBA")).To(BeTrue())
	g.Expect(IsPalindrome("abba")).To(BeTrue())
	g.Expect(IsPalindrome("ABBa")).To(BeTrue())
	g.Expect(IsPalindrome("abba ")).To(BeFalse())
	g.Expect(IsPalindrome("Aba aBa")).To(BeTrue())
	g.Expect(IsPalindrome("Test test")).To(BeFalse())
}
