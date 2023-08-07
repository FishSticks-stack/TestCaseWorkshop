package main

// MLH DW: How to test your Code
// exploring various testing methodologies to ensure your code's accuracy, reliability, and performance

// need these imports to run in the terminal
import (
	"errors"
	"fmt"
)

type Order struct {
	ID       int
	Quantity int
	Price    float64
	// User struct {
	// 	CreditCard struct {
	// 		Number string
	// 		Expire string
	// 		Security string
	// 	}
	// }
}

type PaymentProcessor interface {
	ProcessPayment(orderID int, amount float64) error
}

type InventoryManager interface {
	ReduceStock(ItemID int, quantity int) error
}

type MyPaymentProcessor struct{}

func (p *MyPaymentProcessor) ProcessPayment(orderID int, amount float64) error {
	fmt.Printf("Processing payment for OrderID: %d, Amount: %.2f\n", orderID, amount)

	// add logic for the payment and payment gateway

	return nil
}

type MyInventoryManager struct{}

func (im *MyInventoryManager) ReduceStock(itemID int, quantity int) error {
	fmt.Printf("Reducing stock for ItemID: %d, Quantity: %d\n", itemID, quantity)
	return nil
}

func ProcessOrder(order Order, PaymentProcessor PaymentProcessor, InventoryManager InventoryManager) error {

	// verify the order
	if order.Quantity <= 0 {
		return errors.New("invalid order quantity")
	}

	totalAmount := order.Price * float64(order.Quantity)

	// process the payment

	if err := PaymentProcessor.ProcessPayment(order.ID, totalAmount); err != nil {
		return err
	}

	if err := InventoryManager.ReduceStock(order.ID, order.Quantity); err != nil {
		_ = PaymentProcessor.ProcessPayment(order.ID, -totalAmount)
		return err
	}

	return nil
}

func main() {
	order := Order{
		ID:       1,
		Quantity: 2,
		Price:    2.5,
	}

	paymentProcessor := &MyPaymentProcessor{}
	InventoryManager := &MyInventoryManager{}

	err := ProcessOrder(order, paymentProcessor, InventoryManager)

	if err != nil {
		fmt.Println("Error processing order: ", err)
	} else {
		fmt.Println("Order processed successfully!")
	}
}
