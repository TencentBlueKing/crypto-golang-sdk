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

/*
#cgo windows CFLAGS: -DWIN32_LEAN_AND_MEAN
#cgo windows pkg-config: libssl libcrypto
#include <stdio.h>
#include <string.h>
#include <openssl/evp.h>
#include <openssl/err.h>
#include <openssl/ec.h>
#include <openssl/pem.h>
#include <openssl/ossl_typ.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
#include <openssl/evp.h>
#include <openssl/err.h>
#include <openssl/ec.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

type sm4 struct {
	key []byte
	iv  []byte
}

const (
	sm4KeyLen = 16
	sm4IvLen  = 16
)

// NewSm4 returns a new SM4 crypto instance
func NewSm4(key, iv []byte) (Crypto, error) {
	keyLen := len(key)
	if keyLen != sm4KeyLen {
		return nil, fmt.Errorf("key length %d is invalid, should be %d", keyLen, sm4KeyLen)
	}

	ivLen := len(iv)
	if ivLen != sm4IvLen {
		return nil, fmt.Errorf("iv length %d is invalid, should be %d", ivLen, sm4IvLen)
	}
	return &sm4{key: key, iv: iv}, nil
}

// Encrypt plaintext.
func (s *sm4) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return make([]byte, 0), nil
	}

	// new sm4 encrypt context
	encryptCtx := C.EVP_CIPHER_CTX_new()
	if encryptCtx == nil {
		return nil, errors.New("new evp cipher context for encryption failed")
	}
	defer func(encryptCtx *C.EVP_CIPHER_CTX) {
		C.EVP_CIPHER_CTX_free(encryptCtx)
	}(encryptCtx)

	if C.EVP_EncryptInit_ex(encryptCtx, C.EVP_sm4_ctr(), nil, nil, nil) != 1 {
		return nil, errors.New("init sm4 encrypt context failed")
	}

	if C.EVP_EncryptInit_ex(encryptCtx, nil, nil, (*C.uchar)(&s.key[0]), (*C.uchar)(&s.iv[0])) != 1 {
		return nil, fmt.Errorf("init sm4 encrypt context by key(%s) and iv(%s) failed", s.key, s.iv)
	}

	// encrypt plain text
	ciphertextLen := len(plaintext) + sm4KeyLen
	ciphertext := make([]byte, ciphertextLen)
	cipherCLen := C.int(ciphertextLen)
	res := C.EVP_EncryptUpdate(encryptCtx, (*C.uchar)(&ciphertext[0]), &cipherCLen, (*C.uchar)(&plaintext[0]),
		C.int(len(plaintext)))
	if res != 1 {
		return nil, fmt.Errorf("encrypt %s failed, res: %v", plaintext, res)
	}

	// do final encryption
	final := make([]byte, sm4KeyLen)
	var finalLen C.int
	if C.EVP_EncryptFinal_ex(encryptCtx, (*C.uchar)(&final[0]), &finalLen) != 1 {
		return nil, fmt.Errorf("encrypt final failed, plaintext: %s, res: %v", plaintext, res)
	}

	return append(ciphertext[:cipherCLen], final[:finalLen]...), nil
}

// Decrypt ciphertext.
func (s *sm4) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return make([]byte, 0), nil
	}

	// new sm4 decrypt context
	decryptCtx := C.EVP_CIPHER_CTX_new()
	if decryptCtx == nil {
		return nil, errors.New("new evp cipher context for decryption failed")
	}

	if C.EVP_DecryptInit_ex(decryptCtx, C.EVP_sm4_ctr(), nil, nil, nil) != 1 {
		return nil, errors.New("init sm4 decrypt context failed")
	}

	if C.EVP_DecryptInit_ex(decryptCtx, nil, nil, (*C.uchar)(&s.key[0]), (*C.uchar)(&s.iv[0])) != 1 {
		return nil, fmt.Errorf("init sm4 decrypt context by key(%s) and iv(%s) failed", s.key, s.iv)
	}
	defer func(decryptCtx *C.EVP_CIPHER_CTX) {
		C.EVP_CIPHER_CTX_free(decryptCtx)
	}(decryptCtx)

	// decrypt cipher text
	plaintext := make([]byte, len(ciphertext)+sm4KeyLen)
	plainCLen := C.int(len(plaintext))
	res := C.EVP_DecryptUpdate(decryptCtx, (*C.uchar)(&plaintext[0]), &plainCLen, (*C.uchar)(&ciphertext[0]),
		C.int(len(ciphertext)))
	if res != 1 {
		return nil, fmt.Errorf("decrypt %s failed, res: %v", ciphertext, res)
	}

	// do final decryption
	final := make([]byte, sm4KeyLen)
	var finalLen C.int
	if C.EVP_DecryptFinal_ex(decryptCtx, (*C.uchar)(&final[0]), &finalLen) != 1 {
		return nil, fmt.Errorf("decrypt final failed, ciphertext: %s, res: %v", ciphertext, res)
	}

	return append(plaintext[:plainCLen], final[:finalLen]...), nil
}
