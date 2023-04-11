package models

type Pagination struct {
	Total int64 `json:"total,omitempty"`
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
}

type ApiResponseOverview struct {
	Network    string   `json:"network"`
	LegendData []string `json:"legendData"` // 种类
	XaxisData  []string `json:"xaxisData"`  // 时间
	SeriesData []Serie  `json:"seriesData"` // 数据
}

type Serie struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}
