package college
import "just.com/model/db/table"

func LoadCollegeVo(college *table.CollegeTable) *CollegeVo {
	return &CollegeVo{
		Id:college.UUID,
		Name:college.Name,
	}
}


func LoadMajorVo(major *table.MajorTable) *MajorVo {
	return &MajorVo{
		Id:major.UUID,
		Name:major.Name,
	}
}
