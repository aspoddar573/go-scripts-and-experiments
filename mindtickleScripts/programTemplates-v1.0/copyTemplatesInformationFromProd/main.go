package main

import (
	"fmt"
	"goProject/mindtickleScripts/programTemplates-v1.0/programTemplates"
	"goProject/mindtickleScripts/programTemplates-v1.0/templateTypes"
)

const TARGET_ENV = "staging"

var TEMPLATES_TO_COPY = []string{"'1418224246089186514'"}
var TEMPLATE_CREATORS_TO_COPY = []string{"'1650262556293827453'"}
var additionalTemplatesCopyInformation = map[string]templateTypes.AdditionalProgramTemplateCopyInformation{
	"1418224246089186514": {
		SeriesId: 1432270446000912962,
	},
	"1418224246089186515": {
		SeriesId: 1432270446000912962,
	},
	"1418224246089186519": {
		SeriesId: 1432270446000912962,
	},
	"1418224246089186520": {
		SeriesId: 1432270446000912962,
	},
	"1418224246089186516": {
		SeriesId: 1432270446000912962,
	},
}
var additionalTemplateCreatorsCopyInformation = map[string]templateTypes.AdditionalTemplateCreatorCopyInformation{}

func main() {
	fmt.Printf("Starting operation!\n")
	//templateCreators.CopyProgramTemplateCreators(TARGET_ENV, TEMPLATE_CREATORS_TO_COPY, additionalTemplateCreatorsCopyInformation)
	programTemplates.CopyProgramTemplates(TARGET_ENV, TEMPLATES_TO_COPY, additionalTemplatesCopyInformation)
	fmt.Printf("\nCompleted operation!\n")
}
