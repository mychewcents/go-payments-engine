package tx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEx(t *testing.T) {
	initializeSM()

	type testCases struct {
		tx          *Tx
		expectedErr error
	}

	scenarios := []testCases{
		{
			tx: &Tx{
				amount:       10,
				senderName:   "Bob",
				senderBal:    20,
				receiverName: "Alice",
				receiverBal:  10,
			},
			expectedErr: nil,
		},
		{ // Low Balance error
			tx: &Tx{
				amount:       100,
				senderName:   "Bob",
				senderBal:    20,
				receiverName: "Alice",
				receiverBal:  10,
			},
			expectedErr: errLowBalance,
		},
	}

	for _, scenario := range scenarios {
		err := execute(scenario.tx)
		if scenario.expectedErr != nil {
			assert.Contains(t, err.Error(), scenario.expectedErr.Error(), "wrong error was thrown")
		} else {
			assert.Nil(t, err)
		}
	}
}
