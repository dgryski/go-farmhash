// Package farmhash is cgo bindings against Google's FarmHash family of non-cryptographic hash functions
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


*/
import "C"

import "unsafe"

func Hash32(s []byte) uint32 {
	return uint32(C.cHash32((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

func Hash64(s []byte) uint64 {
	return uint64(C.cHash64((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

func Hash32Seed(s []byte, seed uint32) uint32 {
	return uint32(C.cHash32Seed((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint32_t(seed)))
}

func Hash64Seed(s []byte, seed uint64) uint64 {
	return uint64(C.cHash64Seed((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint64_t(seed)))
}

func Hash64Seeds(s []byte, seed0, seed1 uint64) uint64 {
	return uint64(C.cHash64Seeds((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s)), C.uint64_t(seed0), C.uint64_t(seed1)))
}

func Fingerprint32(s []byte) uint32 {
	return uint32(C.cFingerprint32((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}

func Fingerprint64(s []byte) uint64 {
	return uint64(C.cFingerprint64((*C.char)(unsafe.Pointer(&s[0])), C.size_t(len(s))))
}
