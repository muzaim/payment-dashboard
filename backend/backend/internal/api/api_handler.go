package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth *ah.AuthHandler
	DB   *sql.DB
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(
	w http.ResponseWriter,
	r *http.Request,
	params openapigen.GetDashboardV1PaymentsParams,
) {

	queryBase := "FROM payments WHERE 1=1"
	args := []interface{}{}

	if params.Status != nil && *params.Status != "" {
		queryBase += " AND status = ?"
		args = append(args, *params.Status)
	}

	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	offset := (page - 1) * limit

	var total int
	err := h.DB.QueryRow("SELECT COUNT(*) "+queryBase, args...).Scan(&total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderClause := ""

	if params.Sort != nil && *params.Sort != "" {
		sortField := *params.Sort

		switch sortField {
		case "latest":
			orderClause = " ORDER BY created_at DESC"
		case "earliest":
			orderClause = " ORDER BY created_at ASC"
		default:
			order := "ASC"

			if strings.HasPrefix(sortField, "-") {
				sortField = strings.TrimPrefix(sortField, "-")
				order = "DESC"
			}

			switch sortField {
			case "created_at", "amount", "payment_id":
				orderClause = " ORDER BY " + sortField + " " + order
			}
		}
	}

	finalQuery := "SELECT id, payment_id, merchant, status, amount, created_at " +
		queryBase +
		orderClause +
		" LIMIT ? OFFSET ?"

	args = append(args, limit, offset)

	rows, err := h.DB.Query(finalQuery, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	payments := []openapigen.Payment{}

	for rows.Next() {
		var p openapigen.Payment
		err := rows.Scan(
			&p.Id,
			&p.PaymentId,
			&p.Merchant,
			&p.Status,
			&p.Amount,
			&p.CreatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payments = append(payments, p)
	}

	resp := map[string]interface{}{
		"payments": payments,
		"meta": map[string]interface{}{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + limit - 1) / limit,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *APIHandler) GetDashboardV1PaymentsSummary(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.DB.Query(`
		SELECT status, COUNT(*) as total
		FROM payments
		GROUP BY status
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	summary := map[string]int{
		"pending":    0,
		"processing": 0,
		"completed":  0,
		"failed":     0,
	}

	for rows.Next() {
		var status string
		var total int

		if err := rows.Scan(&status, &total); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		summary[status] = total
	}

	resp := map[string]interface{}{
		"summary": summary,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *APIHandler) PutDashboardV1PaymentIdReview(
	w http.ResponseWriter,
	r *http.Request,
	id string,
) {

	_, err := h.DB.Exec(`
		UPDATE payments
		SET status='completed'
		WHERE payment_id=?
	`, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "success",
	})
}

func (h *APIHandler) GetDashboardV1PaymentsTopMerchants(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows, err := h.DB.Query(`
		SELECT merchant, COUNT(*) as total
		FROM payments
		GROUP BY merchant
		ORDER BY total DESC
		LIMIT 3
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type MerchantStat struct {
		Merchant string `json:"merchant"`
		Total    int    `json:"total"`
	}

	stats := []MerchantStat{}

	for rows.Next() {
		var m MerchantStat
		if err := rows.Scan(&m.Merchant, &m.Total); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		stats = append(stats, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": stats,
	})
}
