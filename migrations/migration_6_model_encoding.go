// Code generated by fastssz. DO NOT EDIT.
// Hash: 84911d88fef697c42b566afc0ab058f03bae7c645b0f24f3634e6953380c3b27
// Version: 0.1.3
package migrations

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the migration_6_OldStorageOperator object
func (m *migration_6_OldStorageOperator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the migration_6_OldStorageOperator object to a target array
func (m *migration_6_OldStorageOperator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Field (0) 'OperatorID'
	dst = ssz.MarshalUint64(dst, m.OperatorID)

	// Offset (1) 'PubKey'
	dst = ssz.WriteOffset(dst, offset)

	// Field (1) 'PubKey'
	if size := len(m.PubKey); size > 48 {
		err = ssz.ErrBytesLengthFn("migration_6_OldStorageOperator.PubKey", size, 48)
		return
	}
	dst = append(dst, m.PubKey...)

	return
}

// UnmarshalSSZ ssz unmarshals the migration_6_OldStorageOperator object
func (m *migration_6_OldStorageOperator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'OperatorID'
	m.OperatorID = ssz.UnmarshallUint64(buf[0:8])

	// Offset (1) 'PubKey'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 != 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'PubKey'
	{
		buf = tail[o1:]
		if len(buf) > 48 {
			return ssz.ErrBytesLength
		}
		if cap(m.PubKey) == 0 {
			m.PubKey = make([]byte, 0, len(buf))
		}
		m.PubKey = append(m.PubKey, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the migration_6_OldStorageOperator object
func (m *migration_6_OldStorageOperator) SizeSSZ() (size int) {
	size = 12

	// Field (1) 'PubKey'
	size += len(m.PubKey)

	return
}

// HashTreeRoot ssz hashes the migration_6_OldStorageOperator object
func (m *migration_6_OldStorageOperator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the migration_6_OldStorageOperator object with a hasher
func (m *migration_6_OldStorageOperator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'OperatorID'
	hh.PutUint64(m.OperatorID)

	// Field (1) 'PubKey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.PubKey))
		if byteLen > 48 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(m.PubKey)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (48+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the migration_6_OldStorageOperator object
func (m *migration_6_OldStorageOperator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the migration_6_OldStorageShare object
func (m *migration_6_OldStorageShare) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the migration_6_OldStorageShare object to a target array
func (m *migration_6_OldStorageShare) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(129)

	// Field (0) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, m.ValidatorIndex)

	// Field (1) 'ValidatorPubKey'
	if size := len(m.ValidatorPubKey); size != 48 {
		err = ssz.ErrBytesLengthFn("migration_6_OldStorageShare.ValidatorPubKey", size, 48)
		return
	}
	dst = append(dst, m.ValidatorPubKey...)

	// Offset (2) 'SharePubKey'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.SharePubKey)

	// Offset (3) 'Committee'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(m.Committee); ii++ {
		offset += 4
		offset += m.Committee[ii].SizeSSZ()
	}

	// Field (4) 'DomainType'
	dst = append(dst, m.DomainType[:]...)

	// Field (5) 'FeeRecipientAddress'
	dst = append(dst, m.FeeRecipientAddress[:]...)

	// Offset (6) 'Graffiti'
	dst = ssz.WriteOffset(dst, offset)

	// Field (7) 'Status'
	dst = ssz.MarshalUint64(dst, m.Status)

	// Field (8) 'ActivationEpoch'
	dst = ssz.MarshalUint64(dst, m.ActivationEpoch)

	// Field (9) 'OwnerAddress'
	dst = append(dst, m.OwnerAddress[:]...)

	// Field (10) 'Liquidated'
	dst = ssz.MarshalBool(dst, m.Liquidated)

	// Field (2) 'SharePubKey'
	if size := len(m.SharePubKey); size > 48 {
		err = ssz.ErrBytesLengthFn("migration_6_OldStorageShare.SharePubKey", size, 48)
		return
	}
	dst = append(dst, m.SharePubKey...)

	// Field (3) 'Committee'
	if size := len(m.Committee); size > 13 {
		err = ssz.ErrListTooBigFn("migration_6_OldStorageShare.Committee", size, 13)
		return
	}
	{
		offset = 4 * len(m.Committee)
		for ii := 0; ii < len(m.Committee); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += m.Committee[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(m.Committee); ii++ {
		if dst, err = m.Committee[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (6) 'Graffiti'
	if size := len(m.Graffiti); size > 32 {
		err = ssz.ErrBytesLengthFn("migration_6_OldStorageShare.Graffiti", size, 32)
		return
	}
	dst = append(dst, m.Graffiti...)

	return
}

// UnmarshalSSZ ssz unmarshals the migration_6_OldStorageShare object
func (m *migration_6_OldStorageShare) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 129 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o6 uint64

	// Field (0) 'ValidatorIndex'
	m.ValidatorIndex = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorPubKey'
	if cap(m.ValidatorPubKey) == 0 {
		m.ValidatorPubKey = make([]byte, 0, len(buf[8:56]))
	}
	m.ValidatorPubKey = append(m.ValidatorPubKey, buf[8:56]...)

	// Offset (2) 'SharePubKey'
	if o2 = ssz.ReadOffset(buf[56:60]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 != 129 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'Committee'
	if o3 = ssz.ReadOffset(buf[60:64]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'DomainType'
	copy(m.DomainType[:], buf[64:68])

	// Field (5) 'FeeRecipientAddress'
	copy(m.FeeRecipientAddress[:], buf[68:88])

	// Offset (6) 'Graffiti'
	if o6 = ssz.ReadOffset(buf[88:92]); o6 > size || o3 > o6 {
		return ssz.ErrOffset
	}

	// Field (7) 'Status'
	m.Status = ssz.UnmarshallUint64(buf[92:100])

	// Field (8) 'ActivationEpoch'
	m.ActivationEpoch = ssz.UnmarshallUint64(buf[100:108])

	// Field (9) 'OwnerAddress'
	copy(m.OwnerAddress[:], buf[108:128])

	// Field (10) 'Liquidated'
	m.Liquidated = ssz.UnmarshalBool(buf[128:129])

	// Field (2) 'SharePubKey'
	{
		buf = tail[o2:o3]
		if len(buf) > 48 {
			return ssz.ErrBytesLength
		}
		if cap(m.SharePubKey) == 0 {
			m.SharePubKey = make([]byte, 0, len(buf))
		}
		m.SharePubKey = append(m.SharePubKey, buf...)
	}

	// Field (3) 'Committee'
	{
		buf = tail[o3:o6]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		m.Committee = make([]*migration_6_OldStorageOperator, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if m.Committee[indx] == nil {
				m.Committee[indx] = new(migration_6_OldStorageOperator)
			}
			if err = m.Committee[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (6) 'Graffiti'
	{
		buf = tail[o6:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(m.Graffiti) == 0 {
			m.Graffiti = make([]byte, 0, len(buf))
		}
		m.Graffiti = append(m.Graffiti, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the migration_6_OldStorageShare object
func (m *migration_6_OldStorageShare) SizeSSZ() (size int) {
	size = 129

	// Field (2) 'SharePubKey'
	size += len(m.SharePubKey)

	// Field (3) 'Committee'
	for ii := 0; ii < len(m.Committee); ii++ {
		size += 4
		size += m.Committee[ii].SizeSSZ()
	}

	// Field (6) 'Graffiti'
	size += len(m.Graffiti)

	return
}

// HashTreeRoot ssz hashes the migration_6_OldStorageShare object
func (m *migration_6_OldStorageShare) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the migration_6_OldStorageShare object with a hasher
func (m *migration_6_OldStorageShare) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ValidatorIndex'
	hh.PutUint64(m.ValidatorIndex)

	// Field (1) 'ValidatorPubKey'
	if size := len(m.ValidatorPubKey); size != 48 {
		err = ssz.ErrBytesLengthFn("migration_6_OldStorageShare.ValidatorPubKey", size, 48)
		return
	}
	hh.PutBytes(m.ValidatorPubKey)

	// Field (2) 'SharePubKey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.SharePubKey))
		if byteLen > 48 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(m.SharePubKey)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (48+31)/32)
	}

	// Field (3) 'Committee'
	{
		subIndx := hh.Index()
		num := uint64(len(m.Committee))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range m.Committee {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	// Field (4) 'DomainType'
	hh.PutBytes(m.DomainType[:])

	// Field (5) 'FeeRecipientAddress'
	hh.PutBytes(m.FeeRecipientAddress[:])

	// Field (6) 'Graffiti'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(m.Graffiti))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(m.Graffiti)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	// Field (7) 'Status'
	hh.PutUint64(m.Status)

	// Field (8) 'ActivationEpoch'
	hh.PutUint64(m.ActivationEpoch)

	// Field (9) 'OwnerAddress'
	hh.PutBytes(m.OwnerAddress[:])

	// Field (10) 'Liquidated'
	hh.PutBool(m.Liquidated)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the migration_6_OldStorageShare object
func (m *migration_6_OldStorageShare) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}
