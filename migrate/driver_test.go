package migrate

import (
	"testing"

	st "github.com/golang-migrate/migrate/v4/source/testing"
	"github.com/stretchr/testify/require"

	"github.com/klingtnet/embed/internal"
)

func Test(t *testing.T) {
	d, err := WithInstance(internal.Embeds)
	require.NoError(t, err, "WithInstance")

	// For test details refer to:
	// https://github.com/golang-migrate/migrate/blob/04a572114abc0db2633c2cc473542bef437c8a8c/source/testing/testing.go#L22
	st.Test(t, d)
}
