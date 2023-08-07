package main

// to run: >go mod init TestCaseWorkshop (will add a go.mod file)
// >go mod tidy (adds TestCaseWorkshop file) >go build
// >./TestCaseWorkshop (finally runs the code)

import (
	"errors"
	"testing"
)

type MockPaymentProcessor struct {
	processError error
}

func (p *MockPaymentProcessor) processPayment(orderID int, amount float64) error {
	return p.processError
}

type MockInventoryManager struct {
	ReduceStockError error
}

func (im *MockInventoryManager) ReduceStock(ItemID int, quantity int) error {
	return im.ReduceStockError
}

func TestProcessesOrder(t *testing.T) {
	order := Order{
		ID:       1,
		Quantity: 2,
		Price:    10}

	t.Run("Valid order", func(t *testing.T) {
		PaymentProcessor := &MyPaymentProcessor{}
		InventoryManager := &MyInventoryManager{}

		err := ProcessOrder(order, PaymentProcessor, InventoryManager)

		if err != nil {
			t.Errorf("Processing order returned an error for a valid order: %s", err)
		}
	})

	t.Run("Invalid order quantity", func(t *testing.T) {
		invalidOrder := Order{
			ID:       2,
			Quantity: 0,
			Price:    5.0,
		}

		PaymentProcessor := &MyPaymentProcessor{}
		InventoryManager := &MyInventoryManager{}

		err := ProcessOrder(invalidOrder, PaymentProcessor, InventoryManager)

		if err == nil {
			t.Error("ProcessOrder did not return an error for an invalid order quantity.")
		}
		expectError := "invalid order quantity"

		if err.Error() != expectError {
			t.Errorf("ProcessOrder returned the wrong error for an invalid order quantity.")
		}
	})

	t.Run("Inventory reduction error", func(t *testing.T) {
		PaymentProcessor := &MyPaymentProcessor{}
		InventoryManager := &MockInventoryManager{
			ReduceStockError: errors.New("inventory reduction error"),
		}

		err := ProcessOrder(order, PaymentProcessor, InventoryManager)
		if err == nil {
			t.Error("ProcessOrder did not return an error for Inventory reduction error.")
		}

		expectedError := "inventory reduction error"

		if err.Error() != expectedError {
			t.Error("ProcessOrder returned the wrong error for an inventory reduction.")
		}
	})
}
