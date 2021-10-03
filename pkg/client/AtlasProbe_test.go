package client

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
)

type FakeHTTPClient struct{}

func (h FakeHTTPClient) Get(url string) (resp *http.Response, err error) {
	ap := atlasProbe{
		AddressV4:   "1.2.3.4",
		Description: "Fake response.",
	}
	apJSON, err := json.Marshal(ap)
	if err != nil {
		return nil, err
	}
	return &http.Response{Status: "200 OK", Body: io.NopCloser(strings.NewReader(string(apJSON)))}, nil
}

func TestAtlasClient_getIPFromProbe(t *testing.T) {
	type fields struct {
		httpClient IHTTPClient
		addressV4  net.IP
	}
	type args struct {
		probeID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test getIPFromProbe",
			fields: fields{
				httpClient: FakeHTTPClient{},
				addressV4:  nil,
			},
			args: args{
				probeID: "456",
			},
			want:    "1.2.3.4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := AtlasClient{
				httpClient: tt.fields.httpClient,
				addressV4:  &tt.fields.addressV4,
			}
			got, err := c.getIPFromProbe(tt.args.probeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AtlasClient.getIPFromProbe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got != tt.want {
				t.Errorf("AtlasClient.getIPFromProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}
