package service
import (
	"just.com/service/course"
	"just.com/test"
	"just.com/dto"
	"testing"
	"time"
)

func TestCourseAdd(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()
	dto := dto.CourseDto{}
	dto.Name = "NAME"
	dto.College = "COLLEGE"
	dto.Experiment = "EXPERIMENT"
	dto.Introduction = "INTRODUCTION"
	dto.Major = "MAJOR"
	dto.Plan = "PLAN"
	dto.Syllabus = "SYLLABUS"
	_, addErr := cs.Add(dto, "aa5eba0a-703c-4801-955b-1f44997738fe")
	if addErr != nil {
		Logger.Println(addErr)
		cs.Session.Rollback()
		t.Fail()
	}
}

func TestCourseMark(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()

	cs.Session.Begin()

	markErr := cs.Mark("35bfcb48-9f2a-4938-bb54-cab72e951970", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if markErr != nil {
		Logger.Println(markErr)
		cs.Session.Rollback()
		t.Fail()
	}
	cs.Session.Commit()
}

func TestCourseMarkCancel(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()

	markErr := cs.MarkCancel("35bfcb48-9f2a-4938-bb54-cab72e951970", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if markErr != nil {
		Logger.Println(markErr)
		t.Fail()
	}
}

func TestCommentAdd(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()

	commentId, addErr := cs.AddComment("2333", "35bfcb48-9f2a-4938-bb54-cab72e951970", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if addErr != nil {
		t.Fail()
	}
	Logger.Println(commentId)
	time.Sleep(time.Second * 5)
}

func TestPointAdd(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()

	addErr := cs.AddPoint(20, "35bfcb48-9f2a-4938-bb54-cab72e951970", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if addErr != nil {
		Logger.Println(addErr)
		t.Fail()
	}
}

func TestPointUpdate(t *testing.T) {
	cs := service.CourseService{}
	cs.Log = Logger
	cs.Session = test.DataSource.NewSession()
	defer cs.Session.Close()

	updateErr := cs.UpdatePoint(100, "35bfcb48-9f2a-4938-bb54-cab72e951970", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if updateErr != nil {
		Logger.Println(updateErr)
		t.Fail()
	}
}