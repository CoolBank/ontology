/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package payload

import (
	"io"

	"github.com/Ontology/common/serialization"
	"github.com/ontio/ontology-crypto/keypair"
	"fmt"
)

const BookKeeperPayloadVersion byte = 0x00

type BookKeeperAction byte

const (
	BookKeeperAction_ADD BookKeeperAction = 0
	BookKeeperAction_SUB BookKeeperAction = 1
)

type Bookkeeper struct {
	PubKey keypair.PublicKey
	Action BookKeeperAction
	Cert   []byte
	Issuer keypair.PublicKey
}

func (self *Bookkeeper) Serialize(w io.Writer) error {
	err := serialization.WriteVarBytes(w, keypair.SerializePublicKey(self.PubKey))
	if err != nil {
		return fmt.Errorf("[Bookkeeper], serializing PubKey failed: %s", err)
	}
	err = serialization.WriteVarBytes(w, []byte{byte(self.Action)})
	if err != nil {
		return fmt.Errorf("[Bookkeeper], serializing Action failed: %s", err)
	}
	err = serialization.WriteVarBytes(w, self.Cert)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], serializing Cert failed: %s", err)
	}
	err = serialization.WriteVarBytes(w, keypair.SerializePublicKey(self.Issuer))
	if err != nil {
		return fmt.Errorf("[Bookkeeper], serializing Issuer failed: %s", err)
	}
	return nil
}

func (self *Bookkeeper) Deserialize(r io.Reader) error {
	buf, err := serialization.ReadVarBytes(r)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], deserializing PubKey failed: %s", err)
	}
	self.PubKey, err = keypair.DeserializePublicKey(buf)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], deserializing PubKey failed: %s", err)
	}

	var p [1]byte
	n, err := r.Read(p[:])
	if n == 0 {
		return fmt.Errorf("[Bookkeeper], deserializing Action failed: %s", err)
	}
	self.Action = BookKeeperAction(p[0])
	self.Cert, err = serialization.ReadVarBytes(r)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], deserializing Cert failed: %s", err)
	}

	buf, err = serialization.ReadVarBytes(r)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], deserializing Issuer failed: %s", err)
	}
	self.Issuer, err = keypair.DeserializePublicKey(buf)
	if err != nil {
		return fmt.Errorf("[Bookkeeper], deserializing Issuer failed: %s", err)
	}

	return nil
}
