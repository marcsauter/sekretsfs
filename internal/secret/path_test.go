package secret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Run("invalid path", func(t *testing.T) {
		invalid := []string{
			"default",
			"/default",
			"/default/",
			"default/",
			"default/secret/key/more",
			"default/secret/key/more/",
			"/default/secret/key/more/",
			"/default/secret/key/more",
		}

		for _, n := range invalid {
			p, err := splitPath(n)
			assert.Error(t, err)
			assert.Nil(t, p)
		}
	})

	t.Run("valid path", func(t *testing.T) {
		valid := []string{
			"default/secret",
			"/default/secret",
			"/default/secret/",
			"default/secret/",
			"default/secret/key",
			"default/secret/key/",
			"/default/secret/key/",
			"/default/secret/key",
		}

		for _, n := range valid {
			p, err := splitPath(n)
			assert.NoError(t, err)
			assert.NotNil(t, p)
		}
	})

	t.Run("valid dir path", func(t *testing.T) {
		validDir := []string{
			"default/secret",
			"/default/secret",
			"/default/secret/",
			"default/secret/",
		}

		for _, n := range validDir {
			p, err := splitPath(n)
			assert.NoError(t, err)
			assert.NotNil(t, p)
			assert.True(t, p.IsDir())
			assert.Equal(t, "default", p.Namespace())
			assert.Equal(t, "secret", p.Secret())
			assert.Empty(t, p.Key())
		}
	})

	t.Run("valid file path", func(t *testing.T) {
		validDir := []string{
			"default/secret/key",
			"default/secret/key/",
			"/default/secret/key/",
			"/default/secret/key",
		}

		for _, n := range validDir {
			p, err := splitPath(n)
			assert.NoError(t, err)
			assert.NotNil(t, p)
			assert.False(t, p.IsDir())
			assert.Equal(t, "default", p.Namespace())
			assert.Equal(t, "secret", p.Secret())
			assert.Equal(t, "key", p.Key())
		}
	})
}