package cli

import (
	"Pair_Project_Clothing_Ecommerce/internal/domain"
	"fmt"
	"strconv"
	"strings"
)

func (c *App) handleCheckout() {
	fmt.Println("\n--- PROSES CHECKOUT ---")
	var items []domain.OrderDetail
	for {
		fmt.Print("Masukkan ID Produk (Ketik '0' jika selesai memilih): ")
		c.scanner.Scan()
		id, _ := strconv.Atoi(strings.TrimSpace(c.scanner.Text()))
		if id == 0 {
			break
		}

		fmt.Print("Jumlah Kuantitas: ")
		c.scanner.Scan()
		qty, _ := strconv.Atoi(strings.TrimSpace(c.scanner.Text()))
		items = append(items, domain.OrderDetail{ProductID: id, Quantity: qty})
	}
	if len(items) > 0 {
		err := c.orderUC.Checkout(c.user.ID, items, "transfer_bank")
		if err != nil {
			fmt.Println("❌ Checkout Gagal:", err.Error())
		} else {
			fmt.Println("🎉 Berhasil! Pesanan dan pemotongan stok sukses dilakukan.")
		}
	}
}
