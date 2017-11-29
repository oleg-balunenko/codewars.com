package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSumEvenFibonacciNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumEvenFibonacciNumbers Suite")
}
