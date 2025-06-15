package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	problemsScore := s.scoreForOneTask * float64(s.solvedProblems)
	return problemsScore > s.passingScore
}
