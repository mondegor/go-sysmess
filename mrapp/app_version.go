package mrapp

import (
	"context"
	"os/exec"
	"strings"
)

// Version возвращает версию приложения из системы контроля версий Git.
// Если текущая ветка отличается от master, main или HEAD, возвращается название ветки.
// В противном случае возвращается результат git describe с информацией о теге, коммите и состоянии изменений.
// Если Git недоступен или произошла ошибка, возвращается пустая строка или название ветки.
func Version(ctx context.Context) string {
	if _, err := exec.LookPath("git"); err != nil {
		return ""
	}

	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--abbrev-ref", "HEAD")

	b, err := cmd.Output()
	if err != nil {
		return ""
	}

	value := strings.TrimSpace(string(b))

	// если указана любая ветка кроме мастера
	if value != "master" && value != "main" && value != "HEAD" {
		return value
	}

	// Примеры тегов:
	//   v0.14.7-0-de1493e0
	//   v0.8.1-0-gd3a5efc-dirty
	cmd = exec.CommandContext(ctx, "git", "describe", "--long", "--always", "--dirty")

	b, err = cmd.Output()
	if err != nil {
		return value
	}

	return strings.TrimSpace(string(b))
}
