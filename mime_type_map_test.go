package mimetype

import (
	"testing"
)

func Test_mappings_GetMimeTypeFromExtension(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name         string
		args         args
		wantMimeType string
		wantOk       bool
	}{
		{
			name: "should return valid mp4 mime type",
			args: args{
				extension: ".mp4",
			},
			wantMimeType: "video/mp4",
			wantOk:       true,
		},
		{
			name: "should return valid webm mime type",
			args: args{
				extension: ".webm",
			},
			wantMimeType: "video/webm",
			wantOk:       true,
		},
		{
			name: "should return valid 3gpp mime type",
			args: args{
				extension: ".3gpp",
			},
			wantMimeType: "video/3gpp",
			wantOk:       true,
		},
		{
			name: "should fail, empty",
			args: args{
				extension: "",
			},
			wantMimeType: "",
			wantOk:       false,
		},
		{
			name: "should fail, missing .",
			args: args{
				extension: "mp4",
			},
			wantMimeType: "",
			wantOk:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMimeType, gotOk := Mappings.GetMimeTypeFromExtension(tt.args.extension)
			if gotMimeType != tt.wantMimeType {
				t.Errorf("Mappings.GetMimeTypeFromExtension() gotMimeType = %v, want %v", gotMimeType, tt.wantMimeType)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Mappings.GetMimeTypeFromExtension() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
