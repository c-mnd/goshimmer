package test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/iotaledger/goshimmer/packages/binary/signature/ed25119"

	"github.com/iotaledger/goshimmer/packages/ledgerstate/coloredcoins"

	"github.com/panjf2000/ants/v2"

	"github.com/iotaledger/goshimmer/packages/ledgerstate/transfer"

	"github.com/iotaledger/goshimmer/packages/binary/transaction"

	"github.com/iotaledger/goshimmer/packages/binary/address"
	"github.com/iotaledger/goshimmer/packages/binary/identity"
	"github.com/iotaledger/goshimmer/packages/binary/transaction/payload/data"
	"github.com/iotaledger/goshimmer/packages/binary/transaction/payload/valuetransfer"
	"github.com/stretchr/testify/assert"
)

func BenchmarkVerifySignature(b *testing.B) {
	transactions := make([]*transaction.Transaction, b.N)
	for i := 0; i < b.N; i++ {
		transactions[i] = transaction.New(transaction.EmptyId, transaction.EmptyId, identity.Generate(), data.New([]byte("test")))
		transactions[i].GetBytes()
	}

	var wg sync.WaitGroup

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		currentIndex := i
		if err := ants.Submit(func() {
			transactions[currentIndex].VerifySignature()

			wg.Done()
		}); err != nil {
			b.Error(err)
		}
	}

	wg.Wait()
}

func TestNew(t *testing.T) {
	newTransaction1 := transaction.New(transaction.EmptyId, transaction.EmptyId, identity.Generate(), data.New([]byte("test")))
	assert.Equal(t, newTransaction1.VerifySignature(), true)

	keyPairOfSourceAddress := ed25119.GenerateKeyPair()
	keyPairOfTargetAddress := ed25119.GenerateKeyPair()

	newValueTransaction1 := transaction.New(transaction.EmptyId, transaction.EmptyId, identity.Generate(),
		valuetransfer.New().
			AddInput(transfer.NewHash("test"), address.FromPublicKey(keyPairOfSourceAddress.PublicKey)).
			AddOutput(address.FromPublicKey(keyPairOfTargetAddress.PublicKey), coloredcoins.NewColoredBalance(coloredcoins.NewColor("IOTA"), 12)).
			Sign(keyPairOfSourceAddress),
	)
	assert.Equal(t, newValueTransaction1.VerifySignature(), true)

	newValueTransaction2, _ := transaction.FromBytes(newValueTransaction1.GetBytes())
	assert.Equal(t, newValueTransaction2.VerifySignature(), true)

	fmt.Println(newValueTransaction1.GetPayload().(*valuetransfer.ValueTransfer).MarshalBinary())
	fmt.Println(newValueTransaction2.GetPayload().(*valuetransfer.ValueTransfer).MarshalBinary())

	if newValueTransaction2.GetPayload().GetType() == valuetransfer.Type {
		fmt.Println(newValueTransaction1.GetPayload().(*valuetransfer.ValueTransfer).VerifySignatures())
		fmt.Println(newValueTransaction2.GetPayload().(*valuetransfer.ValueTransfer).VerifySignatures())
	}

	newTransaction2 := transaction.New(newTransaction1.GetId(), transaction.EmptyId, identity.Generate(), data.New([]byte("test1")))
	assert.Equal(t, newTransaction2.VerifySignature(), true)

	if newTransaction1.GetPayload().GetType() == data.Type {
		fmt.Println("DATA TRANSACTION")
	}

	newTransaction3, _ := transaction.FromBytes(newTransaction2.GetBytes())
	assert.Equal(t, newTransaction3.VerifySignature(), true)

	fmt.Println(newTransaction1)
	fmt.Println(newTransaction2)
	fmt.Println(newTransaction3)

	//fmt.Println(newValueTransaction1)
	//fmt.Println(newValueTransaction2)
}