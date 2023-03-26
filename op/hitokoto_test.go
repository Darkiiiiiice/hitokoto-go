package op

import (
	"reflect"
	"testing"
)

func TestHitokotoUUIDMarkResponse_Parse(t *testing.T) {
	type fields struct {
		Data []int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "array",
			fields: fields{
				Data: []int{1, 2, 3},
			},
			args: args{
				data: []byte(`[1,2,3]`),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &HitokotoUUIDMarkResponse{
				Data: []int{},
			}
			if err := r.Parse(tt.args.data); err != nil {
				t.Errorf("HitokotoUUIDMarkResponse.Parse() error = %v, wantErr %v", err, tt.want)
			}
			if !reflect.DeepEqual(r.Data, tt.want) {
				t.Errorf("HitokotoUUIDMarkResponse.Parse() r.Data = %v, want %v", r.Data, tt.want)
			}
		})
	}
}
