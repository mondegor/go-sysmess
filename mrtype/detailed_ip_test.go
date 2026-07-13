package mrtype_test

import (
	"encoding/json"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/mrtype"
)

func TestNewIP(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name   string
		realIP netip.Addr
		want   mrtype.DetailedIP
	}

	tests := []testCase{
		{
			name:   "empty",
			realIP: netip.Addr{},
			want:   mrtype.DetailedIP{},
		},
		{
			name:   "ipv4",
			realIP: netip.MustParseAddr("127.0.0.1"),
			want:   mrtype.DetailedIP{Real: netip.MustParseAddr("127.0.0.1")},
		},
		{
			name:   "ipv6",
			realIP: netip.MustParseAddr("2001:db8::1"),
			want:   mrtype.DetailedIP{Real: netip.MustParseAddr("2001:db8::1")},
		},
		{
			name:   "ipv4_mapped_is_unmapped",
			realIP: netip.MustParseAddr("::ffff:127.0.0.1"),
			want:   mrtype.DetailedIP{Real: netip.MustParseAddr("127.0.0.1")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, mrtype.NewIP(tt.realIP))
		})
	}
}

func TestNewDetailedIP(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		realIP  netip.Addr
		proxyIP netip.Addr
		want    mrtype.DetailedIP
	}

	tests := []testCase{
		{
			name:    "both_empty",
			realIP:  netip.Addr{},
			proxyIP: netip.Addr{},
			want:    mrtype.DetailedIP{},
		},
		{
			name:   "real_only",
			realIP: netip.MustParseAddr("127.0.0.1"),
			want:   mrtype.DetailedIP{Real: netip.MustParseAddr("127.0.0.1")},
		},
		{
			name:    "proxy_only",
			proxyIP: netip.MustParseAddr("10.0.0.1"),
			want:    mrtype.DetailedIP{Proxy: netip.MustParseAddr("10.0.0.1")},
		},
		{
			name:    "both_ipv4",
			realIP:  netip.MustParseAddr("127.0.0.1"),
			proxyIP: netip.MustParseAddr("10.0.0.1"),
			want: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("127.0.0.1"),
				Proxy: netip.MustParseAddr("10.0.0.1"),
			},
		},
		{
			name:    "both_ipv6",
			realIP:  netip.MustParseAddr("2001:db8::1"),
			proxyIP: netip.MustParseAddr("fe80::1"),
			want: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("2001:db8::1"),
				Proxy: netip.MustParseAddr("fe80::1"),
			},
		},
		{
			name:    "mixed_versions",
			realIP:  netip.MustParseAddr("2001:db8::1"),
			proxyIP: netip.MustParseAddr("10.0.0.1"),
			want: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("2001:db8::1"),
				Proxy: netip.MustParseAddr("10.0.0.1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, mrtype.NewDetailedIP(tt.realIP, tt.proxyIP))
		})
	}
}

func TestDetailedIP_String(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		ip   mrtype.DetailedIP
		want string
	}

	tests := []testCase{
		{
			name: "empty",
			ip:   mrtype.DetailedIP{},
			want: "",
		},
		{
			name: "real_only",
			ip:   mrtype.DetailedIP{Real: netip.MustParseAddr("127.0.0.1")},
			want: "127.0.0.1",
		},
		{
			name: "real_only_ipv6",
			ip:   mrtype.DetailedIP{Real: netip.MustParseAddr("2001:db8::1")},
			want: "2001:db8::1",
		},
		{
			name: "proxy_unspecified",
			ip: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("127.0.0.1"),
				Proxy: netip.MustParseAddr("0.0.0.0"),
			},
			want: "127.0.0.1",
		},
		{
			name: "real_and_proxy",
			ip: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("127.0.0.1"),
				Proxy: netip.MustParseAddr("10.0.0.1"),
			},
			want: "127.0.0.1, 10.0.0.1",
		},
		{
			name: "real_and_proxy_ipv6",
			ip: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("2001:db8::1"),
				Proxy: netip.MustParseAddr("fe80::1"),
			},
			want: "2001:db8::1, fe80::1",
		},
		{
			name: "proxy_only",
			ip:   mrtype.DetailedIP{Proxy: netip.MustParseAddr("10.0.0.1")},
			want: "0, 10.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.ip.String())
		})
	}
}

func TestDetailedIP_JSON(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		ip   mrtype.DetailedIP
		want string
	}

	tests := []testCase{
		{
			name: "empty",
			ip:   mrtype.DetailedIP{},
			want: `{"real":"","proxy":""}`,
		},
		{
			name: "real_only",
			ip:   mrtype.DetailedIP{Real: netip.MustParseAddr("127.0.0.1")},
			want: `{"real":"127.0.0.1","proxy":""}`,
		},
		{
			name: "real_and_proxy",
			ip: mrtype.DetailedIP{
				Real:  netip.MustParseAddr("2001:db8::1"),
				Proxy: netip.MustParseAddr("10.0.0.1"),
			},
			want: `{"real":"2001:db8::1","proxy":"10.0.0.1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tt.ip)
			require.NoError(t, err)
			assert.JSONEq(t, tt.want, string(data))

			var got mrtype.DetailedIP

			require.NoError(t, json.Unmarshal(data, &got))
			assert.Equal(t, tt.ip, got)
		})
	}
}
