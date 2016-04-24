package college
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
)

func LoadCollegeVoList(session *xorm.Session, log *log.Logger) (collegeVoList []CollegeVo) {
	// vo
	collegeList := make([]table.CollegeTable, 0)
	// table
	collegeVoList = make([]CollegeVo, 0)
	findErr := session.Find(&collegeList, &table.CollegeTable{FrozenStatus:"N"})
	if findErr != nil {
		log.Println(findErr)
		return
	}
	// range college
	for _, college := range collegeList {
		majorList := make([]table.MajorTable, 0)
		majorVoList := make([]MajorVo, 0)
		findErr := session.Find(&majorList, &table.MajorTable{CollegeId:college.UUID, FrozenStatus:"N"})
		if findErr != nil {
			log.Println(findErr)
			return
		}
		// range major
		for _, major := range majorList {
			majorVoList = append(majorVoList, MajorVo{Id:major.UUID, Name:major.Name})
		}
		collegeVoList = append(collegeVoList, CollegeVo{Id:college.UUID, Name:college.Name, MajorList:majorVoList})
	}
	return
}
