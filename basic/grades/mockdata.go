package grades

import "log"

func init() {
	log.Println("Setting students")
	students = []Student{
		{
			ID:        1,
			FirstName: "Soner",
			LastName:  "Eker",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 50,
				},
				{
					Title: "Quiz 2",
					Type:  GradeHomework,
					Score: 100,
				},
				{
					Title: "Quiz 3",
					Type:  GradeTest,
					Score: 60,
				},
			},
		},
	}
}
