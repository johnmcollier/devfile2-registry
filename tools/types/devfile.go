package types

// Devfile200 partial Devfile schema.
// Only contains structs for projects at the moment, as we don't need to parse the entire devfile
type Devfile200 struct {
	// Projects worked on in the workspace, containing names and sources locations
	Projects []Project `json:"projects,omitempty" yaml:"projects,omitempty"`
}

// Project project defined in devfile
type Project struct {

	// Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.
	ClonePath string `json:"clonePath,omitempty" yaml:"clonePath,omitempty"`

	// Project's Git source
	Git *Git `json:"git,omitempty" yaml:"git,omitempty"`

	// Project's GitHub source
	Github *Github `json:"github,omitempty" yaml:"github,omitempty"`

	// Project name
	Name string `json:"name" yaml:"name"`

	// Project's Zip source
	Zip *Zip `json:"zip,omitempty" yaml:"zip,omitempty"`
}

// Git Project's Git source
type Git struct {

	// The branch to check
	Branch string `json:"branch,omitempty"`

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty"`

	// The tag or commit id to reset the checked out branch to
	StartPoint string `json:"startPoint,omitempty"`
}

// Github Project's GitHub source
type Github struct {

	// The branch to check
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty" yaml:"sparseCheckoutDir,omitempty"`

	// The tag or commit id to reset the checked out branch to
	StartPoint string `json:"startPoint,omitempty" yaml:"startPoint,omitempty"`
}

// Zip Project's Zip source
type Zip struct {

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty" yaml:"sparseCheckoutDir,omitempty"`
}
