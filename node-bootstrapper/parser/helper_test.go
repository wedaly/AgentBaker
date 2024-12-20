/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"encoding/base64"
	"reflect"
	"testing"

	nbcontractv1 "github.com/Azure/agentbaker/pkg/proto/nbcontract/v1"
	"github.com/Azure/go-autorest/autorest/to"
)

func Test_getSysctlContent(t *testing.T) {
	// Test_getSysctlContent tests the getSysctlContent function.
	type args struct {
		s *nbcontractv1.SysctlConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default SysctlConfig",
			args: args{
				s: &nbcontractv1.SysctlConfig{},
			},
			want: base64.StdEncoding.EncodeToString(
				[]byte(`net.core.message_burst=80
net.core.message_cost=40
net.core.somaxconn=16384
net.ipv4.neigh.default.gc_thresh1=4096
net.ipv4.neigh.default.gc_thresh2=8192
net.ipv4.neigh.default.gc_thresh3=16384
net.ipv4.tcp_max_syn_backlog=16384
net.ipv4.tcp_retries2=8`)),
		},
		{
			name: "SysctlConfig with custom values",
			args: args{
				s: &nbcontractv1.SysctlConfig{
					NetIpv4TcpMaxSynBacklog: to.Int32Ptr(int32(9999)),
					NetCoreRmemDefault:      to.Int32Ptr(int32(9999)),
					NetIpv4IpLocalPortRange: to.StringPtr("32768 62535"),
				},
			},
			want: base64.StdEncoding.EncodeToString(
				[]byte(`net.core.message_burst=80
net.core.message_cost=40
net.core.rmem_default=9999
net.core.somaxconn=16384
net.ipv4.ip_local_port_range=32768 62535
net.ipv4.neigh.default.gc_thresh1=4096
net.ipv4.neigh.default.gc_thresh2=8192
net.ipv4.neigh.default.gc_thresh3=16384
net.ipv4.tcp_max_syn_backlog=9999
net.ipv4.tcp_retries2=8`)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSysctlContent(tt.args.s); got != tt.want {
				t.Errorf("getSysctlContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUlimitContent(t *testing.T) {
	type args struct {
		u *nbcontractv1.UlimitConfig
	}
	str9999 := "9999"
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default UlimitConfig",
			args: args{
				u: &nbcontractv1.UlimitConfig{},
			},
			want: base64.StdEncoding.EncodeToString(
				[]byte("[Service]\n")),
		},
		{
			name: "UlimitConfig with custom values",
			args: args{
				u: &nbcontractv1.UlimitConfig{
					NoFile:          &str9999,
					MaxLockedMemory: &str9999,
				},
			},
			want: base64.StdEncoding.EncodeToString(
				[]byte("[Service]\nLimitMEMLOCK=9999 LimitNOFILE=9999")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUlimitContent(tt.args.u); got != tt.want {
				t.Errorf("getUlimitContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createSortedKeyValueStringPairs(t *testing.T) {
	type args struct {
		m         map[string]interface{}
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty map",
			args: args{
				m:         map[string]interface{}{},
				delimiter: ",",
			},
			want: "",
		},
		{
			name: "Single key-value pair",
			args: args{
				m:         map[string]interface{}{"key1": "value1"},
				delimiter: " ",
			},
			want: "key1=value1",
		},
		{
			name: "Multiple key-value pairs with delimiter ,",
			args: args{
				m:         map[string]interface{}{"key1": "value1", "key2": "value2"},
				delimiter: ",",
			},
			want: "key1=value1,key2=value2",
		},
		{
			name: "Multiple key-value pairs with delimiter space",
			args: args{
				m:         map[string]interface{}{"key1": "value1", "key2": "value2"},
				delimiter: " ",
			},
			want: "key1=value1 key2=value2",
		},
		{
			name: "Sorting key-value pairs",
			args: args{
				m:         map[string]interface{}{"b": "valb", "a": "vala", "c": "valc"},
				delimiter: ",",
			},
			want: "a=vala,b=valb,c=valc",
		},
		{
			name: "Multiple key-value pairs with delimiter line breaker \\n where values are a combination of strings and integers",
			args: args{
				m:         map[string]interface{}{"key1": "value1", "key2": "value2", "key3": 3, "key4": 4},
				delimiter: "\n",
			},
			want: "key1=value1\nkey2=value2\nkey3=3\nkey4=4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSortedKeyValuePairs(tt.args.m, tt.args.delimiter); got != tt.want {
				t.Errorf("createSortedKeyValuePairs() with map[string]string = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createSortedKeyValueInt32Pairs(t *testing.T) {
	type args struct {
		m         map[string]int32
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty map",
			args: args{
				m:         map[string]int32{},
				delimiter: ",",
			},
			want: "",
		},
		{
			name: "Single key-value pair",
			args: args{
				m:         map[string]int32{"key1": 1},
				delimiter: " ",
			},
			want: "key1=1",
		},
		{
			name: "Multiple key-value pairs",
			args: args{
				m:         map[string]int32{"key1": 1, "key2": 2},
				delimiter: ",",
			},
			want: "key1=1,key2=2",
		},
		{
			name: "Multiple key-value pairs with delimiter space",
			args: args{
				m:         map[string]int32{"key1": 1, "key2": 2},
				delimiter: " ",
			},
			want: "key1=1 key2=2",
		},
		{
			name: "Sorting key-value pairs",
			args: args{
				m:         map[string]int32{"b": 2, "a": 1, "c": 3},
				delimiter: ",",
			},
			want: "a=1,b=2,c=3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSortedKeyValuePairs(tt.args.m, tt.args.delimiter); got != tt.want {
				t.Errorf("createSortedKeyValuePairs() with map[string]int32 = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getContainerdConfig(t *testing.T) {
	type args struct {
		nbcontract *nbcontractv1.Configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default Configuration",
			args: args{
				nbcontract: &nbcontractv1.Configuration{
					NeedsCgroupv2: to.BoolPtr(true),
				},
			},
			want: base64.StdEncoding.EncodeToString([]byte(`version = 2
oom_score = 0
[plugins."io.containerd.grpc.v1.cri"]
  sandbox_image = ""
  [plugins."io.containerd.grpc.v1.cri".containerd]
    default_runtime_name = "runc"
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
      runtime_type = "io.containerd.runc.v2"
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
      BinaryName = "/usr/bin/runc"
      SystemdCgroup = true
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
      runtime_type = "io.containerd.runc.v2"
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted.options]
      BinaryName = "/usr/bin/runc"
  [plugins."io.containerd.grpc.v1.cri".registry.headers]
    X-Meta-Source-Client = ["azure/aks"]
[metrics]
  address = "0.0.0.0:10257"
`)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContainerdConfig(tt.args.nbcontract); got != tt.want {
				t.Errorf("getContainerdConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKubenetTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Kubenet template",
			want: base64.StdEncoding.EncodeToString([]byte(`{
	"cniVersion": "0.3.1",
	"name": "kubenet",
	"plugins": [{
		"type": "bridge",
		"bridge": "cbr0",
		"mtu": 1500,
		"addIf": "eth0",
		"isGateway": true,
		"ipMasq": false,
		"promiscMode": true,
		"hairpinMode": false,
		"ipam": {
			"type": "host-local",
			"ranges": [{{range $i, $range := .PodCIDRRanges}}{{if $i}}, {{end}}[{"subnet": "{{$range}}"}]{{end}}],
			"routes": [{{range $i, $route := .Routes}}{{if $i}}, {{end}}{"dst": "{{$route}}"}{{end}}]
		}
	},
	{
		"type": "portmap",
		"capabilities": {"portMappings": true},
		"externalSetMarkChain": "KUBE-MARK-MASQ"
	}]
}
`)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKubenetTemplate(); got != tt.want {
				t.Errorf("getKubenetTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAzureEnvironmentFilepath(t *testing.T) {
	type args struct {
		v *nbcontractv1.Configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Nil CustomCloudConfig",
			args: args{
				v: &nbcontractv1.Configuration{},
			},
			want: "",
		},
		{
			name: "Empty AzureEnvironmentFilepath",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{},
				},
			},
			want: "",
		},
		{
			name: "AzureEnvironmentFilepath when it is AKSCustomCloud",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{
						CustomCloudEnvName: nbcontractv1.AksCustomCloudName,
					},
				},
			},
			want: "/etc/kubernetes/akscustom.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAzureEnvironmentFilepath(tt.args.v); got != tt.want {
				t.Errorf("getAzureEnvironmentFilepath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnsureNoDupePromiscuousBridge(t *testing.T) {
	type args struct {
		nc *nbcontractv1.NetworkConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "NetworkConfig with no promiscuous bridge",
			args: args{
				nc: &nbcontractv1.NetworkConfig{
					NetworkPlugin: nbcontractv1.NetworkPlugin_NP_AZURE,
					NetworkPolicy: nbcontractv1.NetworkPolicy_NPO_AZURE,
				},
			},
			want: false,
		},
		{
			name: "NetworkConfig with promiscuous bridge",
			args: args{
				nc: &nbcontractv1.NetworkConfig{
					NetworkPlugin: nbcontractv1.NetworkPlugin_NP_KUBENET,
					NetworkPolicy: nbcontractv1.NetworkPolicy_NPO_AZURE,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnsureNoDupePromiscuousBridge(tt.args.nc); got != tt.want {
				t.Errorf("getEnsureNoDupePromiscuousBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHasSearchDomain(t *testing.T) {
	type args struct {
		csd *nbcontractv1.CustomSearchDomainConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CustomSearchDomain with empty search domain should return false",
			args: args{
				csd: &nbcontractv1.CustomSearchDomainConfig{},
			},
			want: false,
		},
		{
			name: "CustomSearchDomain with empty search domain user should return false",
			args: args{
				csd: &nbcontractv1.CustomSearchDomainConfig{
					DomainName:    "fakedomain.com",
					RealmPassword: "fakepassword",
				},
			},
			want: false,
		},
		{
			name: "CustomSearchDomain with search domain, user and password should return true",
			args: args{
				csd: &nbcontractv1.CustomSearchDomainConfig{
					DomainName:    "fakedomain.com",
					RealmUser:     "fakeuser",
					RealmPassword: "fakepassword",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHasSearchDomain(tt.args.csd); got != tt.want {
				t.Errorf("getHasSearchDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCustomCACertsStatus(t *testing.T) {
	type args struct {
		customCACerts []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty customCACerts",
			args: args{
				customCACerts: []string{},
			},
			want: false,
		},
		{
			name: "Non-empty customCACerts",
			args: args{
				customCACerts: []string{"cert1", "cert2"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCustomCACertsStatus(tt.args.customCACerts); got != tt.want {
				t.Errorf("getCustomCACertsStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringifiedStringArray(t *testing.T) {
	type args struct {
		arr       []string
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty array input should return empty string",
			args: args{
				arr:       []string{},
				delimiter: ",",
			},
			want: "",
		},
		{
			name: "Single element array input should return the element without delimiter",
			args: args{
				arr:       []string{"element1"},
				delimiter: ",",
			},
			want: "element1",
		},
		{
			name: "Multiple element array input should return the elements separated by delimiter",
			args: args{
				arr:       []string{"element1", "element2", "element3"},
				delimiter: ",",
			},
			want: "element1,element2,element3",
		},
		{
			name: "Multiple element array input with space delimiter should return the elements separated by space",
			args: args{
				arr:       []string{"element1", "element2", "element3"},
				delimiter: " ",
			},
			want: "element1 element2 element3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringifiedStringArray(tt.args.arr, tt.args.delimiter); got != tt.want {
				t.Errorf("getStringifiedStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deref(t *testing.T) {
	type args struct {
		p *interface{}
	}
	var s interface{} = "test"
	var i interface{} = 123
	var b interface{} = true
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Dereference nil pointer should not panic and should return interface{}(nil)",
			args: args{
				p: nil,
			},
			want: interface{}(nil),
		},
		{
			name: "Dereference pointer to string should return the string",
			args: args{
				p: &s,
			},
			want: "test",
		},
		{
			name: "Dereference pointer to int should return the int",
			args: args{
				p: &i,
			},
			want: 123,
		},
		{
			name: "Dereference pointer to bool should return the bool",
			args: args{
				p: &b,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deref(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKubernetesVersionGe(t *testing.T) {
	type args struct {
		actualVersion string
		version       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1.24.5 is greater than or equal to 1.24.5",
			args: args{
				actualVersion: "1.24.5",
				version:       "1.24.5",
			},
			want: true,
		},
		{
			name: "1.23.5 is greater than or equal to 1.23.4",
			args: args{
				actualVersion: "1.23.5",
				version:       "1.23.4",
			},
			want: true,
		},
		{
			name: "1.23.5 is not greater than or equal to 1.24.4",
			args: args{
				actualVersion: "1.23.5",
				version:       "1.24.4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nbcontractv1.IsKubernetesVersionGe(tt.args.actualVersion, tt.args.version); got != tt.want {
				t.Errorf("IsKubernetesVersionGe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPortRangeEndValue(t *testing.T) {
	type args struct {
		portRange string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty string",
			args: args{
				portRange: "",
			},
			want: -1,
		},
		{
			name: "Port range with single port",
			args: args{
				portRange: "80",
			},
			want: -1,
		},
		{
			name: "Port range with valid start and end port",
			args: args{
				portRange: "80 90",
			},
			want: 90,
		},
		{
			name: "Port range with start and end port separated by - (invalid delimiter)",
			args: args{
				portRange: "80-90",
			},
			want: -1,
		},
		{
			name: "Port range with 3 ports (invalid)",
			args: args{
				portRange: "80 90 100",
			},
			want: -1,
		},
		{
			name: "start value is larger than end value (invalid)",
			args: args{
				portRange: "110 90",
			},
			want: -1,
		},
		{
			name: "start value is equal to end value (invalid)",
			args: args{
				portRange: "80 80",
			},
			want: -1,
		},
		{
			name: "either value is not a number (invalid)",
			args: args{
				portRange: "80 abc",
			},
			want: -1,
		},
		{
			name: "either value is smaller than 0 (invalid)",
			args: args{
				portRange: "-1 80",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPortRangeEndValue(tt.args.portRange); got != tt.want {
				t.Errorf("getPortRangeEndValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShouldConfigureHTTPProxy(t *testing.T) {
	type args struct {
		httpProxyConfig *nbcontractv1.HTTPProxyConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Nil HTTPProxyConfig",
			args: args{},
			want: false,
		},
		{
			name: "Empty HTTPProxyConfig",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{},
			},
			want: false,
		},
		{
			name: "HTTPProxyConfig with empty HttpProxy and valid HttpsProxy",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{
					HttpProxy:  "",
					HttpsProxy: "https://fakeproxy.com:8080",
				},
			},
			want: true,
		},
		{
			name: "HTTPProxyConfig with valid HttpProxy and empty HttpsProxy",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{
					HttpProxy:  "http://fakeproxy.com:8080",
					HttpsProxy: "",
				},
			},
			want: true,
		},
		{
			name: "HTTPProxyConfig with empty HttpProxy and empty HttpsProxy",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{
					HttpProxy:      "",
					HttpsProxy:     "",
					NoProxyEntries: []string{"fakesite1.com", "fakesite2.com"},
				},
			},
			want: false,
		},
		{
			name: "HTTPProxyConfig with valid HttpProxy",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{
					HttpProxy: "http://fakeproxy.com:8080",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShouldConfigureHTTPProxy(tt.args.httpProxyConfig); got != tt.want {
				t.Errorf("getShouldConfigureHTTPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShouldConfigureHTTPProxyCA(t *testing.T) {
	type args struct {
		httpProxyConfig *nbcontractv1.HTTPProxyConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Nil HTTPProxyConfig",
			args: args{},
			want: false,
		},
		{
			name: "Empty HTTPProxyConfig",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{},
			},
			want: false,
		},
		{
			name: "HTTPProxyConfig with empty CA",
			args: args{
				httpProxyConfig: &nbcontractv1.HTTPProxyConfig{
					HttpProxy:      "http://fakeproxy.com:8080",
					ProxyTrustedCa: "",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShouldConfigureHTTPProxyCA(tt.args.httpProxyConfig); got != tt.want {
				t.Errorf("getShouldConfigureHTTPProxyCA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTargetEnvironment(t *testing.T) {
	type args struct {
		v *nbcontractv1.Configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Nil CustomCloudConfig",
			args: args{},
			want: nbcontractv1.DefaultCloudName,
		},
		{
			name: "Empty CustomCloudConfig",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{},
				},
			},
			want: nbcontractv1.DefaultCloudName,
		},
		{
			name: "CustomCloudConfig with empty TargetEnvironment",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{},
				},
			},
			want: nbcontractv1.DefaultCloudName,
		},
		{
			name: "CustomCloudConfig with TargetEnvironment",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{
						CustomCloudEnvName: nbcontractv1.AksCustomCloudName,
					},
				},
			},
			want: nbcontractv1.AksCustomCloudName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTargetEnvironment(tt.args.v); got != tt.want {
				t.Errorf("getTargetEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTargetCloud(t *testing.T) {
	type args struct {
		v *nbcontractv1.Configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Nil CustomCloudConfig",
			args: args{},
			want: nbcontractv1.DefaultCloudName,
		},
		{
			name: "Empty CustomCloudConfig",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{},
				},
			},
			want: nbcontractv1.DefaultCloudName,
		},
		{
			name: "CustomCloudConfig with TargetEnvironment",
			args: args{
				v: &nbcontractv1.Configuration{
					CustomCloudConfig: &nbcontractv1.CustomCloudConfig{
						CustomCloudEnvName: nbcontractv1.AksCustomCloudName,
					},
				},
			},
			want: nbcontractv1.AzureStackCloud,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTargetCloud(tt.args.v); got != tt.want {
				t.Errorf("getTargetCloud() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLinuxAdminUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty username",
			args: args{
				username: "",
			},
			want: nbcontractv1.DefaultLinuxUser,
		},
		{
			name: "Non-empty username",
			args: args{
				username: "testuser",
			},
			want: "testuser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLinuxAdminUsername(tt.args.username); got != tt.want {
				t.Errorf("getLinuxAdminUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIsSgxEnabledSKU(t *testing.T) {
	type args struct {
		vmSize string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty vmSize",
			args: args{
				vmSize: "",
			},
			want: false,
		},
		{
			name: "Standard_D2s_v3",
			args: args{
				vmSize: "Standard_D2s_v3",
			},
			want: false,
		},
		{
			name: nbcontractv1.VMSizeStandardDc2s,
			args: args{
				vmSize: nbcontractv1.VMSizeStandardDc2s,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIsSgxEnabledSKU(tt.args.vmSize); got != tt.want {
				t.Errorf("getIsSgxEnabledSKU() = %v, want %v", got, tt.want)
			}
		})
	}
}
