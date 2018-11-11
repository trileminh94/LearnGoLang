package queue

import (
	"testing"
)

func Test_send(t *testing.T) {
	type args struct {
		mesg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"send_ping", args{"ping"}},
		{"send_ping", args{"ping1"}},
		{"send_ping", args{"ping2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			send(tt.args.mesg)
		})
	}
}

func Test_receive(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"receive_ping", "ping"},
		{"receive_ping", "ping1"},
		{"receive_ping", "ping2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := receive(); got != tt.want {
				t.Errorf("receive() = %v, want %v", got, tt.want)
			}
		})
	}
}
