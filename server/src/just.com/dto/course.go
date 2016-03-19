package dto

type CourseDto struct {
	Name         string
	Introduction string
	Syllabus     string
	Plan         string
	Experiment   string
	Major        string
	College      string
}

func NewCouseDto(name ,intro,syllabus,plan,experiment,major,college string) *CourseDto {
	courseDto := new(CourseDto)
	courseDto.Name = name
	courseDto.Introduction = intro
	courseDto.Syllabus = syllabus
	courseDto.Plan = plan
	courseDto.Experiment = experiment
	courseDto.Major = major
	courseDto.College = college
	return courseDto
}