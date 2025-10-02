package esepunittests

//Changed to have 1 list
type GradeCalculator struct {
	allAssignments []Grade
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

//One list named allAssignments holds the different types
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		allAssignments: make([]Grade, 0),
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

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	//Each have the same 3 needs. They are made more generic
	gc.allAssignments = append(gc.allAssignments, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	//Made the specific assignment in allAssignment list
	assignment_average := computeAverage(gc.allAssignments, Assignment)
	exam_average := computeAverage(gc.allAssignments, Exam)
	essay_average := computeAverage(gc.allAssignments, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, t GradeType) int {
	sum := 0
	count := 0

	//Looks through all the grades
	for _, grade := range grades {
		if grade.Type == t {
			sum += grade.Grade
			count++
		}
	}

	return sum / count
}
