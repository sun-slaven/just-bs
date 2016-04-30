package college
import "just.com/model/db/table"

type CollegeVo struct {
	Id        string  `json:"id"`
	Name      string        `json:"name"`
	MajorList []MajorVo `json:"major_list"`
}


type MajorVo struct {
	Id   string        `json:"id"`
	Name string        `json:"name"`
}


func NewCollegeVo(college *table.CollegeTable) *CollegeVo {
	return &CollegeVo{
		Id:college.UUID,
		Name:college.Name,
	}
}

func NewMajorVo(major *table.MajorTable) *MajorVo {
	return &MajorVo{
		Id:major.UUID,
		Name:major.Name,
	}
}