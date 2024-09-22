package check_for_doubles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{12, 2, 3, 3, 7, 8, 10, 6} // Input array
	var expectedFlag = true
	var flag = checkIfExist(nums) // Calls your implementation

	assert.Equal(t, expectedFlag, flag)
}
func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); {
			if arr[i] == arr[j]*2 || arr[j] == arr[i]*2 {
				return true
			}
			j++
		}
	}
	return false
}
