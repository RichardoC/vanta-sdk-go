package v1

// PersonTaskType values used for personnel tasks.
const (
	PersonTaskTypeAcceptPolicies      = "ACCEPT_POLICIES"
	PersonTaskTypeCompleteTrainings   = "COMPLETE_TRAININGS"
	PersonTaskTypeCompleteCustomTasks = "COMPLETE_CUSTOM_TASKS"
)

// Person represents a person-like record returned from people/group endpoints.
type Person struct {
	ID           string             `json:"id"`
	EmailAddress string             `json:"emailAddress"`
	Employment   PersonEmployment   `json:"employment"`
	LeaveInfo    any                `json:"leaveInfo"`
	GroupIDs     []string           `json:"groupIds"`
	Name         PersonName         `json:"name"`
	Sources      map[string]any     `json:"sources"`
	TasksSummary PersonTasksSummary `json:"tasksSummary"`
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

// PersonTasksSummary provides completion and status details for people tasks.
type PersonTasksSummary struct {
	CompletionDate *string                     `json:"completionDate"`
	DueDate        *string                     `json:"dueDate"`
	Status         string                      `json:"status"`
	Details        map[string]PersonTaskDetail `json:"details"`
}

// PersonTaskDetail describes a specific personnel task instance.
type PersonTaskDetail struct {
	TaskType       string  `json:"taskType"`
	Status         string  `json:"status"`
	DueDate        *string `json:"dueDate"`
	CompletionDate *string `json:"completionDate"`
}

// PeopleListPeopleResponse is the typed response envelope for listing people.
type PeopleListPeopleResponse struct {
	Results struct {
		Data     []Person `json:"data"`
		PageInfo PageInfo `json:"pageInfo"`
	} `json:"results"`
}
