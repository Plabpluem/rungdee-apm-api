package pkg

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetPerPage() int {
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	return p.PerPage
}

func (p *Pagination) GetOffSet() int {
	return (p.GetPage() - 1) * p.GetPerPage()
}
