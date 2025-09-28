package main

import (
	"fmt"
	"os"

	"github.com/mondegor/go-sysmess/mrlib/crypt"
	"github.com/mondegor/go-sysmess/mrlog/litelog"
	"github.com/mondegor/go-sysmess/mrlog/slog"
)

func main() {
	l, _ := slog.NewLoggerAdapter(slog.WithWriter(os.Stdout))
	logger := litelog.NewLogger(l)

	value, _ := crypt.GenerateDigits(16)
	logger.Debug("GenerateDigits", "value", value)

	value, _ = crypt.GenerateHex(16)
	logger.Info("GenTokenHex", "value", value)

	value, _ = crypt.GenerateToken(64)
	logger.Info("GenerateToken", "value", value)

	valueBytes, _ := crypt.GenerateBytes([]byte("abc123.,:"), 16)
	logger.Info("GenerateBytes", "value", string(valueBytes))

	pwgen := crypt.NewPasswordGenerator()

	logger.Info("GenPassword", "password", pwgen.Generate(16, crypt.PassAll))

	pw := pwgen.Generate(12, crypt.PassAbc)
	logger.Info("PasswordStrength 12 abc", "password", pw, "strength", crypt.PasswordStrength(pw))

	pw = pwgen.Generate(9, crypt.PassAbcNumerals)
	logger.Info("PasswordStrength 9 abc+num", "password", pw, "strength", crypt.PasswordStrength(pw))

	pw = pwgen.Generate(12, crypt.PassAll)
	logger.Info("PasswordStrength 12 all", "password", pw, "strength", crypt.PasswordStrength(pw))

	fmt.Println(crypt.PasswordStrength("<rin>24zD*~"))
	fmt.Println(crypt.PasswordStrength("<rin>24xX.vD"))
	fmt.Println(crypt.PasswordStrength("12345aAlowD"))
	fmt.Println(crypt.PasswordStrength("12345aAl.D"))
	fmt.Println(crypt.PasswordStrength("123eeeeddggDDll"))
	fmt.Println(crypt.PasswordStrength("1234567890a"))
	fmt.Println(crypt.PasswordStrength("12345678.a"))
	fmt.Println(crypt.PasswordStrength("123456D.a"))
	fmt.Println(crypt.PasswordStrength("12345678D"))
	fmt.Println(crypt.PasswordStrength("123456.D"))
	fmt.Println(crypt.PasswordStrength("1234s.D"))
}
