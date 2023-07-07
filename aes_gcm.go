/*
 * Tencent is pleased to support the open source community by making
 * 蓝鲸智云 - crypto-golang-sdk (BlueKing - crypto-golang-sdk) available.
 * Copyright (C) 2023 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package bkcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type aesGcm struct {
	nonce []byte
	aead  cipher.AEAD
}

// NewAesGcm returns a new AES-GCM crypto instance
func NewAesGcm(key []byte, nonce []byte) (Crypto, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(nonce) != aead.NonceSize() {
		return nil, fmt.Errorf("nonce length %d is invalid, should be %d", len(nonce), aead.NonceSize())
	}

	return &aesGcm{
		nonce: nonce,
		aead:  aead,
	}, nil
}

// Encrypt plaintext.
func (a *aesGcm) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext := a.aead.Seal(plaintext[:0], a.nonce, plaintext, nil)
	return ciphertext, nil
}

// Decrypt ciphertext.
func (a *aesGcm) Decrypt(ciphertext []byte) ([]byte, error) {
	return a.aead.Open(nil, a.nonce, ciphertext, nil)
}
