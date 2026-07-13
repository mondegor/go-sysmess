package parse_test

import (
	"net/netip"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/mrtype/parse"
)

func TestIP(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		value    string
		required bool
		want     netip.Addr
		wantErr  bool
	}

	tests := []testCase{
		{
			name:     "empty_not_required",
			value:    "",
			required: false,
			want:     netip.Addr{},
		},
		{
			name:     "spaces_only_not_required",
			value:    "   ",
			required: false,
			want:     netip.Addr{},
		},
		{
			name:     "empty_required",
			value:    "",
			required: true,
			wantErr:  true,
		},
		{
			name:     "ipv4",
			value:    "127.0.0.1",
			required: true,
			want:     netip.MustParseAddr("127.0.0.1"),
		},
		{
			name:     "ipv4_trimmed",
			value:    "  10.0.0.1  ",
			required: true,
			want:     netip.MustParseAddr("10.0.0.1"),
		},
		{
			name:     "ipv4_with_port",
			value:    "10.0.0.1:8080",
			required: true,
			want:     netip.MustParseAddr("10.0.0.1"),
		},
		{
			name:     "ipv6",
			value:    "2001:db8::1",
			required: true,
			want:     netip.MustParseAddr("2001:db8::1"),
		},
		{
			name:     "ipv6_full_form",
			value:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			required: true,
			want:     netip.MustParseAddr("2001:db8:85a3::8a2e:370:7334"),
		},
		{
			name:     "ipv6_with_port",
			value:    "[2001:db8::1]:8080",
			required: true,
			want:     netip.MustParseAddr("2001:db8::1"),
		},
		{
			name:     "ipv4_mapped_is_unmapped",
			value:    "::ffff:127.0.0.1",
			required: true,
			want:     netip.MustParseAddr("127.0.0.1"),
		},
		{
			name:     "too_long",
			value:    strings.Repeat("1", 65),
			required: true,
			wantErr:  true,
		},
		{
			name:     "incorrect",
			value:    "abc",
			required: true,
			wantErr:  true,
		},
		{
			name:     "incorrect_ipv4",
			value:    "127.0.0.256",
			required: true,
			wantErr:  true,
		},
		{
			name:     "incorrect_ipv6",
			value:    "2001:db8::zz",
			required: true,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := parse.IP(tt.value, tt.required)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
