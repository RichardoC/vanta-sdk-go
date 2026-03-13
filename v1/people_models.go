package v1

// PersonTaskType values used for personnel tasks.
const (
	PersonTaskTypeAcceptPolicies                 = "ACCEPT_POLICIES"
	PersonTaskTypeCompleteTrainings              = "COMPLETE_TRAININGS"
	PersonTaskTypeCompleteCustomTasks            = "COMPLETE_CUSTOM_TASKS"
	PersonTaskTypeCompleteCustomOffboardingTasks = "COMPLETE_CUSTOM_OFFBOARDING_TASKS"
	PersonTaskTypeInstallDeviceMonitoring        = "INSTALL_DEVICE_MONITORING"
	PersonTaskTypeCompleteBackgroundChecks       = "COMPLETE_BACKGROUND_CHECKS"
)

// Person represents a person-like record returned from people/group endpoints.
type Person struct {
	ID           string              `json:"id"`
	EmailAddress string              `json:"emailAddress"`
	Employment   PersonEmployment    `json:"employment"`
	LeaveInfo    *PersonLeaveInfo    `json:"leaveInfo"`
	GroupIDs     []string            `json:"groupIds"`
	Name         PersonName          `json:"name"`
	Sources      PersonSources       `json:"sources"`
	TasksSummary *PersonTasksSummary `json:"tasksSummary"`
}

type PersonEmployment struct {
	EndDate   *string `json:"endDate"`
	JobTitle  string  `json:"jobTitle"`
	StartDate string  `json:"startDate"`
	Status    string  `json:"status"`
}

type PersonName struct {
	Display string `json:"display"`
	Last    string `json:"last"`
	First   string `json:"first"`
}

type PersonLeaveInfo struct {
	EndDate   string `json:"endDate"`
	StartDate string `json:"startDate"`
}

type PersonSources struct {
	EmailAddress *PersonSourceRef         `json:"emailAddress"`
	Employment   *PersonEmploymentSources `json:"employment"`
}

type PersonEmploymentSources struct {
	EndDate   *PersonSourceRef `json:"endDate"`
	JobTitle  *PersonSourceRef `json:"jobTitle"`
	StartDate *PersonSourceRef `json:"startDate"`
	Status    *PersonSourceRef `json:"status"`
}

type PersonSourceRef struct {
	IntegrationID string `json:"integrationId"`
	ResourceID    string `json:"resourceId"`
	Type          string `json:"type"`
}

// PersonTasksSummary provides completion and status details for people tasks.
type PersonTasksSummary struct {
	CompletionDate *string           `json:"completionDate"`
	DueDate        *string           `json:"dueDate"`
	Status         string            `json:"status"`
	Details        PersonTaskDetails `json:"details"`
}

// PersonTaskDetail describes a specific personnel task instance.
type PersonTaskDetail struct {
	TaskType                         string                 `json:"taskType"`
	Status                           string                 `json:"status"`
	DueDate                          *string                `json:"dueDate"`
	CompletionDate                   *string                `json:"completionDate"`
	Disabled                         *PersonTaskDisabled    `json:"disabled"`
	IncompleteTrainings              []PersonNamedReference `json:"incompleteTrainings"`
	CompletedTrainings               []PersonNamedReference `json:"completedTrainings"`
	UnacceptedPolicies               []PersonNamedReference `json:"unacceptedPolicies"`
	AcceptedPolicies                 []PersonNamedReference `json:"acceptedPolicies"`
	IncompleteCustomTasks            []PersonNamedReference `json:"incompleteCustomTasks"`
	CompletedCustomTasks             []PersonNamedReference `json:"completedCustomTasks"`
	IncompleteCustomOffboardingTasks []PersonNamedReference `json:"incompleteCustomOffboardingTasks"`
	CompletedCustomOffboardingTasks  []PersonNamedReference `json:"completedCustomOffboardingTasks"`
}

type PersonTaskDetails struct {
	AcceptPolicies                 *PersonTaskDetail `json:"acceptPolicies"`
	CompleteBackgroundChecks       *PersonTaskDetail `json:"completeBackgroundChecks"`
	CompleteCustomTasks            *PersonTaskDetail `json:"completeCustomTasks"`
	CompleteOffboardingCustomTasks *PersonTaskDetail `json:"completeOffboardingCustomTasks"`
	CompleteTrainings              *PersonTaskDetail `json:"completeTrainings"`
	InstallDeviceMonitoring        *PersonTaskDetail `json:"installDeviceMonitoring"`
}

type PersonTaskDisabled struct {
	Date   string `json:"date"`
	Reason string `json:"reason"`
}

type PersonNamedReference struct {
	Name string `json:"name"`
}

// PeopleListPeopleResponse is the typed response envelope for listing people.
type PeopleListPeopleResponse struct {
	Results PeopleListPeopleResults `json:"results"`
}

type PeopleListPeopleResults struct {
	Data     []Person `json:"data"`
	PageInfo PageInfo `json:"pageInfo"`
}
