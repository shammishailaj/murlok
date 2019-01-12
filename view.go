package murlok

// View is a struct that describes a web view.
type View struct {
	// The url load when the view is created.
	URL string

	// The background color (#rrggbb).
	BackgroundColor string

	// Set the background as a frosted surface.
	FrostedBackground bool
}
