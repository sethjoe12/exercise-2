package auth

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	type args struct {
		email    string
		username string
	}
	tests := []struct {
		name            string
		args            args
		wantTokenString string
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenString, err := GenerateJWT(tt.args.email, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTokenString != tt.wantTokenString {
				t.Errorf("GenerateJWT() = %v, want %v", gotTokenString, tt.wantTokenString)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		signedToken string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateToken(tt.args.signedToken); (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
