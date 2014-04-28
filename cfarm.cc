

#include <farmhash.h>
#include <stdint.h>

extern "C" {

uint32_t cHash32(const char *s, size_t len) { return util::Hash32(s, len); }
uint64_t cHash64(const char *s, size_t len) { return util::Hash64(s, len); }

uint32_t cHash32Seed(const char *s, size_t len, uint32_t seed) { return util::Hash32WithSeed(s, len, seed); }
uint64_t cHash64Seed(const char *s, size_t len, uint64_t seed) { return util::Hash64WithSeed(s, len, seed); }
uint64_t cHash64Seeds(const char *s, size_t len, uint64_t seed0, uint64_t seed1) { return util::Hash64WithSeeds(s, len, seed0, seed1); }

uint32_t cFingerprint32(const char *s, size_t len) { return util::Fingerprint32(s, len); }
uint64_t cFingerprint64(const char *s, size_t len) { return util::Fingerprint64(s, len); }

}
