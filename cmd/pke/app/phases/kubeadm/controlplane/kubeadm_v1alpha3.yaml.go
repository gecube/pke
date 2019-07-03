// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

// kubeadmConfigV1Alpha3Template is a generated function returning the template as a string.
func kubeadmConfigV1Alpha3Template() string {
	var tmpl = "apiVersion: kubeadm.k8s.io/v1alpha3\n" +
		"kind: InitConfiguration\n" +
		"{{ if .APIServerAdvertiseAddress}}\n" +
		"apiEndpoint:\n" +
		"  advertiseAddress: \"{{ .APIServerAdvertiseAddress }}\"\n" +
		"  bindPort: {{ .APIServerBindPort }}{{end}}\n" +
		"nodeRegistration:\n" +
		"  criSocket: \"unix:///run/containerd/containerd.sock\"\n" +
		"  taints:{{ if not .Taints }} []{{end}}{{range .Taints}}\n" +
		"    - key: \"{{.Key}}\"\n" +
		"      value: \"{{.Value}}\"\n" +
		"      effect: \"{{.Effect}}\"{{end}}\n" +
		"  kubeletExtraArgs:\n" +
		"    {{ if .Nodepool }}\n" +
		"    node-labels: \"nodepool.banzaicloud.io/name={{ .Nodepool }}\"{{end}}\n" +
		"    # pod-infra-container-image: {{ .ImageRepository }}/pause:3.1 # only needed by docker\n" +
		"    {{ if .CloudProvider }}\n" +
		"    cloud-provider: \"{{ .CloudProvider }}\"{{end}}\n" +
		"    {{if .KubeletCloudConfig }}cloud-config: \"/etc/kubernetes/{{ .CloudProvider }}.conf\"{{end}}\n" +
		"    read-only-port: \"0\"\n" +
		"    anonymous-auth: \"false\"\n" +
		"    streaming-connection-idle-timeout: \"5m\"\n" +
		"    protect-kernel-defaults: \"true\"\n" +
		"    event-qps: \"0\"\n" +
		"    client-ca-file: \"/etc/kubernetes/pki/ca.crt\"\n" +
		"    feature-gates: \"RotateKubeletServerCertificate=true\"\n" +
		"    rotate-certificates: \"true\"\n" +
		"    tls-cipher-suites: \"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256\"\n" +
		"    authorization-mode: \"Webhook\"\n" +
		"---\n" +
		"apiVersion: kubeadm.k8s.io/v1alpha3\n" +
		"kind: ClusterConfiguration\n" +
		"clusterName: \"{{ .ClusterName }}\"\n" +
		"imageRepository: {{ .ImageRepository }}\n" +
		"unifiedControlPlaneImage: {{ .ImageRepository }}/hyperkube:v{{ .KubernetesVersion }}\n" +
		"networking:\n" +
		"  serviceSubnet: \"{{ .ServiceCIDR }}\"\n" +
		"  podSubnet: \"{{ .PodCIDR }}\"\n" +
		"  dnsDomain: \"cluster.local\"\n" +
		"kubernetesVersion: \"v{{ .KubernetesVersion }}\"\n" +
		"{{ if .ControlPlaneEndpoint }}controlPlaneEndpoint: \"{{ .ControlPlaneEndpoint }}\"{{end}}\n" +
		"certificatesDir: \"/etc/kubernetes/pki\"\n" +
		"{{ if .APIServerCertSANs }}\n" +
		"apiServerCertSANs:\n" +
		"{{range $k, $san := .APIServerCertSANs}}  - \"{{ $san }}\"\n" +
		"{{end}}{{end}}\n" +
		"apiServerExtraArgs:\n" +
		"  # anonymous-auth: \"false\"\n" +
		"  profiling: \"false\"\n" +
		"  enable-admission-plugins: \"AlwaysPullImages,DenyEscalatingExec,EventRateLimit,NodeRestriction,ServiceAccount{{ if .WithPluginPSP }},PodSecurityPolicy{{end}}\"\n" +
		"  disable-admission-plugins: \"\"\n" +
		"  admission-control-config-file: \"{{ .AdmissionConfig }}\"\n" +
		"  audit-log-path: \"{{ .AuditLogDir }}/apiserver.log\"\n" +
		"  audit-log-maxage: \"30\"\n" +
		"  audit-log-maxbackup: \"10\"\n" +
		"  audit-log-maxsize: \"100\"\n" +
		"  {{ if .WithAuditLog }}audit-policy-file: \"{{ .AuditPolicyFile }}\"{{ end }}\n" +
		"  {{ if .EtcdPrefix }}etcd-prefix: \"{{ .EtcdPrefix }}\"{{end}}\n" +
		"  service-account-lookup: \"true\"\n" +
		"  kubelet-certificate-authority: \"{{ .KubeletCertificateAuthority }}\"\n" +
		"  tls-cipher-suites: \"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256\"\n" +
		"  {{ .EncryptionProviderPrefix }}encryption-provider-config: \"/etc/kubernetes/admission-control/encryption-provider-config.yaml\"\n" +
		"  {{ if (and .OIDCIssuerURL .OIDCClientID) }}\n" +
		"  oidc-issuer-url: \"{{ .OIDCIssuerURL }}\"\n" +
		"  oidc-client-id: \"{{ .OIDCClientID }}\"\n" +
		"  oidc-username-claim: \"email\"\n" +
		"  oidc-username-prefix: \"oidc:\"\n" +
		"  oidc-groups-claim: \"groups\"{{end}}\n" +
		"  {{ if .CloudProvider }}\n" +
		"  cloud-provider: \"{{ .CloudProvider }}\"\n" +
		"  cloud-config: /etc/kubernetes/{{ .CloudProvider }}.conf{{end}}\n" +
		"schedulerExtraArgs:\n" +
		"  profiling: \"false\"\n" +
		"apiServerExtraVolumes:\n" +
		"  {{ if .WithAuditLog }}\n" +
		"  - name: audit-log-dir\n" +
		"    hostPath: {{ .AuditLogDir }}\n" +
		"    mountPath: {{ .AuditLogDir }}\n" +
		"    pathType: DirectoryOrCreate\n" +
		"  - name: audit-policy-file\n" +
		"    hostPath: {{ .AuditPolicyFile }}\n" +
		"    mountPath: {{ .AuditPolicyFile }}\n" +
		"    readOnly: true\n" +
		"    pathType: FileOrCreate{{ end }}\n" +
		"  - name: admission-control-config-file\n" +
		"    hostPath: {{ .AdmissionConfig }}\n" +
		"    mountPath: {{ .AdmissionConfig }}\n" +
		"    writable: false\n" +
		"    pathType: File\n" +
		"  - name: admission-control-config-dir\n" +
		"    hostPath: /etc/kubernetes/admission-control/\n" +
		"    mountPath: /etc/kubernetes/admission-control/\n" +
		"    writable: false\n" +
		"    pathType: Directory\n" +
		"  {{ if .CloudProvider }}\n" +
		"  - name: cloud-config\n" +
		"    hostPath: /etc/kubernetes/{{ .CloudProvider }}.conf\n" +
		"    mountPath: /etc/kubernetes/{{ .CloudProvider }}.conf{{end}}\n" +
		"controllerManagerExtraArgs:\n" +
		"  cluster-name: \"{{ .ClusterName }}\"\n" +
		"  profiling: \"false\"\n" +
		"  terminated-pod-gc-threshold: \"10\"\n" +
		"  feature-gates: \"RotateKubeletServerCertificate=true\"\n" +
		"  {{ if .ControllerManagerSigningCA }}cluster-signing-cert-file: {{ .ControllerManagerSigningCA }}{{end}}\n" +
		"  {{ if .CloudProvider }}\n" +
		"  cloud-provider: \"{{ .CloudProvider }}\"\n" +
		"  cloud-config: /etc/kubernetes/{{ .CloudProvider }}.conf\n" +
		"controllerManagerExtraVolumes:\n" +
		"  - name: cloud-config\n" +
		"    hostPath: /etc/kubernetes/{{ .CloudProvider }}.conf\n" +
		"    mountPath: /etc/kubernetes/{{ .CloudProvider }}.conf{{end}}\n" +
		"etcd:\n" +
		"  {{ if .EtcdEndpoints }}\n" +
		"  external:\n" +
		"    endpoints:\n" +
		"    {{range $k, $endpoint := .EtcdEndpoints }}  - \"{{ $endpoint }}\"\n" +
		"    {{end}}\n" +
		"    caFile: {{ .EtcdCAFile }}\n" +
		"    certFile: {{ .EtcdCertFile }}\n" +
		"    keyFile: {{ .EtcdKeyFile }}\n" +
		"  {{else}}\n" +
		"  local:\n" +
		"    extraArgs:\n" +
		"      peer-auto-tls: \"false\"\n" +
		"  {{end}}\n" +
		"---\n" +
		"apiVersion: kubelet.config.k8s.io/v1beta1\n" +
		"kind: KubeletConfiguration\n" +
		"serverTLSBootstrap: true\n" +
		""
	return tmpl
}
