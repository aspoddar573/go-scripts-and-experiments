package templateTypes

type CompetencyDBDocument struct {
	Competency string `json:"competency,omitempty"`
}

type SetupGuideline struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
}

type FrequentlyAskedQuestion struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

type TemplateSupportingDocument struct {
	Guide string `json:"guide,omitempty"`
}

type ProgramTemplateRow struct {
	TemplateId                               int64                         `json:"template_id,omitempty,string"`
	CompanyId                                int64                         `json:"company_id,omitempty,string"`
	TemplateSeriesId                         int64                         `json:"template_series_id,omitempty,string"`
	Name                                     string                        `json:"name,omitempty"`
	Purpose                                  string                        `json:"purpose,omitempty"`
	Description                              string                        `json:"description,omitempty"`
	PostCreationTitle                        string                        `json:"postCreationTitle,omitempty"`
	PostCreationDescription                  string                        `json:"postCreationDescription,omitempty"`
	State                                    string                        `json:"state,omitempty"`
	CreatedBy                                int64                         `json:"created_by,omitempty,string"`
	UpdatedBy                                int64                         `json:"updated_by,omitempty,string"`
	EstimatedCompletionTimeInSeconds         int64                         `json:"estimatedCompletionInSeconds,omitempty,string"`
	EstimatedCompletionTimeIntervalInSeconds int64                         `json:"estimatedCompletionTimeIntervalInSeconds,omitempty,string"`
	Scope                                    string                        `json:"scope,omitempty"`
	CreatorId                                int64                         `json:"creator_id,omitempty,string"`
	Thumbnail                                *Thumbnail                    `json:"thumbnail,omitempty"`
	ListingThumbnail                         *Thumbnail                    `json:"listing_thumbnail,omitempty"`
	Competencies                             []*CompetencyDBDocument       `json:"competencies,omitempty"`
	SetupGuidelines                          []*SetupGuideline             `json:"setupGuidelines,omitempty"`
	FrequentlyAskedQuestions                 []*FrequentlyAskedQuestion    `json:"frequentlyAskedQuestions,omitempty"`
	TemplateSupportingDocuments              []*TemplateSupportingDocument `json:"templateSupportingDocuments,omitempty"`
}

type ProgramTemplateDbRow struct {
	TemplateId                               int64  `json:"template_id,omitempty,string"`
	CompanyId                                int64  `json:"company_id,omitempty,string"`
	TemplateSeriesId                         int64  `json:"template_series_id,omitempty,string"`
	Name                                     string `json:"name,omitempty"`
	Purpose                                  string `json:"purpose,omitempty"`
	Description                              string `json:"description,omitempty"`
	PostCreationTitle                        string `json:"postCreationTitle,omitempty"`
	PostCreationDescription                  string `json:"postCreationDescription,omitempty"`
	Thumbnail                                string `json:"thumbnail,omitempty"`
	ListingThumbnail                         string `json:"listing_thumbnail,omitempty"`
	Competencies                             string `json:"competencies,omitempty"`
	TemplateSupportingDocuments              string `json:"templateSupportingDocuments,omitempty"`
	FrequentlyAskedQuestions                 string `json:"frequentlyAskedQuestions,omitempty"`
	SetupGuidelines                          string `json:"setupGuidelines,omitempty"`
	State                                    string `json:"state,omitempty"`
	CreatedBy                                int64  `json:"created_by,omitempty,string"`
	UpdatedBy                                int64  `json:"updated_by,omitempty,string"`
	EstimatedCompletionTimeInSeconds         int64  `json:"estimatedCompletionInSeconds,omitempty,string"`
	EstimatedCompletionTimeIntervalInSeconds int64  `json:"estimatedCompletionTimeIntervalInSeconds,omitempty,string"`
	Scope                                    string `json:"scope,omitempty"`
	CreatorId                                int64  `json:"creator_id,omitempty,string"`
}

type Thumbnail struct {
	OriginalUrl          string `json:"original_url,omitempty"`
	ProcessedUrl         string `json:"processed_url,omitempty"`
	ProcessedUrl_180X120 string `json:"processed_url_180x120,omitempty"`
	ProcessedUrl_600X360 string `json:"processed_url_600x360,omitempty"`
	MediaType            string `json:"media_type,omitempty"`
}

type ProgramTemplateCreatorMetadata struct {
	CompanyLogo        Thumbnail `json:"company_logo,omitempty"`
	CompanyListingLogo Thumbnail `json:"company_listing_logo,omitempty"`
	Email              string    `json:"email,omitempty"`
	Website            string    `json:"website,omitempty"`
}

type ProgramTemplateCreatorDbRow struct {
	CreatorId   int64  `json:"creator_id,omitempty,string"`
	Name        string `json:"creator_name,omitempty"`
	Type        string `json:"creator_type,omitempty"`
	Description string `json:"description,omitempty"`
	Metadata    string `json:"metadata,omitempty"`
}

type ProgramTemplateCreatorRow struct {
	CreatorId   int64                          `json:"creator_id,omitempty,string"`
	Name        string                         `json:"creator_name,omitempty"`
	Type        string                         `json:"creator_type,omitempty"`
	Description string                         `json:"description,omitempty"`
	Metadata    ProgramTemplateCreatorMetadata `json:"metadata,omitempty"`
}
