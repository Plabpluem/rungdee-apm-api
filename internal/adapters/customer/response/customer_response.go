package response

import "rungdee-apm-api/internal/entities"

type CustomerPaginatedResponse struct {
	Data       []*entities.Customer `json:"data"`
	Total      int64                `json:"total"`
	Page       int                  `json:"page"`
	PerPage    int                  `json:"per_page"`
	TotalPages int                  `json:"total_pages"`
}
