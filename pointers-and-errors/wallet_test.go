package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		// given
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		t.Run("happy", func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(10)}
			err := wallet.Withdraw(Bitcoin(1))
			assertBalance(t, wallet, Bitcoin(9))
			assertNoError(t, err)
		})

		t.Run("insufficient funds", func(t *testing.T) {
			initBalance := Bitcoin(0)
			wallet := Wallet{balance: initBalance}
			err := wallet.Withdraw(Bitcoin(1))
			assertBalance(t, wallet, initBalance)
			assertError(t, ErrInsufficientFunds, err)
		})
	})
}

func assertError(t *testing.T, want error, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("wanted no error but got one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}
