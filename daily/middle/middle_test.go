package middle

import (
	"fmt"
	"testing"
)

func TestTopStudents2512(t *testing.T) {
	cases := []struct {
		positive_feedback []string
		negative_feedback []string
		report            []string
		student_id        []int
		k                 int
	}{
		{
			[]string{"smart", "brilliant", "studious"},
			[]string{"not"},
			[]string{"studious", "the student is smart"},
			[]int{1, 2},
			2,
		},
	}

	for _, cas := range cases {
		fmt.Println(TopStudents2512(cas.positive_feedback, cas.negative_feedback, cas.report, cas.student_id, cas.k))
	}
}
