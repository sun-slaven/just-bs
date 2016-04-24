package college

type CollegeVo struct {
	Id        string
	Name      string
	MajorList []MajorVo `json:"major_list"`
}


type MajorVo struct {
	Id   string
	Name string
}
