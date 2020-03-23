package UtilsHelp

import (
	"github.com/Cc360428/HelpPackage/UtilsHelp/logs"
	"testing"
)

func TestEmail(t *testing.T) {
	email := "li_chao_cheng@163.com"
	gotB := Email(email);
	logs.Info(gotB)
}

func TestPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantBools bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBools := Password(tt.args.password); gotBools != tt.wantBools {
				t.Errorf("Password() = %v, want %v", gotBools, tt.wantBools)
			}
		})
	}
}

func TestPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name      string
		args      args
		wantBools bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBools := Phone(tt.args.phone); gotBools != tt.wantBools {
				t.Errorf("Phone() = %v, want %v", gotBools, tt.wantBools)
			}
		})
	}
}
