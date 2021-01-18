package rdnaemu

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ALU", func() {

	var (
		state *mockInstState
	)

	BeforeEach(func() {
		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

})
