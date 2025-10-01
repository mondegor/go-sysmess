package main

import (
	"fmt"
	"os"

	"github.com/mondegor/go-sysmess/mrlib/crypt"
	"github.com/mondegor/go-sysmess/mrlib/crypt/password"
	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrlog/slog"
)

func main() {
	logger, _ := slog.NewLoggerAdapter(slog.WithWriter(os.Stdout))

	value, _ := crypt.GenerateDigits(16)
	mrlog.Debug(logger, "GenerateDigits", "value", value)

	value, _ = crypt.GenerateHex(16)
	mrlog.Info(logger, "GenTokenHex", "value", value)

	value, _ = crypt.GenerateToken(64)
	mrlog.Info(logger, "GenerateToken", "value", value)

	valueBytes, _ := crypt.GenerateBytes([]byte("abc123.,:"), 16)
	mrlog.Info(logger, "GenerateBytes", "value", string(valueBytes))

	pwgen := password.NewGenerator()

	mrlog.Info(logger, "GenPassword", "password", pwgen.Generate(16, password.CharAll))

	pw := pwgen.Generate(12, password.CharAbc)
	mrlog.Info(logger, "PasswordStrength 12 abc", "password", pw, "strength", password.CalcStrength(pw))

	pw = pwgen.Generate(9, password.CharAbcNumerals)
	mrlog.Info(logger, "PasswordStrength 9 abc+num", "password", pw, "strength", password.CalcStrength(pw))

	pw = pwgen.Generate(12, password.CharAll)
	mrlog.Info(logger, "PasswordStrength 12 all", "password", pw, "strength", password.CalcStrength(pw))

	fmt.Println(password.CalcStrength("<rin>24zD*~"))
	fmt.Println(password.CalcStrength("<rin>24xX.vD"))
	fmt.Println(password.CalcStrength("12345aAlowD"))
	fmt.Println(password.CalcStrength("12345aAl.D"))
	fmt.Println(password.CalcStrength("123eeeeddggDDll"))
	fmt.Println(password.CalcStrength("1234567890a"))
	fmt.Println(password.CalcStrength("12345678.a"))
	fmt.Println(password.CalcStrength("123456D.a"))
	fmt.Println(password.CalcStrength("12345678D"))
	fmt.Println(password.CalcStrength("123456.D"))
	fmt.Println(password.CalcStrength("1234s.D"))
}
