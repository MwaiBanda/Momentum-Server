package stripe

import (
	"context"
	"strings"
	"testing"

	applicationpayment "Momentum/application/service/payment"
)

func TestClientCreateRequiresSecretKey(t *testing.T) {
	client := NewClient("", "2024-06-20")

	_, err := client.Create(context.Background(), applicationpayment.Input{
		Amount: 100,
		Email:  "customer@example.com",
	})
	if err == nil || !strings.Contains(err.Error(), "not configured") {
		t.Fatalf("Create() error = %v, want unconfigured client error", err)
	}
}
