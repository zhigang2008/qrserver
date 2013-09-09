package util

type Pagination struct {
	Count       int
	PageSize    int
	CurrentPage int
	NavPages    []int
	PageCount   int
	QueryParams map[string]string
	Data        interface{}
}

//计算分页数据
func (p *Pagination) Compute() {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}

	p.PageCount = p.Count / p.PageSize
	if p.Count%p.PageSize > 0 {
		p.PageCount += 1
	}

	//计算分页展示页码
	p.NavPages = []int{}
	if p.PageCount <= 10 {
		for i := 1; i <= p.PageCount; i++ {
			p.NavPages = append(p.NavPages, i)
		}
	} else {
		if p.CurrentPage-5 > 0 {
			for pa := p.CurrentPage - 5; pa < p.CurrentPage; pa++ {
				p.NavPages = append(p.NavPages, pa)
			}
		} else {
			for i := 1; i < p.CurrentPage; i++ {
				p.NavPages = append(p.NavPages, i)
			}
		}

		if p.CurrentPage+5 <= p.PageCount {
			for i := p.CurrentPage; i < p.CurrentPage+5; i++ {
				p.NavPages = append(p.NavPages, i)
			}
		} else {
			for i := p.CurrentPage; i < p.PageCount; i++ {
				p.NavPages = append(p.NavPages, i)
			}
		}
	}

}

//添加的查询参数
func (p *Pagination) AddParams(k string, v string) {
	if p.QueryParams == nil {
		p.QueryParams = make(map[string]string)
	}
	p.QueryParams[k] = v
}

//供数据库分页,跳过的数据量
func (p *Pagination) SkipNum() int {
	return p.PageSize * (p.CurrentPage - 1)
}
