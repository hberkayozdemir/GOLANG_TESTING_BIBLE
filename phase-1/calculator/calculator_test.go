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

	t.Run("should return true for 371",func(t *testing.T){
		testCase := TestCase{
		value:    371,
		expected: true,
	}	
	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
	})


	t.Run("should return true for 370",func(t *testing.T){
		testCase := TestCase{
			value:    370,
			expected: true,
		}	
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}

	})



	/*
	   === RUN   TestCalculateIsArmstrong
	   --- PASS: TestCalculateIsArmstrong (0.00s)
	   PASS
	   ok      github.com/hberkayozdemir/GOLANG_TESTING_BIBLE/calculator       0.001s */
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {

	t.Run("Should fail for case 350",func(t *testing.T){

	testCase := TestCase{
		value:    350,
		expected: false,
	}
	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
	})

	t.Run("Should fail for case 300",func(t *testing.T){

		testCase := TestCase{
			value:    300,
			expected: false,
		}
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
		})
}
