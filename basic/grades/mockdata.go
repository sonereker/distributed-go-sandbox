package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Soner",
			LastName:  "Eker",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 50,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeHomework,
					Score: 100,
				},
				Grade{
					Title: "Quiz 3",
					Type:  GradeTest,
					Score: 60,
				},
			},
		},
	}
}
