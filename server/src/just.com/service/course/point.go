package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
)

func (self *CourseService) AddPoint(point int64, courseId, userId string) error {
	// is mark?
	markTable := new(table.CourseMarkTable)
	markTable.CourseId = courseId
	markTable.UserId = userId
	count, countErr := self.Session.Count(markTable)
	if count == 0 {
		if countErr != nil {
			self.Log.Println(countErr)
		}
		return COURSE_POINT_ADD_ERR
	}
	// 删除所有之前的
	pointTable := new(table.CoursePointTable)
	pointTable.CourseId = courseId
	pointTable.UserID = userId
	_, deleteErr := self.Session.Delete(pointTable)
	if deleteErr != nil {
		self.Log.Println(deleteErr)
		return COURSE_POINT_ADD_ERR
	}
	pointTable.UUID = uuid.New()
	pointTable.Point = point
	insertNum, insertErr := self.Session.InsertOne(pointTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return COURSE_POINT_ADD_ERR
	}
	return nil
}

func (self *CourseService) UpdatePoint(point int64, courseId, userId string) error {
	// is mark?
	markTable := new(table.CourseMarkTable)
	markTable.CourseId = courseId
	markTable.UserId = userId
	count, countErr := self.Session.Count(markTable)
	if count == 0 {
		if countErr != nil {
			self.Log.Println(countErr)
		}
		return COURSE_POINT_UPDATE_ERR
	}
	// update
	pointTable := new(table.CoursePointTable)
	pointTable.CourseId = courseId
	pointTable.UserID = userId
	newOne := new(table.CoursePointTable)
	newOne.Point = point
	newOne.UpdateTime = time.Now()
	updateNum, updateErr := self.Session.Update(newOne, pointTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
			return COURSE_POINT_UPDATE_ERR
		}
		// 没有就insert
		newOne.UUID = uuid.New()
		newOne.CourseId = courseId
		newOne.UserID = userId
		newOne.CreateTime = time.Now()
		newOne.UpdateTime = time.Now()
		insertNum, insertErr := self.Session.InsertOne(newOne)
		if insertNum == 0 {
			if insertErr != nil {
				self.Log.Println(insertErr)
			}
			return COURSE_POINT_UPDATE_ERR
		}
	}
	return nil
}
