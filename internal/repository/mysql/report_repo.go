package mysql

import (
	"Pair_Project_Clothing_Ecommerce/internal/domain"
	"database/sql"
	"fmt"
)

type reportRepo struct{ db *sql.DB }

func NewReportRepository(db *sql.DB) domain.ReportRepository { return &reportRepo{db: db} }

func (r *reportRepo) GetUserReports() ([]domain.UserReportDTO, error) {
	query := `
        SELECT 
            u.id, 
            u.email, 
            u.role, 
            COALESCE(p.full_name, '-'), 
            COALESCE(p.phone, '-') 
        FROM users u 
        LEFT JOIN profiles p ON u.id = p.user_id 
        ORDER BY u.id ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.UserReportDTO
	for rows.Next() {
		var d domain.UserReportDTO

		// Tambahkan penanganan error di sini
		err := rows.Scan(&d.ID, &d.Email, &d.Role, &d.FullName, &d.Phone)
		if err != nil {
			fmt.Printf("DEBUG: Error saat scan data user: %v\n", err)
			continue // Lewati baris yang bermasalah agar aplikasi tidak crash
		}

		list = append(list, d)
	}
	return list, nil
}

func (r *reportRepo) GetStockReports() ([]domain.StockReportDTO, error) {
	// 1. Tambahkan COALESCE(c.name, ...) di SELECT dan LEFT JOIN ke tabel categories
	query := `
        SELECT 
            COALESCE(c.name, 'Tanpa Kategori'), 
            p.sku, 
            p.name, 
            COALESCE(p.size, '-'), 
            COALESCE(p.color, '-'), 
            COALESCE(p.description, '-'), 
            p.stock,
            CASE 
                WHEN p.stock <= 0 THEN 'Habis'
                WHEN p.stock < 5 THEN 'Menipis'
                ELSE 'Aman'
            END AS status_stok
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id 
        ORDER BY p.id ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.StockReportDTO
	for rows.Next() {
		var item domain.StockReportDTO

		// 2. Masukkan &item.Category di urutan pertama (sesuai urutan SELECT di atas)
		err := rows.Scan(
			&item.Category,
			&item.SKU,
			&item.Name,
			&item.Size,
			&item.Color,
			&item.Description,
			&item.Stock,
			&item.StatusStok,
		)
		if err != nil {
			fmt.Printf("DEBUG: Error saat scan stok: %v\n", err)
			continue
		}

		list = append(list, item)
	}
	return list, nil
}

func (r *reportRepo) GetOrderReports() ([]domain.OrderReportDTO, error) {
	// 1. Tambahkan LEFT JOIN ke tabel payments dan ambil pay.payment_method
	query := `
        SELECT 
            o.id, 
            u.email, 
            o.total_amount, 
            o.status, 
            COALESCE(pay.payment_method, '-'), 
            DATE_FORMAT(o.order_date, '%Y-%m-%d %H:%i') 
        FROM orders o 
        JOIN users u ON o.user_id = u.id 
        LEFT JOIN payments pay ON o.id = pay.order_id
        ORDER BY o.order_date DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.OrderReportDTO
	for rows.Next() {
		var d domain.OrderReportDTO
		// 2. Tambahkan &d.PaymentMethod di urutan ke-5 sesuai SELECT di atas
		err := rows.Scan(&d.OrderID, &d.Pembeli, &d.TotalAmount, &d.Status, &d.PaymentMethod, &d.WaktuTransaksi)
		if err != nil {
			fmt.Printf("DEBUG: Error saat scan order: %v\n", err)
			continue
		}
		list = append(list, d)
	}
	return list, nil
}
