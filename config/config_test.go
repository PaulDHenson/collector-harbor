package config

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestConfig(t *testing.T) {
	Expect := NewGomegaWithT(t).Expect

	for _, tc := range []struct {
		name        string
		flags       []string
		expected    *Config
		expectError bool
	}{
		{
			name:  "defaults",
			flags: []string{},
			expected: &Config{
				Port:  8080,
				Debug: false,
				RodeConfig: &RodeConfig{
					Host: "rode:50051",
				},
				HarborConfig: &HarborConfig{
					Host: "http://harbor-harbor-core",
				},
			},
		},
		{
			name:        "bad port",
			flags:       []string{"--port=foo"},
			expectError: true,
		},
		{
			name:        "bad debug",
			flags:       []string{"--debug=bar"},
			expectError: true,
		},
		{
			name:  "Rode host",
			flags: []string{"--rode-host=bar"},
			expected: &Config{
				Port:  8080,
				Debug: false,
				RodeConfig: &RodeConfig{
					Host: "bar",
				},
				HarborConfig: &HarborConfig{
					Host: "http://harbor-harbor-core",
				},
			},
		},
		{
			name:  "Harbor host",
			flags: []string{"--harbor-host=foo", "--harbor-username=bar", "--harbor-password=baz", "--harbor-insecure"},
			expected: &Config{
				Port:  8080,
				Debug: false,
				RodeConfig: &RodeConfig{
					Host: "rode:50051",
				},
				HarborConfig: &HarborConfig{
					Host:     "foo",
					Username: "bar",
					Password: "baz",
					Insecure: true,
				},
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c, err := Build("rode-collector-harbor", tc.flags)

			if tc.expectError {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(err).ToNot(HaveOccurred())
				Expect(c).To(BeEquivalentTo(tc.expected))
			}
		})
	}
}
