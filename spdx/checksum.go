// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx

type ChecksumAlgorithm2_2 int

const (
	SHA1 ChecksumAlgorithm2_2 = iota
	SHA224
	SHA256
	SHA384
	SHA512
	MD2
	MD4
	MD5
	MD6
)

type Checksum2_2 struct {
	Algorithm ChecksumAlgorithm2_2
	Value     string
}
