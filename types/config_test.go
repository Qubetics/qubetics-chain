package types

import (
	"sync"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestConfig_GetBech32NodeAddrPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_addr": "qubetics"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubetics",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_addr": "qubeticsnode"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticsnode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32NodeAddrPrefix(); got != tt.want {
				t.Errorf("GetBech32NodeAddrPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32NodePubPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_pub": "qubeticspub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticspub",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_pub": "qubeticsnodepub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticsnodepub",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32NodePubPrefix(); got != tt.want {
				t.Errorf("GetBech32NodePubPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32ProviderAddrPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_addr": "qubetics"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubetics",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_addr": "qubeticsprov"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticsprov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32ProviderAddrPrefix(); got != tt.want {
				t.Errorf("GetBech32ProviderAddrPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32ProviderPubPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_pub": "qubeticspub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticspub",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_pub": "qubeticsprovpub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"qubeticsprovpub",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32ProviderPubPrefix(); got != tt.want {
				t.Errorf("GetBech32ProviderPubPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
