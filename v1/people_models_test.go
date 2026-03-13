package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestPersonUnmarshalIncludesDocumentedTaskDetails(t *testing.T) {
	const payload = `{
		"id": "65e1efde08e8478f143a8ff9",
		"emailAddress": "example-person@email.com",
		"employment": {
			"endDate": null,
			"jobTitle": "Customer success manager",
			"startDate": "2021-01-01T00:00:00.000Z",
			"status": "CURRENT"
		},
		"leaveInfo": {
			"startDate": "2026-01-01T00:00:00.000Z",
			"endDate": "2026-02-01T00:00:00.000Z"
		},
		"groupIds": ["5f2c939a52855e725c8d5824"],
		"name": {
			"display": "Example Person",
			"last": "Person",
			"first": "Example"
		},
		"sources": {
			"emailAddress": {
				"integrationId": "gsuiteadmin",
				"resourceId": "660c701d3d344e660b032306",
				"type": "INTEGRATION"
			},
			"employment": {
				"startDate": {
					"integrationId": "gusto",
					"resourceId": "660c70783d344e660b032323",
					"type": "INTEGRATION"
				},
				"endDate": {
					"integrationId": "gusto",
					"resourceId": "660c70783d344e660b032323",
					"type": "INTEGRATION"
				}
			}
		},
		"tasksSummary": {
			"completionDate": null,
			"dueDate": "2021-12-01T00:00:00.000Z",
			"status": "OVERDUE",
			"details": {
				"completeTrainings": {
					"taskType": "COMPLETE_TRAININGS",
					"status": "COMPLETE",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": "2021-11-01T00:00:00.000Z",
					"disabled": {
						"date": "2021-11-01T00:00:00.000Z",
						"reason": "Training Vanta tests have been disabled for this person"
					},
					"incompleteTrainings": [
						{"name": "Security training 1"}
					],
					"completedTrainings": [
						{"name": "Security training 2"}
					]
				},
				"acceptPolicies": {
					"taskType": "ACCEPT_POLICIES",
					"status": "COMPLETE",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": "2021-11-01T00:00:00.000Z",
					"disabled": null,
					"unacceptedPolicies": [
						{"name": "Policy 1"}
					],
					"acceptedPolicies": [
						{"name": "Policy 2"}
					]
				},
				"completeCustomTasks": {
					"taskType": "COMPLETE_CUSTOM_TASKS",
					"status": "OVERDUE",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": "2021-11-01T00:00:00.000Z",
					"disabled": {
						"date": "2021-11-01T00:00:00.000Z",
						"reason": "Custom task Vanta tests have been disabled for this person"
					},
					"incompleteCustomTasks": [
						{"name": "Custom task 1"},
						{"name": "Custom task 2"}
					],
					"completedCustomTasks": [
						{"name": "Custom task 3"},
						{"name": "Custom task 4"}
					]
				},
				"completeOffboardingCustomTasks": {
					"taskType": "COMPLETE_CUSTOM_OFFBOARDING_TASKS",
					"status": "COMPLETE",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": "2021-11-01T00:00:00.000Z",
					"disabled": null,
					"incompleteCustomOffboardingTasks": [],
					"completedCustomOffboardingTasks": [
						{"name": "Custom offboarding task 1"},
						{"name": "Custom offboarding task 2"}
					]
				},
				"installDeviceMonitoring": {
					"taskType": "INSTALL_DEVICE_MONITORING",
					"status": "DUE_SOON",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": null,
					"disabled": null
				},
				"completeBackgroundChecks": {
					"taskType": "COMPLETE_BACKGROUND_CHECKS",
					"status": "COMPLETE",
					"dueDate": "2021-12-01T00:00:00.000Z",
					"completionDate": "2021-11-01T00:00:00.000Z",
					"disabled": null
				}
			}
		}
	}`

	var person Person
	if err := json.Unmarshal([]byte(payload), &person); err != nil {
		t.Fatalf("unmarshal person: %v", err)
	}

	if person.LeaveInfo == nil || person.LeaveInfo.StartDate != "2026-01-01T00:00:00.000Z" {
		t.Fatalf("leaveInfo not decoded: %+v", person.LeaveInfo)
	}
	if person.Sources.EmailAddress == nil || person.Sources.EmailAddress.IntegrationID != "gsuiteadmin" {
		t.Fatalf("email source not decoded: %+v", person.Sources.EmailAddress)
	}
	if person.Sources.Employment == nil || person.Sources.Employment.StartDate == nil || person.Sources.Employment.StartDate.IntegrationID != "gusto" {
		t.Fatalf("employment sources not decoded: %+v", person.Sources.Employment)
	}
	if person.TasksSummary == nil {
		t.Fatal("tasksSummary not decoded")
	}

	customTasks := person.TasksSummary.Details.CompleteCustomTasks
	if customTasks == nil {
		t.Fatal("completeCustomTasks not decoded")
	}
	if got := len(customTasks.IncompleteCustomTasks); got != 2 {
		t.Fatalf("incompleteCustomTasks len = %d, want 2", got)
	}
	if got := len(customTasks.CompletedCustomTasks); got != 2 {
		t.Fatalf("completedCustomTasks len = %d, want 2", got)
	}
	if customTasks.Disabled == nil || customTasks.Disabled.Reason == "" {
		t.Fatalf("completeCustomTasks disabled block not decoded: %+v", customTasks.Disabled)
	}

	offboardingTasks := person.TasksSummary.Details.CompleteOffboardingCustomTasks
	if offboardingTasks == nil {
		t.Fatal("completeOffboardingCustomTasks not decoded")
	}
	if got := len(offboardingTasks.CompletedCustomOffboardingTasks); got != 2 {
		t.Fatalf("completedCustomOffboardingTasks len = %d, want 2", got)
	}
	if offboardingTasks.IncompleteCustomOffboardingTasks == nil {
		t.Fatal("incompleteCustomOffboardingTasks should decode as an empty slice, not nil")
	}
}

func TestPeopleBulkMutationRequestsAndResponses(t *testing.T) {
	t.Run("mark as people", func(t *testing.T) {
		var gotBody map[string]any
		httpClient := &http.Client{
			Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Method != http.MethodPost {
					t.Fatalf("method = %s, want POST", r.Method)
				}
				if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
					t.Fatalf("decode request body: %v", err)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     "200 OK",
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       io.NopCloser(strings.NewReader(`{"results":[{"id":"65e1efde08e8478f143a8ff9","status":"SUCCESS"}]}`)),
				}, nil
			}),
		}

		client, err := NewClient(WithHTTPClient(httpClient))
		if err != nil {
			t.Fatalf("NewClient returned error: %v", err)
		}

		resp, err := client.Services.People.MarkAsPeople(context.Background(), &PeopleMarkAsPeopleParams{
			Body: &PeopleMarkAsPeopleRequestBody{
				Updates: []PeopleMarkAsPeopleUpdate{{
					ID: "65e1efde08e8478f143a8ff9",
				}},
			},
		})
		if err != nil {
			t.Fatalf("MarkAsPeople returned error: %v", err)
		}

		wantBody := map[string]any{
			"updates": []any{
				map[string]any{
					"id": "65e1efde08e8478f143a8ff9",
				},
			},
		}
		if !reflect.DeepEqual(gotBody, wantBody) {
			t.Fatalf("request body = %#v, want %#v", gotBody, wantBody)
		}
		if len(resp.Results) != 1 || resp.Results[0].Status != "SUCCESS" {
			t.Fatalf("unexpected mark-as-people response: %+v", resp.Results)
		}
	})

	t.Run("mark as not people", func(t *testing.T) {
		var gotBody map[string]any
		httpClient := &http.Client{
			Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Method != http.MethodPost {
					t.Fatalf("method = %s, want POST", r.Method)
				}
				if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
					t.Fatalf("decode request body: %v", err)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     "200 OK",
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body: io.NopCloser(strings.NewReader(
						`{"results":[{"id":"65e1efde08e8478f143a8ff9","status":"SUCCESS"},{"id":"OTHER_USER_ID","status":"ERROR","message":"Invalid Input"}]}`,
					)),
				}, nil
			}),
		}

		client, err := NewClient(WithHTTPClient(httpClient))
		if err != nil {
			t.Fatalf("NewClient returned error: %v", err)
		}

		resp, err := client.Services.People.MarkAsNotPeople(context.Background(), &PeopleMarkAsNotPeopleParams{
			Body: &PeopleMarkAsNotPeopleRequestBody{
				Updates: []PeopleMarkAsNotPeopleUpdate{{
					ID:     "65e1efde08e8478f143a8ff9",
					Reason: "duplicate account",
				}},
			},
		})
		if err != nil {
			t.Fatalf("MarkAsNotPeople returned error: %v", err)
		}

		wantBody := map[string]any{
			"updates": []any{
				map[string]any{
					"id":     "65e1efde08e8478f143a8ff9",
					"reason": "duplicate account",
				},
			},
		}
		if !reflect.DeepEqual(gotBody, wantBody) {
			t.Fatalf("request body = %#v, want %#v", gotBody, wantBody)
		}
		if len(resp.Results) != 2 {
			t.Fatalf("results len = %d, want 2", len(resp.Results))
		}
		if resp.Results[1].Message == nil || *resp.Results[1].Message != "Invalid Input" {
			t.Fatalf("response message not decoded: %+v", resp.Results[1])
		}
	})

	t.Run("offboard people", func(t *testing.T) {
		var gotBody map[string]any
		httpClient := &http.Client{
			Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Method != http.MethodPost {
					t.Fatalf("method = %s, want POST", r.Method)
				}
				if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
					t.Fatalf("decode request body: %v", err)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     "200 OK",
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       io.NopCloser(strings.NewReader(`{"results":[{"id":"65e1efde08e8478f143a8ff9","status":"SUCCESS"}]}`)),
				}, nil
			}),
		}

		client, err := NewClient(WithHTTPClient(httpClient))
		if err != nil {
			t.Fatalf("NewClient returned error: %v", err)
		}

		resp, err := client.Services.People.OffboardPeople(context.Background(), &PeopleOffboardPeopleParams{
			Body: &PeopleOffboardPeopleRequestBody{
				Updates: []PeopleOffboardPeopleUpdate{{
					ID:             "65e1efde08e8478f143a8ff9",
					AcknowledgerID: "ack-user-id",
				}},
			},
		})
		if err != nil {
			t.Fatalf("OffboardPeople returned error: %v", err)
		}

		wantBody := map[string]any{
			"updates": []any{
				map[string]any{
					"id":             "65e1efde08e8478f143a8ff9",
					"acknowledgerId": "ack-user-id",
				},
			},
		}
		if !reflect.DeepEqual(gotBody, wantBody) {
			t.Fatalf("request body = %#v, want %#v", gotBody, wantBody)
		}
		if len(resp.Results) != 1 || resp.Results[0].Status != "SUCCESS" {
			t.Fatalf("unexpected offboard response: %+v", resp.Results)
		}
	})
}

func TestPeopleUpdatePersonMetadataRequestBodyOmitsUnsetFields(t *testing.T) {
	first := "Example"
	startDate := "1999-02-06T01:34:20.878Z"

	body, err := json.Marshal(PeopleUpdatePersonMetadataRequestBody{
		Name: &PeopleUpdatePersonMetadataName{
			First: &first,
		},
		Employment: &PeopleUpdatePersonMetadataEmployment{
			StartDate: &startDate,
		},
	})
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}

	want := map[string]any{
		"name": map[string]any{
			"first": "Example",
		},
		"employment": map[string]any{
			"startDate": "1999-02-06T01:34:20.878Z",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request body = %#v, want %#v", got, want)
	}
}

func TestUnknownFieldWarningsForPeopleModels(t *testing.T) {
	resetUnknownFieldWarningsForTest()
	t.Cleanup(resetUnknownFieldWarningsForTest)

	var warnings []string
	previous := UnknownFieldWarningf
	UnknownFieldWarningf = func(format string, args ...any) {
		warnings = append(warnings, fmt.Sprintf(format, args...))
	}
	t.Cleanup(func() {
		UnknownFieldWarningf = previous
	})

	const payload = `{
		"results": {
			"data": [
				{
					"id": "65e1efde08e8478f143a8ff9",
					"emailAddress": "example-person@email.com",
					"employment": {
						"endDate": null,
						"jobTitle": "Customer success manager",
						"startDate": "2021-01-01T00:00:00.000Z",
						"status": "CURRENT",
						"unexpectedEmploymentField": true
					},
					"leaveInfo": null,
					"groupIds": [],
					"name": {
						"display": "Example Person",
						"last": "Person",
						"first": "Example"
					},
					"sources": {
						"emailAddress": {
							"integrationId": "gsuiteadmin",
							"resourceId": "660c701d3d344e660b032306",
							"type": "INTEGRATION"
						}
					},
					"tasksSummary": {
						"completionDate": null,
						"dueDate": "2021-12-01T00:00:00.000Z",
						"status": "OVERDUE",
						"details": {
							"completeCustomTasks": {
								"taskType": "COMPLETE_CUSTOM_TASKS",
								"status": "OVERDUE",
								"dueDate": "2021-12-01T00:00:00.000Z",
								"completionDate": "2021-11-01T00:00:00.000Z",
								"disabled": null,
								"incompleteCustomTasks": [],
								"completedCustomTasks": [],
								"unexpectedTaskField": "new"
							},
							"newTaskType": {
								"taskType": "NEW_TASK_TYPE",
								"status": "DUE_SOON"
							}
						}
					},
					"unexpectedPersonField": "new"
				}
			],
			"pageInfo": {
				"hasNextPage": false,
				"hasPreviousPage": false,
				"startCursor": "",
				"endCursor": ""
			}
		}
	}`

	var resp PeopleListPeopleResponse
	if err := decodeJSONBytes([]byte(payload), &resp); err != nil {
		t.Fatalf("unmarshal list people response: %v", err)
	}
	if len(resp.Results.Data) != 1 {
		t.Fatalf("data len = %d, want 1", len(resp.Results.Data))
	}

	joined := strings.Join(warnings, "\n")
	for _, want := range []string{
		"PeopleListPeopleResponse.results.data[].unexpectedPersonField",
		"PeopleListPeopleResponse.results.data[].employment.unexpectedEmploymentField",
		"PeopleListPeopleResponse.results.data[].tasksSummary.details.newTaskType",
		"PeopleListPeopleResponse.results.data[].tasksSummary.details.completeCustomTasks.unexpectedTaskField",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("warnings %q do not contain %q", joined, want)
		}
	}
}
