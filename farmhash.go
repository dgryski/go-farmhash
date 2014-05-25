// Package farmhash is cgo bindings against Google's FarmHash family of non-cryptographic hash functions
/*

https://code.google.com/p/farmhash/

*/
package farmhash

/*
#cgo LDFLAGS: -lfarmhash

#include <stdint.h>
#include <stddef.h>

uint32_t cHash32(const char *s, size_t len);
uint64_t cHash64(const char *s, size_t len);

uint32_t cHash32Seed(const char *s, size_t len, uint32_t seed);
uint64_t cHash64Seed(const char *s, size_t len, uint64_t seed);
uint64_t cHash64Seeds(const char *s, size_t len, uint64_t seed0, uint64_t seed1);

uint32_t cFingerprint32(const char *s, size_t len);
uint64_t cFingerprint64(const char *s, size_t len);

void cFingerprint128(const char *s, size_t len, uint64_t *hi, uint64_t *lo);

*/
import "C"

import "unsafe"

// Hash32 hashes a byte slice and returns a uint32 hash value
func Hash32(s []byte) uint32 {
	return uint32(C.cHash32((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

// Hash64 hashes a byte slice and returns a uint64 hash value
func Hash64(s []byte) uint64 {
	return uint64(C.cHash64((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

// Hash32Seed hashes a byte slice and a uint32 seed and returns a uint32 hash value
func Hash32Seed(s []byte, seed uint32) uint32 {
	return uint32(C.cHash32Seed((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint32_t(seed)))
}

// Hash64Seed hashes a byte slice and a uint64 seed and returns a uint64 hash value
func Hash64Seed(s []byte, seed uint64) uint64 {
	return uint64(C.cHash64Seed((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint64_t(seed)))
}

// Hash64Seeds hashes a byte slice and two uint64 seeds and returns a uint64 hash value
func Hash64Seeds(s []byte, seed0, seed1 uint64) uint64 {
	return uint64(C.cHash64Seeds((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint64_t(seed0), C.uint64_t(seed1)))
}

// Fingerprint32 is a 32-bit fingerprint function for byte-slices
func Fingerprint32(s []byte) uint32 {
	return uint32(C.cFingerprint32((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

// Fingerprint64 is a 64-bit fingerprint function for byte-slices
func Fingerprint64(s []byte) uint64 {
	return uint64(C.cFingerprint64((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

// Fingerprint128 is a 128-bit fingerprint function for byte-slices
func Fingerprint128(s []byte) (uint64, uint64) {
	var hi, lo uint64
	C.cFingerprint128((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), (*C.uint64_t)(unsafe.Pointer(&hi)), (*C.uint64_t)(unsafe.Pointer(&lo)))
	return hi, lo
}
