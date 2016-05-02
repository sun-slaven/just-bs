package dto


type CourseAddRequest struct {
	Id           string `json:"id"`
	Name         string `json:"name"`       // required
	Description  string  `json:"description"`
	Introduction string `json:"introduction"`
	Experiment   string `json:"experiment"`
	Syllabus     string `json:"syllabus"`   // required
	Wish         string `json:"wish"`
	MajorId      string `json:"major_id"`   // required
	CollegeId    string `json:"college_id"` // required
	TeacherId    string `json:"teacher_id"` // required
	IconUrl      string `json:"icon_url"`   // required
}