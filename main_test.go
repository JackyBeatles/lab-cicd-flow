package main

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_rootHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_logRequest(t *testing.T) {
	type args struct {
		handler http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logRequest(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty key",
			args: args{
				key: "",
				fallback: "Nil",
			},
			want: "Nil",
		},
		{
			name: "SHELL key",
			args: args{
				key: "SHELL",
				fallback: "test",
			},
			want: "/bin/zsh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
