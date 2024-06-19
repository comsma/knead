package bun

import (
	"bytes"
	"github.com/comsma/knead/pkg/domain"
	"testing"
)

func TestGenerator_WriteFile(t *testing.T) {
	type args struct {
		table *domain.Table
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Generator{}
			w := &bytes.Buffer{}
			err := g.WriteFile(w, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteFile() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_newBunTagForColumn(t *testing.T) {
	type args struct {
		col *domain.Column
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Has Primary Key",
			args: args{
				col: &domain.Column{
					Name: "id",
					Options: domain.ColumnOptions{
						IsPrimaryKey: true,
						DatabaseType: "varchar(30)",
					},
				},
			},
			want: "id,type:varchar(30),pk",
		},
		{
			name: "Has Identity",
			args: args{
				col: &domain.Column{
					Name: "id",
					Options: domain.ColumnOptions{
						IsIdentity:   true,
						DatabaseType: "decimal(19,0)",
					},
				},
			},
			want: "id,type:decimal(19,0),autoincrement",
		},
		{
			name: "Has Default Value",
			args: args{
				col: &domain.Column{
					Name: "is_visible",
					Options: domain.ColumnOptions{
						DefaultValue: "('N')",
						DatabaseType: "varchar",
					},
				},
			},
			want: "is_visible,type:varchar,default:('N')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBunTagForColumn(tt.args.col); got != tt.want {
				t.Errorf("newBunTagForColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
