package log

import (
	"github.com/SAP/jenkins-library/pkg/ans"
	"github.com/SAP/jenkins-library/pkg/xsuaa"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestANSHook_Levels(t *testing.T) {
	hook := NewANSHook("", "", "")
	assert.Equal(t, []logrus.Level{logrus.InfoLevel, logrus.DebugLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel},
		hook.Levels())
}

func TestNewANSHook(t *testing.T) {
	testClient := ans.ANS{
		XSUAA: xsuaa.XSUAA{
			OAuthURL:     "https://my.test.oauth.provider",
			ClientID:     "myTestClientID",
			ClientSecret: "super secret",
		},
		URL: "https://my.test.backend",
	}
	testServiceKeyJSON := `{
					"url": "https://my.test.backend",
					"client_id": "myTestClientID",
					"client_secret": "super secret",
					"oauth_url": "https://my.test.oauth.provider"
				}`
	type args struct {
		serviceKey    string
		correlationID string
		eventTemplate string
	}
	tests := []struct {
		name string
		args args
		want ANSHook
	}{
		{
			name: "Straight forward test",
			args: args{
				serviceKey:    testServiceKeyJSON,
				correlationID: "1234",
			},
			want: ANSHook{
				correlationID: "1234",
				client:        testClient,
			},
		},
		{
			name: "No service key = no client",
			args: args{
				correlationID: "1234",
			},
			want: ANSHook{
				correlationID: "1234",
				client:        ans.ANS{},
			},
		},
		{
			name: "With event template",
			args: args{
				serviceKey:    testServiceKeyJSON,
				correlationID: "1234",
				eventTemplate: `{"priority":123}`,
			},
			want: ANSHook{
				correlationID: "1234",
				client:        testClient,
				event: ans.Event{
					Priority: 123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewANSHook(tt.args.serviceKey, tt.args.correlationID, tt.args.eventTemplate)
			assert.Equal(t, tt.want, got, "new ANSHook not as expected")
		})
	}
}

func TestANSHook_Fire(t *testing.T) {
	testClient := ansMock{}
	type fields struct {
		correlationID string
		client        ansMock
		levels        []logrus.Level
		event         ans.Event
	}
	tests := []struct {
		name      string
		fields    fields
		entryArg  *logrus.Entry
		wantEvent ans.Event
	}{
		{
			name: "Straight forward test",
			fields: fields{
				correlationID: "1234",
				client:        testClient,
			},
			entryArg: &logrus.Entry{
				Level:   logrus.InfoLevel,
				Time:    time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC),
				Message: "my log message",
				Data:    map[string]interface{}{"stepName": "testStep"},
			},
			wantEvent: ans.Event{
				EventType:      "Piper",
				EventTimestamp: time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC).Unix(),
				Severity:       "INFO",
				Category:       "NOTIFICATION",
				Subject:        "testStep",
				Body:           "my log message",
				Tags:           map[string]interface{}{"ans:correlationId": "1234", "stepName": "testStep", "logLevel": "info"},
			},
		},
		{
			name: "Event already set",
			fields: fields{
				correlationID: "1234",
				client:        testClient,
				event: ans.Event{
					EventType: "My event type",
					Subject:   "My subject line",
					Tags:      map[string]interface{}{"Some": 1, "Additional": "a string", "Tags": true},
				},
			},
			entryArg: &logrus.Entry{
				Level:   logrus.InfoLevel,
				Time:    time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC),
				Message: "my log message",
				Data:    map[string]interface{}{"stepName": "testStep"},
			},
			wantEvent: ans.Event{
				EventType:      "My event type",
				EventTimestamp: time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC).Unix(),
				Severity:       "INFO",
				Category:       "NOTIFICATION",
				Subject:        "My subject line",
				Body:           "my log message",
				Tags:           map[string]interface{}{"ans:correlationId": "1234", "stepName": "testStep", "logLevel": "info", "Some": 1, "Additional": "a string", "Tags": true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansHook := &ANSHook{
				correlationID: tt.fields.correlationID,
				client:        tt.fields.client,
				event:         tt.fields.event,
			}
			defer func() {testEvent = ans.Event{}}()
			ansHook.Fire(tt.entryArg)
			assert.Equal(t, tt.wantEvent, testEvent, "Event is not as expected.")
		})
	}
}

type ansMock struct{}

var testEvent ans.Event

func (ans ansMock) Send(event ans.Event) error {
	testEvent = event
	return nil
}
