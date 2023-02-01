package serviceimpl

import (
	"testing"
	"yyds-pro/repository"
	"yyds-pro/trace"
)

func TestReportService_GenerateReportByKafka(t *testing.T) {
	type fields struct {
		ReportRepo repository.ReportRepo
	}
	type args struct {
		c *trace.Trace
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantErr  bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ReportService{
				ReportRepo: tt.fields.ReportRepo,
			}
			gotCode, err := s.GenerateReportByKafka(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateReportByKafka() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("GenerateReportByKafka() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
