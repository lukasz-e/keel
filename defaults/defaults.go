package defaults

import (
	"fmt"
	"github.com/rusenask/cron"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"strconv"
)

const (

	// ApprovalDeadlineDefault - builtin default deadline in hours
	ApprovalDeadlineDefault = 24

	// PollScheduleDefault - builtin default polling schedule
	PollScheduleDefault = "@every 1m"
)

// global defaults for parameters, calculated from (using first found source):
// - command line flag (--default-<parameter-name>="value", e.g. --default-approval-deadline=24)
// - environment variable (KEEL_DEFAULT_<PARAMETER_NAME>="value, e.g. KEEL_DEFAULT_APPROVAL_DEADLINE=24)
// - builtin default (if there is), or nil (if there is no builtin default)
var (
	// default poll schedule, builtin = "@every 1m"
	PollSchedule *string

	// minimum required approvals, no builtin
	MinimumApprovals *int

	// approval deadline in hours, builtin = 24
	ApprovalDeadline *int
)

func validatePollSchedule(*kingpin.ParseContext) error {
	// check if provided poll schedule is valid
	_, err := cron.Parse(*PollSchedule)
	if err != nil {
		err = fmt.Errorf("invalid format for default poll schedule: %w", err)
	}
	return err
}

// Registers command line flags and environment variables
// used for setting default keel parameter values.
func init() {
	ApprovalDeadline =
		kingpin.Flag("default-approval-deadline", "default approval deadline in hours (24h if not specified)").
			Envar("KEEL_DEFAULT_APPROVAL_DEADLINE").Default(strconv.Itoa(ApprovalDeadlineDefault)).Int()
	MinimumApprovals =
		kingpin.Flag("default-minimum-approvals", "default required approvals count (none if not specified)").
			Envar("KEEL_DEFAULT_MINIMUM_APPROVALS").Int()
	PollSchedule =
		kingpin.Flag("default-poll-schedule", "default poll schedule ('@every 1m' if not specified)").
			Envar("KEEL_DEFAULT_POLL_SCHEDULE").Default(PollScheduleDefault).Action(validatePollSchedule).String()
}
