package governance

import (
	"testing"

	"gitlab.com/q-dev/q-client/common"
)

func Test_newExclusionSet(t *testing.T) {
	type args struct {
		list *common.ValidatorExclusionList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Correct exclusion set",
			args: args{
				list: &common.ValidatorExclusionList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0x25c2d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4"),
					Validators: []common.ExcludedValidator{
						{
							Address: common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
							Block:   0,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    0,
							EndBlock: 1000,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    2000,
							EndBlock: 3000,
						},
						{
							Address: common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
							Block:   1000,
						},
					},
				},
			},
			wantErr: false,
		}, {
			name: "Wrong hash",
			args: args{
				list: &common.ValidatorExclusionList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0x6662d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4"),
					Validators: []common.ExcludedValidator{
						{
							Address: common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
							Block:   0,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    0,
							EndBlock: 1000,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    2000,
							EndBlock: 3000,
						},
						{
							Address: common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
							Block:   1000,
						},
					},
				},
			},
			wantErr: true,
		}, {
			name: "Duplicate address",
			args: args{
				list: &common.ValidatorExclusionList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0x25c2d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4"),
					Validators: []common.ExcludedValidator{
						{
							Address: common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
							Block:   0,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    0,
							EndBlock: 1000,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    2000,
							EndBlock: 3000,
						},
						{
							Address: common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
							Block:   1000,
						},
						{
							Address: common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
							Block:   1000,
						},
					},
				},
			},
			wantErr: true,
		}, {
			name: "Wrong signature",
			args: args{
				list: &common.ValidatorExclusionList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0x25c2d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4"),
					Validators: []common.ExcludedValidator{
						{
							Address: common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
							Block:   0,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    0,
							EndBlock: 1000,
						},
						{
							Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
							Block:    2000,
							EndBlock: 3000,
						},
						{
							Address: common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
							Block:   1000,
						},
					},
					Signatures: [][]byte{[]byte("0x6662d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4")},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newExclusionSet(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("newExclusionSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
