package types

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewTinderboxHTTPOptions(t *testing.T) {
	tests := []struct {
		name string
		want *TinderboxHTTPOptions
	}{
		// TODO: Add test cases.
		{
			name: "defaults",
			want: &TinderboxHTTPOptions{
				retries:      3,
				backoffDelay: 1,
				target:       "",
				headers:      http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTinderboxHTTPOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTinderboxHTTPOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinderboxHTTPOptions_WithMaxRetries(t *testing.T) {
	type fields struct {
		retries      int
		backoffDelay int
		target       string
		headers      http.Header
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TinderboxHTTPOptions
	}{
		// TODO: Add test cases.
		{
			name: "adjusted retries",
			fields: fields{
				retries: 3,
			},
			args: args{
				i: 10,
			},
			want: &TinderboxHTTPOptions{
				retries: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TinderboxHTTPOptions{
				retries:      tt.fields.retries,
				backoffDelay: tt.fields.backoffDelay,
				target:       tt.fields.target,
				headers:      tt.fields.headers,
			}
			if got := h.WithMaxRetries(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinderboxHTTPOptions.WithMaxRetries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinderboxHTTPOptions_WithHeader(t *testing.T) {
	type fields struct {
		retries      int
		backoffDelay int
		target       string
		headers      http.Header
	}
	type args struct {
		headerName string
		headerVal  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TinderboxHTTPOptions
	}{
		// TODO: Add test cases.
		{
			name: "adjustable headers",
			fields: fields{
				headers: http.Header{},
			},
			args: args{
				headerName: "User-Agent",
				headerVal:  "Tinderbox/0.1",
			},
			want: &TinderboxHTTPOptions{
				headers: http.Header{
					"User-Agent": []string{"Tinderbox/0.1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TinderboxHTTPOptions{
				retries:      tt.fields.retries,
				backoffDelay: tt.fields.backoffDelay,
				target:       tt.fields.target,
				headers:      tt.fields.headers,
			}
			if got := h.WithHeader(tt.args.headerName, tt.args.headerVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinderboxHTTPOptions.WithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTinderboxHTTPOptions_WithDelay(t *testing.T) {
	type fields struct {
		retries      int
		backoffDelay int
		target       string
		headers      http.Header
	}
	type args struct {
		f int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TinderboxHTTPOptions
	}{
		// TODO: Add test cases.
		{
			name: "adjustable delay",
			fields: fields{
				backoffDelay: 0,
			},
			args: args{
				f: 5,
			},
			want: &TinderboxHTTPOptions{
				backoffDelay: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TinderboxHTTPOptions{
				retries:      tt.fields.retries,
				backoffDelay: tt.fields.backoffDelay,
				target:       tt.fields.target,
				headers:      tt.fields.headers,
			}
			if got := h.WithDelay(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinderboxHTTPOptions.WithDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}
