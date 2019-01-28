package murlok

// MacOSConfig represents a MacOS package configuration.
type MacOSConfig struct {
	// The app name. It is displayed in the menubar and dock. The default value
	// is the package directory name.
	//
	// It is used only for goapp packaging.
	Name string `json:",omitempty"`

	// The UTI representing the app.
	//
	// It is used only for goapp packaging.
	ID string `json:",omitempty"`

	// The URL scheme that launches the app.
	//
	// It is used only for goapp packaging.
	URLScheme string `json:",omitempty"`

	// The version of the app (minified form eg. 1.42).
	//
	// It is used only for goapp packaging.
	Version string `json:",omitempty"`

	// The build number.
	//
	// It is used only for goapp packaging.
	BuildNumber int `json:",omitempty"`

	// The app icon path relative to the resources directory. It must be a
	// ".png". Provide a big one! Other required icon sizes will be auto
	// generated.
	//
	// It is used only for goapp packaging.
	Icon string `json:",omitempty"`

	// The development region.
	//
	// It is used only for goapp packaging.
	DevRegion string `json:",omitempty"`

	// A human readable copyright.
	//
	// It is used only for goapp packaging.
	Copyright string `json:",omitempty"`

	// The application category.
	//
	// It is used only for goapp packaging.
	Category MacOSCategory `json:",omitempty"`

	// Reports wheter the app runs in background mode. Background apps does not
	// appear in the dock and menubar.
	//
	// It is used only for goapp packaging.
	Background bool `json:",omitempty"`

	// Reports whether the app is a server (accepts incoming connections).
	//
	// It is used only for goapp packaging.
	Server bool `json:",omitempty"`

	// Reports whether the app uses the camera.
	//
	// It is used only for goapp packaging.
	Camera bool `json:",omitempty"`

	// Reports whether the app uses the microphone.
	//
	// It is used only for goapp packaging.
	Microphone bool `json:",omitempty"`

	// Reports whether the app uses the USB devices.
	//
	// It is used only for goapp packaging.
	USB bool `json:",omitempty"`

	// Reports whether the app uses printers.
	//
	// It is used only for goapp packaging.
	Printers bool `json:",omitempty"`

	// Reports whether the app uses bluetooth.
	//
	// It is used only for goapp packaging.
	Bluetooth bool `json:",omitempty"`

	// Reports whether the app has access to contacts.
	//
	// It is used only for goapp packaging.
	Contacts bool `json:",omitempty"`

	// Reports whether the app has access to device location.
	//
	// It is used only for goapp packaging.
	Location bool `json:",omitempty"`

	// Reports whether the app has access to calendars.
	//
	// It is used only for goapp packaging.
	Calendar bool `json:",omitempty"`

	// The file picker access mode.
	//
	// It is used only for goapp packaging.
	FilePickers MacOSFileAccess `json:",omitempty"`

	// The Download directory access mode.
	//
	// It is used only for goapp packaging.
	Downloads MacOSFileAccess `json:",omitempty"`

	// The Pictures directory access mode.
	//
	// It is used only for goapp packaging.
	Pictures MacOSFileAccess `json:",omitempty"`

	// The Music directory access mode.
	//
	// It is used only for goapp packaging.
	Music MacOSFileAccess `json:",omitempty"`

	// The Movies directory access mode.
	//
	// It is used only for goapp packaging.
	Movies MacOSFileAccess `json:",omitempty"`

	// The file types that can be opened by the app.
	//
	// It is used only for goapp packaging.
	SupportedFiles []MacOSFileType `json:",omitempty"`
}

// MacOSRole represents the role of an application.
type MacOSRole string

// Constants that enumerate application roles.
const (
	MacOSEditor MacOSRole = "Editor"
	MacOSViewer MacOSRole = "Viewer"
	MacOSShell  MacOSRole = "Shell"
)

// MacOSCategory represents the app style. The App Store uses this string to
// determine the appropriate categorization for the app.
// See https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/LaunchServicesKeys.html
// for more info.
type MacOSCategory string

// Constants that enumerate application categories.
const (
	MacOSBusinessApp             MacOSCategory = "public.app-category.business"
	MacOSDeveloperToolsApp                     = "public.app-category.developer-tools"
	MacOSEducationApp                          = "public.app-category.education"
	MacOSEntertainmentApp                      = "public.app-category.entertainment"
	MacOSFinanceApp                            = "public.app-category.finance"
	MacOSGamesApp                              = "public.app-category.games"
	MacOSGraphicsAndDesignApp                  = "public.app-category.graphics-design"
	MacOSHealthcareAndFitnessApp               = "public.app-category.healthcare-fitness"
	MacOSLifestyleApp                          = "public.app-category.lifestyle"
	MacOSMedicalApp                            = "public.app-category.medical"
	MacOSMusicApp                              = "public.app-category.music"
	MacOSNewsApp                               = "public.app-category.news"
	MacOSPhotographyApp                        = "public.app-category.photography"
	MacOSProductivityApp                       = "public.app-category.productivity"
	MacOSReferenceApp                          = "public.app-category.reference"
	MacOSSocialNetworkingApp                   = "public.app-category.social-networking"
	MacOSSportsApp                             = "public.app-category.sports"
	MacOSTravelApp                             = "public.app-category.travel"
	MacOSUtilitiesApp                          = "public.app-category.utilities"
	MacOSVideoApp                              = "public.app-category.video"
	MacOSWeatherApp                            = "public.app-category.weather"
	MacOSActionGamesApp                        = "public.app-category.action-games"
	MacOSAdventureGamesApp                     = "public.app-category.adventure-games"
	MacOSArcadeGamesApp                        = "public.app-category.arcade-games"
	MacOSBoardGamesApp                         = "public.app-category.board-games"
	MacOSCardGamesApp                          = "public.app-category.card-games"
	MacOSCasinoGamesApp                        = "public.app-category.casino-games"
	MacOSDiceGamesApp                          = "public.app-category.dice-games"
	MacOSEducationalGamesApp                   = "public.app-category.educational-games"
	MacOSFamilyGamesApp                        = "public.app-category.family-games"
	MacOSKidsGamesApp                          = "public.app-category.kids-games"
	MacOSMusicGamesApp                         = "public.app-category.music-games"
	MacOSPuzzleGamesApp                        = "public.app-category.puzzle-games"
	MacOSRacingGamesApp                        = "public.app-category.racing-games"
	MacOSRolePlayingGamesApp                   = "public.app-category.role-playing-games"
	MacOSSimulationGamesApp                    = "public.app-category.simulation-games"
	MacOSSportsGamesApp                        = "public.app-category.sports-games"
	MacOSStrategyGamesApp                      = "public.app-category.strategy-games"
	MacOSTriviaGamesApp                        = "public.app-category.trivia-games"
	MacOSWordGamesApp                          = "public.app-category.word-games"
)

// MacOSFileAccess represents a file access mode.
type MacOSFileAccess string

// Constants that enumerate file access modes.
const (
	NoAccess  MacOSFileAccess = ""
	ReadOnly  MacOSFileAccess = "read-only"
	ReadWrite MacOSFileAccess = "read-write"
)

// MacOSFileType describes a file type that can be opened by the app.
type MacOSFileType struct {
	// The  name.
	// Must be non empty, a single word and lowercased.
	Name string

	// The app’s role with respect to the type.
	Role MacOSRole

	// The icon path:
	// - Must be relative to the resources directory.
	// - Must be a ".png".
	Icon string

	// A list of UTI defining a supported file type.
	// Eg. "public.png" for ".png" files.
	UTIs []string
}
