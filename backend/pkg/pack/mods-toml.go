package pack

// ModsToml represents the structure of a Minecraft Forge mods.toml file
type ModsToml struct {
	ModLoader          string            `toml:"modLoader"`                    // Mandatory top-level field
	LowestMcVersion    string            `toml:"loaderVersion"`                // Mandatory top-level field
	License            string            `toml:"license,omitempty"`            // Optional top-level field
	IssueTrackerURL    string            `toml:"issueTrackerURL,omitempty"`    // Optional top-level field
	ShowAsResourcePack bool              `toml:"showAsResourcePack,omitempty"` // Optional top-level field (can contain # comments in TOML)
	Dependencies       DependencyAliases `toml:"dependencies,omitempty"`       // Optional aliases to use in dependency entries
	Mods               []Mod             `toml:"mods"`                         // Mandatory [[mods]] array (at least one entry required)
	DependencyEntries  []Dependency      `toml:"dependency,omitempty"`         // Optional array of dependency specifications
}

// DependencyAliases represents dependencies aliasing
// For each alias, there can be an entry in this struct
type DependencyAliases struct {
	Aliases map[string]string `toml:",omitempty"` // Using string maps for dynamic alias names
}

// Mod represents a single mod entry in the mods.toml
type Mod struct {
	// Mandatory fields
	ModID   string `toml:"modId"`
	Version string `toml:"version"`
	// Recommended fields
	DisplayName string `toml:"displayName,omitempty"`
	Description string `toml:"description,omitempty"`
	// Optional fields
	UpdateJSONURL       string   `toml:"updateJSONURL,omitempty"`
	DisplayURL          string   `toml:"displayURL,omitempty"`
	LogoFile            string   `toml:"logoFile,omitempty"`
	Credits             string   `toml:"credits,omitempty"`
	Authors             string   `toml:"authors,omitempty"`
	DisplayTest         string   `toml:"displayTest,omitempty"`
	ConfigCommentHeader []string `toml:"configBackground,omitempty"`
	// Optional array of entrypoints (like clientEntrypoint: "package.path")
	Entrypoints map[string]string `toml:",omitempty"` // Using map for dynamic entrypoint names
}

// Dependency represents a dependency entry
type Dependency struct {
	// Mandatory fields
	ModID     string `toml:"modId"`
	Mandatory bool   `toml:"mandatory"`
	// Versioning fields - all are optional
	VersionRange string `toml:"versionRange,omitempty"`
	Ordering     string `toml:"ordering,omitempty"`
	Side         string `toml:"side,omitempty"`
}
