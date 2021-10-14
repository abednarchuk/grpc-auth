package validators

import "testing"

func TestValidateUserName(t *testing.T) {
	t.Run("Test valid usernames", func(t *testing.T) {
		usernames := []string{"use", "user", "user123", "user_123", "user123user_d"}
		for _, username := range usernames {
			got := ValidateUserName(username)
			want := true
			if got != want {
				t.Errorf("%q: got %v, want %v", username, got, want)
			}
		}
	})
	t.Run("Test invalid usernames", func(t *testing.T) {
		usernames := []string{"us", "user.", "user_", "_user", "1user", "tes_", "_test", "a2a4a6a8a11a14a17a20a23a26"}
		for _, username := range usernames {
			got := ValidateUserName(username)
			want := false

			if got != want {
				t.Errorf("%q: got %v, want %v", username, got, want)

			}
		}
	})
}

func TestValidateEmail(t *testing.T) {
	t.Run("Test valid email", func(t *testing.T) {
		emails := []string{"use@test.com", "user@user", "user.kappa@te.com"}
		for _, email := range emails {
			got := ValidateEmail(email)
			want := true
			if got != want {
				t.Errorf("%q: got %v, want %v", email, got, want)
			}
		}
	})
	t.Run("Test invalid email", func(t *testing.T) {
		emails := []string{"use@", "@user", "user a@te.com"}
		for _, email := range emails {
			got := ValidateEmail(email)
			want := false
			if got != want {
				t.Errorf("%q: got %v, want %v", email, got, want)
			}
		}
	})
}
func TestValidatePassword(t *testing.T) {
	t.Run("Test valid password", func(t *testing.T) {
		passwords := []string{"secret1"}
		for _, password := range passwords {
			got := ValidatePassword(password)
			want := true
			if got != want {
				t.Errorf("%q: got %v, want %v", password, got, want)
			}
		}
	})
	t.Run("Test invalid password", func(t *testing.T) {
		passwords := []string{"123"}
		for _, password := range passwords {
			got := ValidatePassword(password)
			want := false
			if got != want {
				t.Errorf("%q: got %v, want %v", password, got, want)
			}
		}
	})
}
