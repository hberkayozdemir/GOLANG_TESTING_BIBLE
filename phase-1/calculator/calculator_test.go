package calculator_test

import (
	"testing"

	"github.com/hberkayozdemir/GOLANG_TESTING_BIBLE/calculator"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	/* t.Fail() */
	/* 	--- FAIL: TestCalculateIsArmstrong (0.00s)
	   	FAIL
	   	FAIL    github.com/hberkayozdemir/GOLANG_TESTING_BIBLE/calculator       0.001s
	   	FAIL */

	testCase := TestCase{
		value:    371,
		expected: true,
	}
	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}

	/*
	   === RUN   TestCalculateIsArmstrong
	   --- PASS: TestCalculateIsArmstrong (0.00s)
	   PASS
	   ok      github.com/hberkayozdemir/GOLANG_TESTING_BIBLE/calculator       0.001s */
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    350,
		expected: false,
	}
	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
