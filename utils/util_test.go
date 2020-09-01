package utils

import (
	"reflect"
	"testing"
)

func TestToInt64(t *testing.T) {
	type args struct {
		value interface{}
	}
	var tests []struct {
		name    string
		args    args
		wantD   int64
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ToInt64(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotD != tt.wantD {
				t.Errorf("ToInt64() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestToInt64V2(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantD   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ToInt64V2(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64V2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotD != tt.wantD {
				t.Errorf("ToInt64V2() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestSalt(t *testing.T) {
	tests := []struct {
		name     string
		wantSalt string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSalt, err := Salt()
			if (err != nil) != tt.wantErr {
				t.Errorf("Salt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSalt != tt.wantSalt {
				t.Errorf("Salt() = %v, want %v", gotSalt, tt.wantSalt)
			}
		})
	}
}

func TestStringTurnInt(t *testing.T) {
	type args struct {
		pr string
	}
	tests := []struct {
		name    string
		args    args
		wantR   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := StringTurnInt(tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringTurnInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("StringTurnInt() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestIntTurnString(t *testing.T) {
	type args struct {
		pr int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntTurnString(tt.args.pr); got != tt.want {
				t.Errorf("IntTurnString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64TurnString(t *testing.T) {
	type args struct {
		pr int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64TurnString(tt.args.pr); got != tt.want {
				t.Errorf("Int64TurnString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceTurnString(t *testing.T) {
	type args struct {
		pr interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceTurnString(tt.args.pr); got != tt.want {
				t.Errorf("InterfaceTurnString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapTurnStruct(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantObj interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObj, err := MapTurnStruct(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapTurnStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObj, tt.wantObj) {
				t.Errorf("MapTurnStruct() = %v, want %v", gotObj, tt.wantObj)
			}
		})
	}
}

func TestStructuralTurnMap(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructuralTurnMap(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructuralTurnMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInterfaceToMap(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantDest  map[string]interface{}
		wantIsMap bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDest, gotIsMap := ConvertInterfaceToMap(tt.args.src)
			if !reflect.DeepEqual(gotDest, tt.wantDest) {
				t.Errorf("ConvertInterfaceToMap() gotDest = %v, want %v", gotDest, tt.wantDest)
			}
			if gotIsMap != tt.wantIsMap {
				t.Errorf("ConvertInterfaceToMap() gotIsMap = %v, want %v", gotIsMap, tt.wantIsMap)
			}
		})
	}
}
