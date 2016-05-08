package service
import (
	"just.com/model/db/table"
	"time"
	"just.com/model/db"
	"log"
	"just.com/value"
)

/*flush course mark sum
	when mark or mark cancel trigger
*/
func FlushMarkSum(courseId string, ds *db.DataSource, log *log.Logger) error {
	session := ds.NewSession()
	defer session.Close()
	countSql := `SELECT COUNT("UUID") FROM "COURSE_MARK" WHERE "COURSE_ID" = ?`
	count, countErr := session.Sql(countSql, courseId).Count(&table.CourseMarkTable{})
	if countErr != nil {
		log.Println(countErr)
	}
	courseTable := table.CourseTable{}
	courseTable.MarkSum = count
	updateNum, updateErr := session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		if updateErr != nil {
			log.Println(updateErr)
		}
		return COURSE_FLUSH_MARK_NUM_ERR
	}
	return nil
}

/*flush course's comment sum
	when comment or frozen comment trigger
*/
func FlushCommentSum(courseId string, ds *db.DataSource, log *log.Logger) error {
	session := ds.NewSession()
	defer session.Close()
	comment := new(table.CourseCommentTable)
	countSql := `SELECT COUNT("UUID") FROM "COURSE_COMMENT" WHERE "COURSE_ID" = ?
		AND "FROZEN_STATUS" = ?`
	count, countErr := session.Sql(countSql, courseId, value.STATUS_ENABLED).Count(comment)
	if countErr != nil {
		log.Println(count)
	}
	course := new(table.CourseTable)
	course.CommentSum = count
	updateNum, updateErr := session.Id(courseId).Update(course)
	if updateNum == 0 {
		if updateErr != nil {
			log.Println(updateNum)
		}
		return COURSE_FLUSH_COMMENT_NUM_ERR
	}
	return nil
}

/*flush point*/
func FlushPoint(courseId string, ds *db.DataSource, log *log.Logger) error {
	session := ds.NewSession()
	defer session.Close()
	pointList := make([]table.CoursePointTable, 0)
	oldPoint := new(table.CoursePointTable)
	oldPoint.CourseId = courseId
	findErr := session.Find(&pointList, oldPoint)
	if findErr != nil {
		log.Println(findErr)
		return COURSE_FLUSH_POINT_ERR
	}
	// 重新计算评分
	pointPerson := len(pointList)
	var pointSum int64
	for i := 0; i < pointPerson; i++ {
		pointSum += pointList[i].Point
	}
	newCourse := new(table.CourseTable)
	newCourse.UpdateTime = time.Now()
	newCourse.PointPerson = int64(pointPerson)
	newCourse.Points = pointSum
	updateNum, updateErr := session.Id(courseId).Update(newCourse)
	if updateNum == 0 {
		if updateErr != nil {
			log.Println(updateErr)
		}
		return COURSE_FLUSH_POINT_ERR
	}
	return nil
}