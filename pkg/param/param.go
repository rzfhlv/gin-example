package param

var (
	DEFAULTLIMIT = 10
	DEFAULTPAGE  = 1
)

type Param struct {
	Page   int   `json:"page" form:"page"`
	Limit  int   `json:"limit" form:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func (f *Param) CalculateOffset() int {
	return f.Limit * (f.Page - 1)
}
