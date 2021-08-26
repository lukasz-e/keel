package defaults

import (
    kingpin "gopkg.in/alecthomas/kingpin.v2"
    "github.com/rusenask/cron"
    "fmt"
    "strconv"
)

const (

    // ApprovalDeadlineDefault - default deadline in hours
    ApprovalDeadlineDefault = 24
    
    // PollDefaultSchedule - default polling schedule
    PollScheduleDefault = "@every 1m"
)

var PollSchedule *string

var MinimumApprovals *int

var ApprovalDeadline *int

func validatePollSchedule(*kingpin.ParseContext) error {
    // check if provided poll schedule is valid
    _, err := cron.Parse(*PollSchedule)
    if err != nil {
        err = fmt.Errorf("invalid format for default poll schedule: %w", err);
    }
    return err
}

// Registers command line flags and environment variables 
// setting default keel parameter values.
func init() {
    ApprovalDeadline =
        kingpin.Flag("default-approval-deadline", "default approval deadline in hours (24h if not specified)").
            Envar("KEEL_DEFAULT_APPROVAL_DEADLINE").Default(strconv.Itoa(ApprovalDeadlineDefault)).Int()
    MinimumApprovals =
        kingpin.Flag("default-minimum-approvals", "default required approvals count (none if not specified)").
            Envar("KEEL_DEFAULT_MINIMUM_APPROVALS").Int()
    PollSchedule =
        kingpin.Flag("default-poll-schedule", "default poll schedule (none if not specified)").
            Envar("KEEL_DEFAULT_POLL_SCHEDULE").Default(PollScheduleDefault).Action(validatePollSchedule).String()
}

