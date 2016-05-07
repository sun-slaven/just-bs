package service
import (
	"just.com/query/vo/course"
	"just.com/err"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"just.com/dto"
	"just.com/common"
	"time"
)

func (self *CourseService) AddChapter(courseId, userId string, request *dto.CourseChapterRequest) (chapterVo *course.CourseChapterVo, error *err.HttpError) {
	if common.IsEmpty(request.Name, request.Content) {
		error = err.NO_REQUIRED_PARAM_FOUND
		return
	}
	courseTable, getErr := self.GetCourse(courseId)
	if getErr != nil {
		error = getErr
		return
	}
	table := &table.CourseChapterTable{
		UUID:uuid.New(),
		Name:request.Name,
		Content:request.Content,
		CourseId:courseTable.UUID,
		CreateUser:userId,
		CreateTime:time.Now(),
		UpdateTime:time.Now(),
		UpdateUser:userId,
		FrozenStatus:"N",
	}
	insertNum, insertErr := self.Session.Insert(table)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		error = err.CHAPTER_ADD_ERR
		return
	}
	return course.NewChapterVo(table), nil
}

/**
更新章节,没有required
 */
func (self *CourseService) UpdateChapter(chapterId, userId string, request *dto.CourseChapterRequest) (chapterVo *course.CourseChapterVo, error *err.HttpError) {
	if common.IsEmpty(chapterId) {
		error = err.NO_REQUIRED_PARAM_FOUND
		return
	}
	chapter := new(table.CourseChapterTable)
	getFlag, getErr := self.Session.Id(chapterId).Get(chapter)
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.NO_CHAPTER_FOUND
		return
	}
	if request.Name != "" {
		chapter.Name = request.Name
	}
	if request.Content != "" {
		chapter.Content = request.Content
	}
	if request.Order != 0 {
		chapter.Order = request.Order
	}
	chapter.UpdateTime = time.Now()
	chapter.UpdateUser = userId
	updateNum, updateErr := self.Session.Id(chapterId).Update(chapter)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		error = err.CHAPTER_UPDATE_ERR
		return
	}
	return course.NewChapterVo(chapter), nil
}

func (self *CourseService) DeleteChapter(courseId, chapterId, userId string) (error *err.HttpError) {
	if common.IsEmpty(courseId, chapterId) {
		return err.NO_REQUIRED_PARAM_FOUND
	}
	chapterTable := new(table.CourseChapterTable)
	chapterTable.FrozenStatus = "Y"
	chapterTable.FrozenTime = time.Now()
	chapterTable.UpdateUser = userId
	chapterTable.UpdateTime = time.Now()
	updateNum, updateErr := self.Session.Update(chapterTable, &table.CourseChapterTable{UUID:chapterId, CourseId:courseId})
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return err.NO_CHAPTER_FOUND
	}
	return nil
}