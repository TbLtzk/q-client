package governance

import (
	"testing"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/hexutil"
)

var (
	activeExclusionList = &common.ValidatorExclusionList{
		Timestamp: 1680255000,
		Hash:      common.HexToHash("0xf6cc800a504eadaabe2e7dbd5651fdd546d8769323af8ab50245b61e80795982"),
		Validators: []common.ExcludedValidator{
			{
				Address: common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
				Block:   100,
			},
			{
				Address:  common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
				Block:    100,
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
	}
	//Has timestamp with a less value than activeExclusionList
	obsoleteExclusionList = &common.ValidatorExclusionList{
		Timestamp: 1680250000,
		Hash:      common.HexToHash("0x7d32d21a17e09b9d5affd0aa3b28b7f0e767117498a4163323d3f588dae34a99"),
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
	}
	correctExclusionList = &common.ValidatorExclusionList{
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
	}
	exclusionListWithWrongHash = &common.ValidatorExclusionList{
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
	}
	exclusionListWithDuplicateAddresses = &common.ValidatorExclusionList{
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
	}
	exclusionListWithWrongSignature = &common.ValidatorExclusionList{
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
		Signatures: []hexutil.Bytes{hexutil.MustDecode("0x6662d75a2ddb2c63342dcec4ae307dca29adf5aefd5b06a4586a4541afddcdf4")},
	}
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
				list: correctExclusionList,
			},
			wantErr: false,
		}, {
			name: "Wrong hash",
			args: args{
				list: exclusionListWithWrongHash,
			},
			wantErr: false, // Hash mismatch now logs a warning instead of returning an error
		}, {
			name: "Duplicate address",
			args: args{
				list: exclusionListWithDuplicateAddresses,
			},
			wantErr: true,
		}, {
			name: "Wrong signature",
			args: args{
				list: exclusionListWithWrongSignature,
			},
			wantErr: false,
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
