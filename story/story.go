package story

type ArcMap = map[string]Arc

type (
	Arc struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []Option `json:"options"`
	}

	Option struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	}
)
