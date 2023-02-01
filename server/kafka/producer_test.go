package kafka

import (
	"reflect"
	"testing"
	"time"
	"yyds-pro/config"
)

func TestCreateProducer(t *testing.T) {
	type args struct {
		address     string
		topic       string
		duration    time.Duration
		syncOrAsync string
	}
	tests := []struct {
		name string
		args args
		want AbsKafkaProducer
	}{
		// TODO: Add test cases.
		{
			name: "xxxx",
			args: args{
				address:     "159.75.35.70:9092",
				topic:       "test1",
				duration:    5,
				syncOrAsync: Sync,
			},
			want: &SyncKafkaProducer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CreateProducer(tt.args.address, tt.args.topic, tt.args.duration, tt.args.syncOrAsync)
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("CreateProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncKafkaProducer_Send(t *testing.T) {
	con, _ := config.TestKafkaLoadConfig()
	type fields struct {
		KafkaConfig *KafkaConfig
	}
	type args struct {
		value []byte
	}
	conf, _ := NewProducer(con.App.Kafka.Address+":"+con.App.Kafka.Port, "test1", 5)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				KafkaConfig: conf,
			},
			args: args{
				value: []byte("sssss"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &SyncKafkaProducer{
				KafkaConfig: tt.fields.KafkaConfig,
			}
			k.Send(tt.args.value)
		})
	}
}
