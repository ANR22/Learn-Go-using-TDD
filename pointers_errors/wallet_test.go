package pointerserrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
	t.Run("withdraw sufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(10)
		assertNoError(t, err)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(100)

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
