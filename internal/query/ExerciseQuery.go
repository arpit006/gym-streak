package query

type exerciseQueries struct{}

var exerciseQuery *exerciseQueries = nil

func GetExerciseQueryInstance() *exerciseQueries {
	if exerciseQuery != nil {
		return exerciseQuery
	}
	exerciseQuery = &exerciseQueries{}
	return exerciseQuery
}

func (eQuery exerciseQueries) GetCreateExerciseQuery() string {
	return "INSERT INTO exercise(id, name, meta, category, type) VALUES (?, ?, ?, ?, ?)"
}
