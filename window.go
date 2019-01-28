package murlok

// Window is a struct that describes a window.
type Window struct {
	// The background color (#rrggbb).
	BackgroundColor string `json:",omitempty"`

	// Sets the background as a frosted surface.
	FrostedBackground bool `json:",omitempty"`

	// The text color (#rrggbb).
	TextColor string `json:",omitempty"`

	// The title bar color (#rrggbb). When set, the web content appears below
	// the title bar rather than under.
	TitleBarColor string `json:",omitempty"`

	// The url load when the view is created.
	URL string
}
