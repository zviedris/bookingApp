package forms

type errors map[string][]string

//add an error message for a given field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

//returns firs error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]

}
