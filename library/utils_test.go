package library

import (
	"reflect"
	"testing"
)

func TestGetSortedMapValues(t *testing.T) {
	type args struct {
		m map[string]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"GetSortedMapValues",
			args{map[string]float64{"a": 1.2, "c": 9.02, "b": 232.01}},
			[]float64{1.2, 9.02, 232.01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSortedMapValues(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortedMapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInstanceByBNS(t *testing.T) {
	/*
	type args struct {
		bns string
	}
	tests := []struct {
		name        string
		args        args
		wantBnsList []models.BNSItem
		wantErr     bool
	}{
		{
			"get_instance_by_service",
			args{"group.bdrp-bdrp-common-test-v308-2-proxy.redis.all"},
			[]models.BNSItem{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBnsList, err := GetInstanceByBNS(tt.args.bns)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInstanceByBNS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBnsList, tt.wantBnsList) {
				t.Errorf("GetInstanceByBNS() gotBnsList = %v, want %v", gotBnsList, tt.wantBnsList)
			}
		})
	}
	*/
}

func TestGetInstanceByBNS2(t *testing.T) {
	bns := "group.bdrp-bdrp-common-test-v308-2-proxy.redis.all"
	bnsList, err := GetInstanceByBNS(bns)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(bnsList)
	}
}