package cli

import (
	"fmt"
	"strings"
)

func (c *App) loginScreen() {
	fmt.Println("\n[ Wajib Login Untuk Melanjutkan ]")
	fmt.Print("Email: ")
	c.scanner.Scan()
	email := strings.TrimSpace(c.scanner.Text())

	fmt.Print("Password: ")
	c.scanner.Scan()
	pass := strings.TrimSpace(c.scanner.Text())

	user, err := c.authUC.Login(email, pass)
	if err != nil {
		fmt.Println("❌ Gagal Login:", err.Error())
		return
	}
	c.user = user
	fmt.Printf("✅ Berhasil Login! Selamat datang, %s (%s)\n", user.Email, user.Role)
}
