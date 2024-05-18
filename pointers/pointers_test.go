package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.balance

		if got != want {
			t.Errorf("want %s got %s", want, got)
			t.Log("error of balance")
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposite(Bitcoin(10))
		assertBalance(t, wallet, 10.0)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		want := Bitcoin(10)

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, want)

	})
}
