package model

// '{"title": "some_title", "text": "some_text", "url": "some_url"}'

type Content struct {
	title string `json:"title"`
	text  string `json:"text"`
	url   string `json:"url"`
}

func (me *Content) GetTitle() string {
	return me.title
}

func (me *Content) GetText() string {
	return me.text
}

func (me *Content) GetUrl() string {
	return me.url
}

func GoGo() {}
