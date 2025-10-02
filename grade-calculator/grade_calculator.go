package esepunittests

//Added pass/fail bool
type GradeCalculator struct {
	allAssignments []Grade
	passOrFail bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

//Initialized with pass/fail bool
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		allAssignments: make([]Grade, 0),
		passOrFail: false,
	}
}


func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

//Created a new final calculator for pass/fail ranking
func (gc *GradeCalculator) GetFinalRanking() string {
	numericalGrade := gc.calculateNumericalGrade()

	//Returns true if passing grade
	if numericalGrade >= 70 {
		return "Pass"
	}
	return "Fail"
}



func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.allAssignments = append(gc.allAssignments, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.allAssignments, Assignment)
	exam_average := computeAverage(gc.allAssignments, Exam)
	essay_average := computeAverage(gc.allAssignments, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, t GradeType) int {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade.Type == t {
			sum += grade.Grade
			count++
		}
	}

	return sum / count
}
