package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestPlanNoLB verifies that planning with create_lb=false produces a valid
// plan. This is a plan-only test — no infrastructure is created.
func TestPlanNoLB(t *testing.T) {
	t.Parallel()

	vars := getBaseVars(t)
	vars["vm_prefix"] = "test-nolb"
	vars["create_lb"] = false

	opts := terraformOptions(t, vars)

	exitCode := terraform.InitAndPlanWithExitCode(t, opts)
	assert.Contains(t, []int{0, 2}, exitCode,
		"terraform plan should succeed (exit 0 or 2) when create_lb=false")
}

// TestPlanWithLB verifies that planning with create_lb=true produces a valid
// plan. This is a plan-only test — no infrastructure is created.
func TestPlanWithLB(t *testing.T) {
	t.Parallel()

	vars := getBaseVars(t)
	vars["vm_prefix"] = "test-withlb"
	vars["create_lb"] = true
	vars["lb_ipam"] = "dhcp"
	vars["lb_listener_port"] = 443
	vars["lb_listener_backend_port"] = 443
	vars["lb_protocol"] = "TCP"
	vars["lb_healthcheck_period_seconds"] = 5
	vars["lb_healthcheck_timeout_seconds"] = 3
	vars["lb_healthcheck_failure_threshold"] = 3
	vars["lb_healthcheck_success_threshold"] = 1

	opts := terraformOptions(t, vars)

	exitCode := terraform.InitAndPlanWithExitCode(t, opts)
	assert.Contains(t, []int{0, 2}, exitCode,
		"terraform plan should succeed when create_lb=true")
}