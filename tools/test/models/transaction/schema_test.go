package transaction

import (
	"fmt"
	"testing"

	"github.com/hoxito/walletgo/models/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {

	var gotTransaction = transaction.NewTransaction()

	assert.NotEmpty(t, gotTransaction)
	assert.NotEmpty(t, gotTransaction.TransactionId)
	assert.NotEmpty(t, gotTransaction.Created)
	assert.Equal(t, "Created", gotTransaction.Status)
	assert.Empty(t, gotTransaction.Amount)
	assert.Empty(t, gotTransaction.DestinationId)
	assert.Empty(t, gotTransaction.OriginId)

}
func TestValidateTrReqErrors(t *testing.T) {

	var tests = []struct {
		transaction transaction.TransactionRequest
		want        string
	}{
		//Test Case Transaction Request with 0 amount
		{transaction.TransactionRequest{
			OriginId:      "someid",
			DestinationId: "anotherid",
			Amount:        00.00,
		},

			//Result
			"Key: 'TransactionRequest.Amount' Error:Field validation for 'Amount' failed on the 'required' tag",
		},

		//Test Case Transaction Request with 0.5 amount
		{transaction.TransactionRequest{
			OriginId:      "someid",
			DestinationId: "anotherid",
			Amount:        00.50,
		},

			//Result
			"Key: 'TransactionRequest.Amount' Error:Field validation for 'Amount' failed on the 'gte' tag",
		},

		//Test Case Transaction Request with 100000000000.00 amount
		{transaction.TransactionRequest{
			OriginId:      "someid",
			DestinationId: "anotherid",
			Amount:        100000000000.00,
		},

			//Result
			"Key: 'TransactionRequest.Amount' Error:Field validation for 'Amount' failed on the 'lte' tag",
		},
		//Test Case Transaction Request without OriginId
		{transaction.TransactionRequest{
			OriginId:      "",
			DestinationId: "anotherid",
			Amount:        100.00,
		},

			//Result
			"Key: 'TransactionRequest.OriginId' Error:Field validation for 'OriginId' failed on the 'required' tag",
		},
		//Test Case valid Transaction Request
		{transaction.TransactionRequest{
			OriginId:      "someId",
			DestinationId: "anotherid",
			Amount:        100.00,
		},

			//Result
			"",
		},
		//Test Case Transaction Request with same origin and destination
		{transaction.TransactionRequest{
			OriginId:      "someId",
			DestinationId: "someId",
			Amount:        100.00,
		},

			//Result
			"Key: 'TransactionRequest.OriginId' Error:Field validation for 'OriginId' failed on the 'necsfield' tag",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Testing Validate Transaction Request with (%v,%v,%v)",
			tt.transaction.OriginId,
			tt.transaction.DestinationId,
			tt.transaction.Amount)
		t.Run(testname, func(t *testing.T) {
			ans := tt.transaction.ValidateSchema()

			if ans == nil {
				assert.Equal(t, tt.want, "")
			} else {
				assert.Equal(t, tt.want, ans.Error())
			}
			// if ans != tt.want {
			// 	t.Errorf("got %d, want %d", ans, tt.want)
			// }
		})
	}
}
