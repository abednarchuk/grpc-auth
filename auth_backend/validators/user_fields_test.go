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
