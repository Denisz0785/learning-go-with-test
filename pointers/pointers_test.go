package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposite(Bitcoin(10))
		assertBalance(t, wallet, 10.0)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		want := Bitcoin(10)

		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, want)

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}

		err := wallet.Withdraw(Bitcoin(30))
		assertError(t, err, ErrInsufficientFound)
		assertBalance(t, wallet, startBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance

	if got != want {
		t.Errorf("want %s got %s", want, got)
		t.Log("error of balance")
	}

}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("wait of nol nil, but got %v", got)
	}
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("wanted nil but got %v", got)
	}
}
