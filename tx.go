package tx

import (
	"errors"

	"github.com/mychewcents/go-payments-engine/internal/sm"
)

var stateMachine *sm.StateMachine

var (
	errLowBalance = errors.New("low balance")
)

type Tx struct {
	amount       int
	senderName   string
	senderBal    int
	receiverName string
	receiverBal  int
}

func initializeSM() {
	stateMachine = sm.New()

	stateMachine.CreateRoute(0, deductSender, 1)
	stateMachine.CreateRoute(1, depositReceiver, 2)
}

func deductSender(sa *sm.CurrentState) error {
	if sa.State != 0 {
		return errors.New("tx not ready to deduct the sender")
	}

	tx := sa.Entity.(*Tx)

	if tx.amount > tx.senderBal {
		return errLowBalance
	}

	tx.senderBal -= tx.amount
	sa.State = 1

	sa.Entity = tx
	return nil
}

func depositReceiver(sa *sm.CurrentState) error {
	if sa.State != 1 {
		return errors.New("tx not ready to deposit the receiver")
	}

	tx := sa.Entity.(*Tx)

	tx.receiverBal += tx.amount

	sa.State = 2
	sa.Entity = tx
	return nil
}

func execute(tx *Tx) error {
	if stateMachine == nil {
		panic("state machine not defined")
	}

	if tx == nil {
		return errors.New("tx is nil")
	}

	currTxState := &sm.CurrentState{
		State:  0,
		Entity: tx,
	}

	// Deduct the sender
	if err := stateMachine.Run(currTxState); err != nil {
		return err
	}

	// Deposit the receiver
	if err := stateMachine.Run(currTxState); err != nil {
		return err
	}

	return nil
}
