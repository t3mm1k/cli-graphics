package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuffer_Clear(t *testing.T) {
	type fields struct {
		W int
		H int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "TestBuffer_Clear",
			fields: fields{
				W: 5,
				H: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := NewBuffer(tt.fields.W, tt.fields.H)
			b.Data[2][2] = NewCellColored('x', ColorBlue, ColorGreen)
			b.Clear()
			assert.Equal(t, b.Data[2][2], NewCell(' '))
		})
	}
}

func TestBuffer_SetSize(t *testing.T) {
	type fields struct {
		W int
		H int
	}
	type args struct {
		w int
		h int
	}
	type want struct {
		w, h int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "IncreaseSize",
			fields: fields{
				W: 1,
				H: 1,
			},
			args: args{
				w: 2,
				h: 2,
			},
			want: want{
				w: 2,
				h: 2,
			},
			wantErr: false,
		},
		{
			name: "DecreaseSize",
			fields: fields{
				W: 3,
				H: 3,
			},
			args: args{
				w: 2,
				h: 2,
			},
			want: want{
				w: 2,
				h: 2,
			},
			wantErr: false,
		},
		{
			name: "Panic on invalid dimensions",
			fields: fields{
				W: 3,
				H: 3,
			},
			args: args{
				w: 0,
				h: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewBuffer(tt.fields.W, tt.fields.H)
			require.NoError(t, err)

			resizeFunc := func() {
				b.SetSize(tt.args.w, tt.args.h)
			}

			if tt.wantErr {
				require.Panics(t, resizeFunc)
				return
			}
			wantB, _ := NewBuffer(tt.want.w, tt.want.h)

			require.NotPanics(t, resizeFunc)

			assert.Equal(t, wantB, b)

			b.SetSize(tt.args.w, tt.args.h)
		})
	}
}

func TestNewBuffer(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name    string
		args    args
		want    *Buffer
		wantErr bool
	}{
		{
			name: "ValidSize",
			args: args{
				width:  3,
				height: 3,
			},
			want: &Buffer{
				Data: [][]Cell{
					{NewCell(' '), NewCell(' '), NewCell(' ')},
					{NewCell(' '), NewCell(' '), NewCell(' ')},
					{NewCell(' '), NewCell(' '), NewCell(' ')},
				},
				W: 3,
				H: 3,
			},
			wantErr: false,
		},
		{
			name: "ZeroWidth",
			args: args{
				width:  0,
				height: 3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NegativeHeight",
			args: args{
				width:  3,
				height: -5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ZeroSize",
			args: args{
				width:  0,
				height: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBuffer(tt.args.width, tt.args.height)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
