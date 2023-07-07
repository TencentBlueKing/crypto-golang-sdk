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
	"testing"
)

var (
	aesGcmKey       = []byte("IHDggaKpZcMWMtVbLzRxVbOmX5c0L2F2")
	aesGcmNonce     = []byte("RDUpW3YNs4Kw")
	aesGcmPlaintext = []byte("37045a07-8d63-494e-b2d1-8ded581b675a")
)

func TestAesGcmEncryptDecrypt(t *testing.T) {
	crypto, err := NewAesGcm(aesGcmKey, aesGcmNonce)
	if err != nil {
		t.Fatalf("new aes gcm failed, err: %v", err)
	}

	ciphertext, err := crypto.Encrypt(aesGcmPlaintext)
	if err != nil {
		t.Fatalf("encrypt failed, err: %v", err)
	}

	decrypted, err := crypto.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("decrypt failed, err: %v", err)
	}

	if string(decrypted) != string(aesGcmPlaintext) {
		t.Fatalf("decrypted text %s is not the same with plaintext %s, key: %s", decrypted, aesGcmPlaintext, aesGcmKey)
	}
}

func BenchmarkAesGcmEncrypt(b *testing.B) {
	crypto, err := NewAesGcm(aesGcmKey, aesGcmNonce)
	if err != nil {
		b.Fatalf("new aes gcm failed, err: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = crypto.Encrypt(aesGcmPlaintext)
	}
}

func BenchmarkAesGcmDecrypt(b *testing.B) {
	crypto, err := NewAesGcm(aesGcmKey, aesGcmNonce)
	if err != nil {
		b.Fatalf("new aes gcm failed, err: %v", err)
	}

	ciphertext, err := crypto.Encrypt(aesGcmPlaintext)
	if err != nil {
		b.Fatalf("encrypt failed, err: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = crypto.Decrypt(ciphertext)
	}
}
