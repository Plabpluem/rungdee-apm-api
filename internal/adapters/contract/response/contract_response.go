package response

import "rungdee-apm-api/internal/entities"

type ContractPaginatedResponse struct {
	Data       []*entities.Contract `json:"data"`
	Total      int64                `json:"total"`
	Page       int                  `json:"page"`
	PerPage    int                  `json:"per_page"`
	TotalPages int                  `json:"total_pages"`
}
