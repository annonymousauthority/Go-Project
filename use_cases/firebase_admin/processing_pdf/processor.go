package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dslipak/pdf"
)

type Processor interface {
	evaluateResume() (string, int)
}

type myResume struct {
	res string
}

func main() {
	content, err := readPdf("../data/augustine_resume_doc.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)
	var processor Processor
	processor = myResume{content}
	fmt.Println(processor.evaluateResume())

}

func readPdf(path string) (string, error) {
	f, err := pdf.Open(path)
	// remember close file
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := f.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func (r myResume) evaluateResume() (string, int) {
	/*
		The algorithm is meant to find the consistency of the resume against the requirements of the job description.
		The algorithm ranks the resume against the following parameters;

		- Total years as a professional
		- Ratio of needed skills to mentioned skills.
		- Position compatibility.
		- Availability of custom keywords to grammar in the resume.
		{Managed, APi, REST, Mentored, Trained.}

		Criteria Checks:
		What is the Big-O notation of the algorithm?

		Considerations:

		Can we format the resume correctly? Into every header.


		Step 1: Convert the resume to a []String
		Step 2: Sort the criteria of the resume storing each score

	*/

	resSection := strings.Split(r.res, " ")
	fmt.Println(resSection)
	return "", 0
}

/*
 * Copyright (C) Augustine Francis, 2023
 */
