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
			gotMimeType, gotOk := GetMimeTypeFromExtension(tt.args.extension)
			if gotMimeType != tt.wantMimeType {
				t.Errorf("Mappings.GetMimeTypeFromExtension() gotMimeType = %v, want %v", gotMimeType, tt.wantMimeType)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Mappings.GetMimeTypeFromExtension() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestGetExtentionFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return .mp4",
			args: args{
				path: "www.foobar.com/file.mp4",
			},
			want:    ".mp4",
			wantErr: false,
		},
		{
			name: "should return nothing because the extension is missing",
			args: args{
				path: "filemp4",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetExtentionFromPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExtentionFromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetExtentionFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
