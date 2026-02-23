package dto

type FilterInvoiceDto struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Query   string `json:"query"`
}
