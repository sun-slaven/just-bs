package college

type CollegeVo struct {
	Id        string  `json:"id"`
	Name      string        `json:"name"`
	MajorList []MajorVo `json:"major_list"`
}


type MajorVo struct {
	Id   string        `json:"id"`
	Name string        `json:"name"`
}
