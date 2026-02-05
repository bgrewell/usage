package internal

import "testing"

func TestConfiguration_Creation(t *testing.T) {
	tests := []struct {
		name        string
		appName     string
		version     string
		buildDate   string
		commitHash  string
		branch      string
		description string
	}{
		{
			name:        "basic configuration",
			appName:     "myapp",
			version:     "1.0.0",
			buildDate:   "2024-01-01",
			commitHash:  "abc123",
			branch:      "main",
			description: "Test app",
		},
		{
			name:        "minimal configuration",
			appName:     "app",
			version:     "",
			buildDate:   "",
			commitHash:  "",
			branch:      "",
			description: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Configuration{
				ApplicationName:        tt.appName,
				ApplicationVersion:     tt.version,
				ApplicationBuildDate:   tt.buildDate,
				ApplicationCommitHash:  tt.commitHash,
				ApplicationBranch:      tt.branch,
				ApplicationDescription: tt.description,
				Groups:                 make(map[string]*Group),
			}

			if config.ApplicationName != tt.appName {
				t.Errorf("ApplicationName = %q, want %q", config.ApplicationName, tt.appName)
			}
			if config.ApplicationVersion != tt.version {
				t.Errorf("ApplicationVersion = %q, want %q", config.ApplicationVersion, tt.version)
			}
			if config.ApplicationBuildDate != tt.buildDate {
				t.Errorf("ApplicationBuildDate = %q, want %q", config.ApplicationBuildDate, tt.buildDate)
			}
			if config.ApplicationCommitHash != tt.commitHash {
				t.Errorf("ApplicationCommitHash = %q, want %q", config.ApplicationCommitHash, tt.commitHash)
			}
			if config.ApplicationBranch != tt.branch {
				t.Errorf("ApplicationBranch = %q, want %q", config.ApplicationBranch, tt.branch)
			}
			if config.ApplicationDescription != tt.description {
				t.Errorf("ApplicationDescription = %q, want %q", config.ApplicationDescription, tt.description)
			}
			if config.Groups == nil {
				t.Error("Groups map should not be nil")
			}
		})
	}
}

func TestConfiguration_GroupsMap(t *testing.T) {
	t.Run("groups can be added to map", func(t *testing.T) {
		config := Configuration{
			ApplicationName: "test",
			Groups:          make(map[string]*Group),
		}

		group := &Group{
			Priority:    0,
			Name:        "TestGroup",
			Description: "Test",
		}

		config.Groups["TestGroup"] = group

		if len(config.Groups) != 1 {
			t.Errorf("Groups length = %d, want 1", len(config.Groups))
		}

		retrieved, ok := config.Groups["TestGroup"]
		if !ok {
			t.Error("Group not found in map")
		}
		if retrieved.Name != "TestGroup" {
			t.Errorf("Retrieved group name = %q, want %q", retrieved.Name, "TestGroup")
		}
	})
}
