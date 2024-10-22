package main

import (
	"reflect"
	"testing"
)

func TestGetFont(t *testing.T) {
	type args struct {
		fontName string
	}
	tests := []struct {
		name    string
		args    args
		want    Font
		wantErr bool
	}{
		{"Test 1", args{fontName: "not here font"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFont(tt.args.fontName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFont() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFont() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLetter(t *testing.T) {
	type args struct {
		letter rune
		font   Font
	}
	tests := []struct {
		name    string
		args    args
		want    Letter
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test 1", args{letter: 'ا', font: nil}, nil, true},
		{"Test 2", args{letter: 'ن', font: nil}, nil, true},
		{"Test 3", args{letter: '␡', font: nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLetter(tt.args.letter, tt.args.font)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLetter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
