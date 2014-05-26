

#include <farmhash.h>
#include <stdint.h>

extern "C" {

uint32_t cHash32(const char *s, size_t len) { return util::Hash32(s, len); }
uint64_t cHash64(const char *s, size_t len) { return util::Hash64(s, len); }

uint32_t cHash32Seed(const char *s, size_t len, uint32_t seed) { return util::Hash32WithSeed(s, len, seed); }
uint64_t cHash64Seed(const char *s, size_t len, uint64_t seed) { return util::Hash64WithSeed(s, len, seed); }
uint64_t cHash64Seeds(const char *s, size_t len, uint64_t seed0, uint64_t seed1) { return util::Hash64WithSeeds(s, len, seed0, seed1); }

void cHash128Seeds(const char *s, size_t len, uint64_t seed0, uint64_t seed1, uint64_t *lo, uint64_t *hi) {
    util::uint128_t f128 = util::Hash128WithSeed(s, len, util::Uint128(seed0, seed1));
    *lo = util::Uint128Low64(f128);
    *hi = util::Uint128High64(f128);
}

void cHash128(const char *s, size_t len, uint64_t *lo, uint64_t *hi) {
    util::uint128_t f128 = util::Hash128(s, len);
    *lo = util::Uint128Low64(f128);
    *hi = util::Uint128High64(f128);
}

uint32_t cFingerprint32(const char *s, size_t len) { return util::Fingerprint32(s, len); }
uint64_t cFingerprint64(const char *s, size_t len) { return util::Fingerprint64(s, len); }

void cFingerprint128(const char *s, size_t len, uint64_t *hi, uint64_t *lo) {
    util::uint128_t f128 = util::Fingerprint128(s, len);
    *lo = util::Uint128Low64(f128);
    *hi = util::Uint128High64(f128);
}

}
