package dto


type CourseAddRequest struct {
	Name         string `json:"name"`       // required
	Description  string  `json:"description"`
	Introduction string `json:"introduction"`
	Experiment   string `json:"experiment"`
	Wish         string `json:"wish"`
	MajorId      string `json:"major_id"`   // required
	CollegeId    string `json:"college_id"` // required
	TeacherId    string `json:"teacher_id"` // required
	ImageUrl     string `json:"image_url"`  // required
}