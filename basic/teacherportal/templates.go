package teacherportal

import "html/template"

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"basic/teacherportal/students.gohtml",
		"basic/teacherportal/student.gohtml",
		"basic/teacherportal/grades.gohtml")
	if err != nil {
		return err
	}
	return nil
}
