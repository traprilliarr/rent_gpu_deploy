package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

func NewValidator(viper *viper.Viper) *validator.Validate {

	validate := validator.New()
	err := validate.RegisterValidation("customEmail", CustomEmailValidator)
	if err != nil {
		return nil
	}
	_ = validate.RegisterValidation("customTelegram", CustomTelegramValidator)
	if err != nil {
		return nil
	}
	_ = validate.RegisterValidation("customSSHKey", CustomSSHKeyValidator)
	if err != nil {
		return nil
	}
	return validate
}

// CustomEmailValidator checks if the email is in a valid format
func CustomEmailValidator(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	// Simple email format validation
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// CustomTelegramValidator checks if the Telegram username starts with "@"
func CustomTelegramValidator(fl validator.FieldLevel) bool {
	telegram := fl.Field().String()
	return strings.HasPrefix(telegram, "@")
}

// CustomSSHKeyValidator checks if the SSH key starts with one of the given prefixes
func CustomSSHKeyValidator(fl validator.FieldLevel) bool {
	sshKey := fl.Field().String()
	// Example prefixes typically used for SSH keys
	validPrefixes := []string{"ssh-rsa", "ssh-ed25519", "ecdsa-sha2-nistp256", "ecdsa-sha2-nistp384", "ecdsa-sha2-nistp521"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(sshKey, prefix) {
			return true
		}
	}
	return false
}
