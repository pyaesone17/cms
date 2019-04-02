package posts_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPosts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Posts Suite")
}
