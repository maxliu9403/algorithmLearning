package isAnagram

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	v := isAnagram("rat", "car")
	fmt.Println(v)
}
