// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ApacheCPULoad:            MetricConfig{Enabled: true},
					ApacheCPUTime:            MetricConfig{Enabled: true},
					ApacheCurrentConnections: MetricConfig{Enabled: true},
					ApacheLoad1:              MetricConfig{Enabled: true},
					ApacheLoad15:             MetricConfig{Enabled: true},
					ApacheLoad5:              MetricConfig{Enabled: true},
					ApacheRequestTime:        MetricConfig{Enabled: true},
					ApacheRequests:           MetricConfig{Enabled: true},
					ApacheScoreboard:         MetricConfig{Enabled: true},
					ApacheTraffic:            MetricConfig{Enabled: true},
					ApacheUptime:             MetricConfig{Enabled: true},
					ApacheWorkers:            MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ApacheServerName: ResourceAttributeConfig{Enabled: true},
					ApacheServerPort: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ApacheCPULoad:            MetricConfig{Enabled: false},
					ApacheCPUTime:            MetricConfig{Enabled: false},
					ApacheCurrentConnections: MetricConfig{Enabled: false},
					ApacheLoad1:              MetricConfig{Enabled: false},
					ApacheLoad15:             MetricConfig{Enabled: false},
					ApacheLoad5:              MetricConfig{Enabled: false},
					ApacheRequestTime:        MetricConfig{Enabled: false},
					ApacheRequests:           MetricConfig{Enabled: false},
					ApacheScoreboard:         MetricConfig{Enabled: false},
					ApacheTraffic:            MetricConfig{Enabled: false},
					ApacheUptime:             MetricConfig{Enabled: false},
					ApacheWorkers:            MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ApacheServerName: ResourceAttributeConfig{Enabled: false},
					ApacheServerPort: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg, confmap.WithIgnoreUnused()))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				ApacheServerName: ResourceAttributeConfig{Enabled: true},
				ApacheServerPort: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ApacheServerName: ResourceAttributeConfig{Enabled: false},
				ApacheServerPort: ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
