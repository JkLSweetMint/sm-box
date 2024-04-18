package closer

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		conf *Config
		ctx  context.Context
		wg   *sync.WaitGroup
	}

	tests := []struct {
		name   string
		args   args
		wantCl Closer
		wantCt context.Context
	}{
		{
			name: "Case 1",
			args: args{
				conf: new(Config).Default(),
				ctx:  context.Background(),
				wg:   new(sync.WaitGroup),
			},
			wantCl: &closer{
				conf:      new(Config).Default(),
				ctx:       nil,
				ctxCancel: nil,
			},
			wantCt: nil,
		},
	}

	for i, tt := range tests {
		tt.wantCl.(*closer).ctx, tt.wantCl.(*closer).ctxCancel = context.WithCancel(tt.args.ctx)
		tt.wantCt = tt.wantCl.(*closer).ctx
		tt.wantCl.(*closer).wg = tt.args.wg

		tests[i] = tt
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCl, gotCt := New(tt.args.conf, tt.args.ctx, tt.args.wg)

			if gotCl.(*closer).ctxCancel == nil || tt.wantCl.(*closer).ctxCancel == nil {
				t.Errorf("New() gotCl.ctxCancel = %v, want.ctxCancel %v", gotCl.(*closer).ctxCancel, tt.wantCl.(*closer).ctxCancel)
			}

			gotCl.(*closer).ctxCancel = nil
			tt.wantCl.(*closer).ctxCancel = nil

			if !reflect.DeepEqual(gotCl, tt.wantCl) {
				t.Errorf("New() gotCl = %v, want %v", gotCl, tt.wantCl)
			}

			if !reflect.DeepEqual(gotCt, tt.wantCt) {
				t.Errorf("New() gotCt = %v, want %v", gotCt, tt.wantCt)
			}
		})
	}
}
