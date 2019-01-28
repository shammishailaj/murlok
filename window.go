package murlok

// Window is a struct that describes a window.
type Window struct {
	// The url load when the view is created.
	URL string

	// The background color (#rrggbb).
	BackgroundColor string `json:",omitempty"`

	// Sets the background as a frosted surface.
	FrostedBackground bool `json:",omitempty"`

	// Sets a title bar color. When set, the web content appears below the title
	// bar rather than under.
	TitleBarColor string `json:",omitempty"`
}
