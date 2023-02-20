package blockchain

import "crypto/sha256"

var characterSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// GenerateRandomNumber : 간단한(가벼운) 의사 난수 생성기(PRNG)
func GenerateRandomNumber(seed uint64) uint64 {
	seed ^= seed << 21
	seed ^= seed >> 35
	seed ^= seed << 4
	return seed
}

// RandomString : 인자로 받은 문자열(str)을 랜덤하게 변경하고 또 다른 난수 반환
// => 인자로 받은 문자열 메모리를 재사용하여 메모리 관련 처리 속도를 최적화
// offset: 기존 블록체인에 해당하는 문자열의 길이
func RandomString(str []byte, offset int, seed uint64) uint64 {
	for i := offset; i < len(str); i++ {
		seed = GenerateRandomNumber(seed)
		str[i] = characterSet[seed%62]
	}
	return seed
}

// Consensus : 블록체인 자격 증명 여부 확인
// 자격증명을 하려는 데이터(data)의 해시값이 자격 조건에 맞는지 검증
// Simple Block Chain 에서는 앞의 특정 비트들(bits 로 결정)이 0인지 여부 확인
func Consensus(data []byte, bits int) bool {
	hashed := sha256.Sum256(data)

	// 자격 조건에서 모두 0이 되어야 하는 바이트 수
	nBytes := bits / 8

	// 자격 조건에서 모두 0이 되어야 하는 바이트 수를 제외한 비트 수
	nBits := bits % 8
	idx := 0
	for ; idx < nBytes; idx++ {
		if hashed[idx] > 0 {
			return false
		}
	}

	// 나머지 비트 수에 대해서도 0이 되어야 함
	return (hashed[idx] >> (8 - nBits)) == 0
}
