package governance

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/q-dev/q-client/common"
)

func Test_newRootSet(t *testing.T) {
	type args struct {
		list *common.RootList
	}
	tests := []struct {
		name    string
		args    args
		want    *rootSet
		wantErr bool
	}{
		{
			name: "Correct root set",
			args: args{
				list: &common.RootList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					Nodes: []common.Address{
						common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					},
				},
			},
			want: &rootSet{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: nil,
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
			},
			wantErr: false,
		}, {
			name: "Wrong hash",
			args: args{
				list: &common.RootList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0x666af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					Nodes: []common.Address{
						common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					},
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Wrong signature",
			args: args{
				list: &common.RootList{
					Timestamp: 1680255617,
					Hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					Nodes: []common.Address{
						common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					},
					Signatures: [][]byte{
						common.Hex2Bytes("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newRootSet(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRootSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !rootSetsAreEqual(t, got, tt.want) {
				t.Errorf("newRootSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func rootSetsAreEqual(t *testing.T, a, b *rootSet) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.timestamp != b.timestamp {
		return false
	}
	if a.hash != b.hash {
		return false
	}
	if !assert.ElementsMatch(t, a.rootAddresses, b.rootAddresses) {
		return false
	}
	if !reflect.DeepEqual(a.aliases, b.aliases) {
		return false
	}
	if !assert.Equal(t, a.roots, b.roots) {
		return false
	}
	return true
}

func Test_rootSet_addSignature(t *testing.T) {
	type fields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	type args struct {
		signer    common.Address
		signature []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Add signature",
			fields: fields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
				},
				aliases: map[common.Address]common.Address{},
				roots:   map[common.Address]struct{}{},
				signers: map[common.Address][]byte{},
			},
			args: args{
				signer:    common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
				signature: []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
			},
			want: true,
		},
		{
			name: "Add signature twice",
			fields: fields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
				},
				aliases: map[common.Address]common.Address{},
				roots:   map[common.Address]struct{}{},
				signers: map[common.Address][]byte{
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
				},
			},
			args: args{
				signer:    common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
				signature: []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
			},
			want: false,
		},
		{
			name: "Add empty signature",
			fields: fields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
				},
				aliases: map[common.Address]common.Address{},
				roots:   map[common.Address]struct{}{},
				signers: map[common.Address][]byte{
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
				},
			},
			args: args{
				signer:    common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
				signature: []byte(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.addSignature(tt.args.signer, tt.args.signature); got != tt.want {
				t.Errorf("addSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO update calcHash to use the all fields
func Test_rootSet_calcHash(t *testing.T) {
	type fields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   common.Hash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.calcHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootSet_copy(t *testing.T) {
	type fields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   *rootSet
	}{
		{
			name: "Copy",
			fields: fields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: map[common.Address]common.Address{},
				roots:   map[common.Address]struct{}{},
				signers: map[common.Address][]byte{},
			},
			want: &rootSet{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: map[common.Address]common.Address{},
				roots:   map[common.Address]struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.copy(); !rootSetsAreEqual(t, got, tt.want) {
				t.Errorf("copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootSet_isAcceptable(t *testing.T) {
	type initialSetFields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	type args struct {
		set *rootSet
	}
	tests := []struct {
		name   string
		fields initialSetFields
		args   args
		want   bool
	}{
		{
			name: "Acceptable",
			fields: initialSetFields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				signers: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
			},
			args: args{
				set: &rootSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					rootAddresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					aliases: map[common.Address]common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					roots: map[common.Address]struct{}{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
					},
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
					},
				},
			},
			want: true,
		}, {
			name: "Not acceptable. Not enough signers",
			fields: initialSetFields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				signers: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
			},
			args: args{
				set: &rootSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					rootAddresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					aliases: map[common.Address]common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					roots: map[common.Address]struct{}{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
					},
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					},
				},
			},
			want: false,
		}, {
			name: "Not acceptable. Alias is not the same",
			fields: initialSetFields{
				timestamp: 1680255617,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x674FCA43B4742EB7E0610C1C64B6D25EDC0C998F"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				signers: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
			},
			args: args{
				set: &rootSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					rootAddresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					aliases: map[common.Address]common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					roots: map[common.Address]struct{}{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
					},
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.isAcceptable(tt.args.set); got != tt.want {
				t.Errorf("isAcceptable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootSet_isEnoughExSetSignatures(t *testing.T) {
	type rootSetFields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	type args struct {
		set *exclusionSet
	}
	tests := []struct {
		name   string
		fields rootSetFields
		args   args
		want   bool
	}{
		{
			name: "Enough signatures",
			fields: rootSetFields{
				timestamp: 1680255618,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
			},
			args: args{
				set: &exclusionSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					addresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					addrToBlock: nil,
					blockRanges: nil,
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					},
				},
			},
			want: true,
		}, {
			name: "Not enough signatures",
			fields: rootSetFields{
				timestamp: 1680255618,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				signers: map[common.Address][]byte{},
			},
			args: args{
				set: &exclusionSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					addresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					addrToBlock: nil,
					blockRanges: nil,
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					},
				},
			},
			want: false,
		}, {
			name: "Aliases differ",
			fields: rootSetFields{
				timestamp: 1680255618,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0xE0B8B589A1DBA9AD4860892913A738FB7F8DEFFF"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				signers: map[common.Address][]byte{},
			},
			args: args{
				set: &exclusionSet{
					timestamp: 1680255618,
					hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
					addresses: []common.Address{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
					},
					addrToBlock: nil,
					blockRanges: nil,
					signers: map[common.Address][]byte{
						common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
						common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
						common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.isEnoughExSetSignatures(tt.args.set); got != tt.want {
				t.Errorf("isEnoughExSetSignatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootSet_knownSigners(t *testing.T) {
	type rootSetFields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	type args struct {
		signers map[common.Address][]byte
	}
	tests := []struct {
		name   string
		fields rootSetFields
		args   args
		want   []common.Address
	}{
		{
			name: "known signers",
			fields: rootSetFields{
				timestamp: 1680255618,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				signers: map[common.Address][]byte{},
			},
			args: args{
				signers: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
				},
			},
			want: []common.Address{
				common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
				common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			assert.ElementsMatch(t, tt.want, s.knownSigners(tt.args.signers))
		})
	}
}

func Test_rootSet_mergeSignatures(t *testing.T) {
	type rootSetFields struct {
		timestamp     uint64
		hash          common.Hash
		rootAddresses []common.Address
		aliases       map[common.Address]common.Address
		roots         map[common.Address]struct{}
		signers       map[common.Address][]byte
	}
	type args struct {
		hash       common.Hash
		signatures map[common.Address][]byte
	}
	tests := []struct {
		name   string
		fields rootSetFields
		args   args
		want   map[common.Address][]byte
	}{
		{
			name: "merge signatures",
			fields: rootSetFields{
				timestamp: 1680255618,
				hash:      common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				rootAddresses: []common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
				},
				roots: map[common.Address]struct{}{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): {},
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): {},
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): {},
				},
				aliases: map[common.Address]common.Address{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
				},
				signers: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
				},
			},
			args: args{
				hash: common.HexToHash("0xf68af076b760d81d3d2a5071817c7def517a0d60f7e5bec5c65daf6e2dcab855"),
				signatures: map[common.Address][]byte{
					common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
					common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
				},
			},
			want: map[common.Address][]byte{
				common.HexToAddress("0XA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"): []byte("GHZqFYSD5bSgkLgJvOX03GJ+lOUrP6p6hBySD2lIfENt9oVIfXrcV4LwQiU0OQ7aI0LPDoM3AsBlCIZsRN4CRAE="),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rootSet{
				timestamp:     tt.fields.timestamp,
				hash:          tt.fields.hash,
				rootAddresses: tt.fields.rootAddresses,
				aliases:       tt.fields.aliases,
				roots:         tt.fields.roots,
				signers:       tt.fields.signers,
			}
			if got := s.mergeSignatures(tt.args.hash, tt.args.signatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSignatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
