package dto

type CourseDto struct {
	Name         string
	Introduction string
	Syllabus     string
	Wish         string
	Experiment   string
	Major        string
	College      string
}

func NewCouseDto(name ,intro,syllabus,wish,experiment,major,college string) *CourseDto {
	courseDto := new(CourseDto)
	courseDto.Name = name
	courseDto.Introduction = intro
	courseDto.Syllabus = syllabus
	courseDto.Wish = wish
	courseDto.Experiment = experiment
	courseDto.Major = major
	courseDto.College = college
	return courseDto
}