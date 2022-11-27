package tx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEx(t *testing.T) {
	initializeSM()

	type testCases struct {
		name        string
		tx          *Tx
		expectedErr error
	}

	scenarios := []testCases{
		{
			name: "happy path",
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
			name: "low balance error case",
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
		fmt.Printf("Scenario: %s\n", scenario.name)
		err := execute(scenario.tx)
		fmt.Printf("tx state after calling the execute: %+v , err=%+v\n\n", *scenario.tx, err)
		if scenario.expectedErr != nil {
			assert.Contains(t, err.Error(), scenario.expectedErr.Error(), "wrong error was thrown")
		} else {
			assert.Nil(t, err)
		}
	}
}
