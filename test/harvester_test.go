package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

// getBaseVars returns the minimum required variables common to all test cases.
// Reads HARVESTER_KUBECONFIG and HARVESTER_NAMESPACE from the environment so
// secrets are never hardcoded.
func getBaseVars(t *testing.T) map[string]interface{} {
	t.Helper()

	kconfig := os.Getenv("HARVESTER_KUBECONFIG")
	require.NotEmpty(t, kconfig, "HARVESTER_KUBECONFIG env var must be set")

	namespace := os.Getenv("HARVESTER_NAMESPACE")
	if namespace == "" {
		namespace = "default"
	}

	harvesterNet := os.Getenv("HARVESTER_NETWORK")
	require.NotEmpty(t, harvesterNet, "HARVESTER_NETWORK env var must be set")

	harvesterNetNS := os.Getenv("HARVESTER_NETWORK_NAMESPACE")
	require.NotEmpty(t, harvesterNetNS, "HARVESTER_NETWORK_NAMESPACE env var must be set")


	imageNamespace := os.Getenv("HARVESTER_IMAGE_NAMESPACE")
	if imageNamespace == "" {
		imageNamespace = namespace
	}

	return map[string]interface{}{
		"kconfig":                  kconfig,
		"namespace":                namespace,
		"harvester_net":            harvesterNet,
		"harvester_net_namespace":  harvesterNetNS,
		"image_display_name":       os.Getenv("HARVESTER_IMAGE_DISPLAY_NAME"),
		"image_namespace":          imageNamespace,
		"image_storageclass":       "harvester-longhorn",
		"download_image":           false,
		"vm_count":                 1,
		"vm_cpu":                   "2",
		"vm_memory":                "2Gi",
		"vm_disksize":              "20Gi",
		"vm_disk_auto_delete":      true,
		"cloud_config_user_data":   baseCloudInit(),
		"cloud_config_network_data": "",
		"ssh_user":                 "ubuntu",
		"vm_description":           "",
		"vm_labels":                map[string]string{},
		"create_lb":                false,
	}
}

func baseCloudInit() string {
	return `#cloud-config
package_update: false
packages:
  - qemu-guest-agent
runcmd:
  - [systemctl, enable, qemu-guest-agent.service]
  - [systemctl, start, qemu-guest-agent.service]
`
}

// terraformOptions builds a standard terraform.Options pointed at the module root.
func terraformOptions(t *testing.T, vars map[string]interface{}) *terraform.Options {
	t.Helper()
	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         vars,
		NoColor:      true,
	})
}
