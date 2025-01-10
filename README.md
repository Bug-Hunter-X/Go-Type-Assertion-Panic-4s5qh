# Go Type Assertion Panic

This repository demonstrates a common error in Go: panicking due to an unsuccessful type assertion. The `bug.go` file contains the erroneous code, while `bugSolution.go` provides a corrected version.

The issue arises from attempting to convert an integer to a string using type assertion without checking if the assertion is valid.  This can lead to runtime crashes in production environments. The solution showcases how to handle this gracefully using a type switch or conditional check before assertion.