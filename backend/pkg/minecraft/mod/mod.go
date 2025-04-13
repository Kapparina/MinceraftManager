package mod

// ModInfo represents the essential information for a Minecraft Forge mod
//
//goland:noinspection GoNameStartsWithPackageN
//goland:noinspection GoNameStartsWithPackageName
type ModInfo struct {
	// ModID is the unique identifier for the mod (lowercase, no spaces)
	ModID string `json:"modId"`

	// Version of the mod following semantic versioning
	Version string `json:"version"`

	// DisplayName is the user-friendly name of the mod
	DisplayName string `json:"displayName"`

	// Description provides information about what the mod does
	Description string `json:"description"`

	// Authors is a list of people who created the mod
	Authors []string `json:"authors"`

	// Credits acknowledges people who contributed but aren't authors
	Credits string `json:"credits,omitempty"`

	// LogoFile is the path to the mod's logo relative to the mod file
	LogoFile string `json:"logoFile,omitempty"`

	// UpdateJSONURL is the URL to a JSON file with update information
	UpdateJSONURL string `json:"updateJSONURL,omitempty"`

	// Dependencies lists other mods this mod depends on
	Dependencies []Dependency `json:"dependencies,omitempty"`

	// IssueTrackerURL is the URL to the mod's issue tracker
	IssueTrackerURL string `json:"issueTrackerURL,omitempty"`

	// License specifies the mod's license (e.g., MIT, GPL)
	License string `json:"license"`
}

// Dependency represents a mod dependency
type Dependency struct {
	// ModID of the dependency
	ModID string `json:"modId"`

	// Mandatory indicates if the dependency is required
	Mandatory bool `json:"mandatory"`

	// VersionRange specifies the acceptable versions of the dependency
	// Format examples: "[1.16.5,1.17)", "1.16.5"
	VersionRange string `json:"versionRange,omitempty"`

	// Ordering defines load order relation (NONE, BEFORE, AFTER)
	Ordering string `json:"ordering,omitempty"`

	// Side specifies where this dependency is needed (CLIENT, SERVER, BOTH)
	Side string `json:"side,omitempty"`
}

// MainClass represents the entry point of the mod
type MainClass struct {
	// Package name and class name that extends the appropriate mod loader base class
	ClassName string
}

// ModConfig represents the configuration of a Minecraft mod
//
//goland:noinspection GoNameStartsWithPacka
//goland:noinspection GoNa
//goland:noinspection GoNameStartsWithPackageName
type ModConfig struct {
	// Basic information about the mod
	Info ModInfo

	// Entry point of the mod
	MainClass MainClass

	// Resources contains paths to assets, data, etc.
	Resources struct {
		Assets string
		Data   string
	}
}
