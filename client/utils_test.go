package client

import "testing"

func Test_generateUniqueInstructionID(t *testing.T) {
	id := GenerateUniqueInstructionID()
	t.Log(id)
}
