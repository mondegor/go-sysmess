package mrtype_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/mrtype"
)

func TestNewIP(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name   string
		realIP uint32
		want   mrtype.DetailedIP
	}

	tests := []testCase{
		{
			name:   "zero",
			realIP: 0,
			want:   mrtype.DetailedIP{},
		},
		{
			name:   "real_only",
			realIP: 0x7F000001, // 127.0.0.1
			want:   mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}},
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
		realIP  uint32
		proxyIP uint32
		want    mrtype.DetailedIP
	}

	tests := []testCase{
		{
			name:    "both_zero",
			realIP:  0,
			proxyIP: 0,
			want:    mrtype.DetailedIP{},
		},
		{
			name:    "real_only",
			realIP:  0x7F000001, // 127.0.0.1
			proxyIP: 0,
			want:    mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}},
		},
		{
			name:    "proxy_only",
			realIP:  0,
			proxyIP: 0x0A000001, // 10.0.0.1
			want:    mrtype.DetailedIP{Proxy: net.IP{10, 0, 0, 1}},
		},
		{
			name:    "both",
			realIP:  0x7F000001, // 127.0.0.1
			proxyIP: 0x0A000001, // 10.0.0.1
			want:    mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}, Proxy: net.IP{10, 0, 0, 1}},
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
			ip:   mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}},
			want: "127.0.0.1",
		},
		{
			name: "proxy_unspecified",
			ip:   mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}, Proxy: net.IPv4zero},
			want: "127.0.0.1",
		},
		{
			name: "real_and_proxy",
			ip:   mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}, Proxy: net.IP{10, 0, 0, 1}},
			want: "127.0.0.1, 10.0.0.1",
		},
		{
			name: "proxy_only",
			ip:   mrtype.DetailedIP{Proxy: net.IP{10, 0, 0, 1}},
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

func TestDetailedIP_ToUint(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		ip          mrtype.DetailedIP
		wantRealIP  uint32
		wantProxyIP uint32
		wantErr     bool
	}

	tests := []testCase{
		{
			name:        "empty",
			ip:          mrtype.DetailedIP{},
			wantRealIP:  0,
			wantProxyIP: 0,
		},
		{
			name:        "real_only",
			ip:          mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}},
			wantRealIP:  0x7F000001,
			wantProxyIP: 0,
		},
		{
			name:        "real_and_proxy",
			ip:          mrtype.DetailedIP{Real: net.IP{127, 0, 0, 1}, Proxy: net.IP{10, 0, 0, 1}},
			wantRealIP:  0x7F000001,
			wantProxyIP: 0x0A000001,
		},
		{
			name:    "ipv6_real",
			ip:      mrtype.DetailedIP{Real: net.ParseIP("2001:db8::1")},
			wantErr: true,
		},
		{
			name:    "incorrect_real",
			ip:      mrtype.DetailedIP{Real: net.IP{1, 2, 3}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			realIP, proxyIP, err := tt.ip.ToUint()
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantRealIP, realIP)
			assert.Equal(t, tt.wantProxyIP, proxyIP)
		})
	}
}
