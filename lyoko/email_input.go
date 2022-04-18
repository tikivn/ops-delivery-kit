package lyoko

type EmailInput struct {
	EmailTemplate string   `json:"email_template"`
	Type          string   `json:"type"`
	BoxCode       string   `json:"box_code"`
	Receiver      string   `json:"receiver"`
	FileURLs      []string `json:"file_urls"`
	Content       string   `json:"content"`
}
