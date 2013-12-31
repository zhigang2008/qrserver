// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

// get gets the string of length n for id from the given 4-byte string index.
func get(idx string, id, n int) string {
	return idx[id<<2:][:n]
}

// cmp returns an integer comparing a and b lexicographically.
func cmp(a string, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i, c := range b[:n] {
		switch {
		case a[i] > c:
			return 1
		case a[i] < c:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}

// search searchs for the insertion point of key in smap, which is a
// string with consecutive 4-byte entries. Only the first len(key)
// bytes from the start of the 4-byte entries will be considered.
func search(smap string, key []byte) int {
	n := len(key)
	return sort.Search(len(smap)>>2, func(i int) bool {
		return cmp(get(smap, i, n), key) != -1
	}) << 2
}

func index(smap string, key []byte) int {
	i := search(smap, key)
	if cmp(smap[i:i+len(key)], key) != 0 {
		return -1
	}
	return i
}

func searchUint(imap []uint16, key uint16) int {
	return sort.Search(len(imap), func(i int) bool {
		return imap[i] >= key
	})
}

// fixCase reformats s to the same pattern of cases as pat.
// If returns false if string s is malformed.
func fixCase(pat string, b []byte) bool {
	if len(pat) != len(b) {
		return false
	}
	for i, c := range b {
		r := pat[i]
		if r <= 'Z' {
			if c >= 'a' {
				c -= 'z' - 'Z'
			}
			if c > 'Z' || c < 'A' {
				return false
			}
		} else {
			if c <= 'Z' {
				c += 'z' - 'Z'
			}
			if c > 'z' || c < 'a' {
				return false
			}
		}
		b[i] = c
	}
	return true
}

type langID uint16

// getLangID returns the langID of s if s is a canonical subtag
// or langUnknown if s is not a canonical subtag.
func getLangID(s []byte) (langID, error) {
	if len(s) == 2 {
		return getLangISO2(s)
	}
	return getLangISO3(s)
}

// mapLang returns the mapped langID of id according to mapping m.
func normLang(m []fromTo, id langID) langID {
	k := sort.Search(len(m), func(i int) bool {
		return m[i].from >= uint16(id)
	})
	if k < len(m) && m[k].from == uint16(id) {
		return langID(m[k].to)
	}
	return id
}

// getLangISO2 returns the langID for the given 2-letter ISO language code
// or unknownLang if this does not exist.
func getLangISO2(s []byte) (langID, error) {
	if len(s) == 2 && fixCase("zz", s) {
		if i := index(lang, s); i != -1 && lang[i+3] != 0 {
			return langID(i >> 2), nil
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

const base = 'z' - 'a' + 1

func strToInt(s []byte) uint {
	v := uint(0)
	for i := 0; i < len(s); i++ {
		v *= base
		v += uint(s[i] - 'a')
	}
	return v
}

// converts the given integer to the original ASCII string passed to strToInt.
// len(s) must match the number of characters obtained.
func intToStr(v uint, s []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		s[i] = byte(v%base) + 'a'
		v /= base
	}
}

// getLangISO3 returns the langID for the given 3-letter ISO language code
// or unknownLang if this does not exist.
func getLangISO3(s []byte) (langID, error) {
	if fixCase("und", s) {
		// first try to match canonical 3-letter entries
		for i := search(lang, s[:2]); cmp(lang[i:i+2], s[:2]) == 0; i += 4 {
			if lang[i+3] == 0 && lang[i+2] == s[2] {
				// We treat "und" as special and always translate it to "unspecified".
				// Note that ZZ and Zzzz are private use and are not treated as
				// unspecified by default.
				id := langID(i >> 2)
				if id == nonCanonicalUnd {
					return 0, nil
				}
				return id, nil
			}
		}
		if i := index(altLangISO3, s); i != -1 {
			return langID(altLangIndex[altLangISO3[i+3]]), nil
		}
		n := strToInt(s)
		if langNoIndex[n/8]&(1<<(n%8)) != 0 {
			return langID(n) + langNoIndexOffset, nil
		}
		// Check for non-canonical uses of ISO3.
		for i := search(lang, s[:1]); lang[i] == s[0]; i += 4 {
			if cmp(lang[i+2:][:2], s[1:3]) == 0 {
				return langID(i >> 2), nil
			}
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

// stringToBuf writes the string to b and returns the number of bytes
// written.  cap(b) must be >= 3.
func (id langID) stringToBuf(b []byte) int {
	if id >= langNoIndexOffset {
		intToStr(uint(id)-langNoIndexOffset, b[:3])
		return 3
	} else if id == 0 {
		return copy(b, "und")
	}
	l := lang[id<<2:]
	if l[3] == 0 {
		return copy(b, l[:3])
	}
	return copy(b, l[:2])
}

// String returns the BCP 47 representation of the langID.
// Use b as variable name, instead of id, to ensure the variable
// used is consistent with that of Base in which this type is embedded.
func (b langID) String() string {
	if b == 0 {
		return "und"
	} else if b >= langNoIndexOffset {
		b -= langNoIndexOffset
		buf := [3]byte{}
		intToStr(uint(b), buf[:])
		return string(buf[:])
	}
	l := lang[b<<2:]
	if l[3] == 0 {
		return l[:3]
	}
	return l[:2]
}

// ISO3 returns the ISO 639-3 language code.
func (b langID) ISO3() string {
	if b == 0 || b >= langNoIndexOffset {
		return b.String()
	}
	l := lang[b<<2:]
	if l[3] == 0 {
		return l[:3]
	} else if l[2] == 0 {
		return get(altLangISO3, int(l[3]), 3)
	}
	// This allocation will only happen for 3-letter ISO codes
	// that are non-canonical BCP 47 language identifiers.
	return l[0:1] + l[2:4]
}

// IsPrivateUse reports whether this language code is reserved for private use.
func (b langID) IsPrivateUse() bool {
	return langPrivateStart <= b && b <= langPrivateEnd
}

type regionID uint16

// getRegionID returns the region id for s if s is a valid 2-letter region code
// or unknownRegion.
func getRegionID(s []byte) (regionID, error) {
	if len(s) == 3 {
		if isAlpha(s[0]) {
			return getRegionISO3(s)
		}
		if i, err := strconv.ParseUint(string(s), 10, 10); err == nil {
			return getRegionM49(int(i))
		}
	}
	return getRegionISO2(s)
}

// getRegionISO2 returns the regionID for the given 2-letter ISO country code
// or unknownRegion if this does not exist.
func getRegionISO2(s []byte) (regionID, error) {
	if fixCase("ZZ", s) {
		if i := index(regionISO, s); i != -1 {
			return regionID(i>>2) + isoRegionOffset, nil
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

// getRegionISO3 returns the regionID for the given 3-letter ISO country code
// or unknownRegion if this does not exist.
func getRegionISO3(s []byte) (regionID, error) {
	if fixCase("ZZZ", s) {
		for i := search(regionISO, s[:1]); regionISO[i] == s[0]; i += 4 {
			if cmp(regionISO[i+2:][:2], s[1:3]) == 0 {
				return regionID(i>>2) + isoRegionOffset, nil
			}
		}
		for i := 0; i < len(altRegionISO3); i += 3 {
			if cmp(altRegionISO3[i:i+3], s) == 0 {
				return regionID(altRegionIDs[i/3]), nil
			}
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

func getRegionM49(n int) (regionID, error) {
	if 0 < n && n <= 999 {
		const (
			searchBits = 7
			regionBits = 9
			regionMask = 1<<regionBits - 1
		)
		idx := n >> searchBits
		buf := fromM49[m49Index[idx]:m49Index[idx+1]]
		val := uint16(n) << regionBits // we rely on bits shifting out
		i := sort.Search(len(buf), func(i int) bool {
			return buf[i] >= val
		})
		if r := buf[i]; r&^regionMask == val {
			return regionID(r & regionMask), nil
		}
	}
	var e ValueError
	fmt.Fprint(bytes.NewBuffer([]byte(e.v[:])), n)
	return 0, e
}

// normRegion returns a region if r is deprecated or 0 otherwise.
// TODO: consider supporting BYS (-> BLR), CSK (-> 200 or CZ), PHI (-> PHL) and AFI (-> DJ).
// TODO: consider mapping split up regions to new most populous one (like CLDR).
func normRegion(r regionID) regionID {
	m := regionOldMap
	k := sort.Search(len(m), func(i int) bool {
		return m[i].from >= uint16(r)
	})
	if k < len(m) && m[k].from == uint16(r) {
		return regionID(m[k].to)
	}
	return 0
}

// String returns the BCP 47 representation for the region.
// It returns "ZZ" for an unspecified region.
func (r regionID) String() string {
	if r < isoRegionOffset {
		if r == 0 {
			return "ZZ"
		}
		return fmt.Sprintf("%03d", r.M49())
	}
	r -= isoRegionOffset
	return get(regionISO, int(r), 2)
}

// ISO3 returns the 3-letter ISO code of r.
// Note that not all regions have a 3-letter ISO code.
// In such cases this method returns the empty string.
func (r regionID) ISO3() string {
	if r < isoRegionOffset {
		return ""
	}
	r -= isoRegionOffset
	reg := regionISO[r<<2:]
	switch reg[2] {
	case 0:
		return altRegionISO3[reg[3]:][:3]
	case ' ':
		return ""
	}
	return reg[0:1] + reg[2:4]
}

// M49 returns the UN M.49 encoding of r, or 0 if this encoding
// is not defined for r.
func (r regionID) M49() int {
	return int(m49[r])
}

// IsPrivateUse reports whether r is reserved for private use.
func (r regionID) IsPrivateUse() bool {
	const m49PrivateUseStart = 900
	return r.M49() >= m49PrivateUseStart
}

type scriptID uint8

// getScriptID returns the script id for string s. It assumes that s
// is of the format [A-Z][a-z]{3}.
func getScriptID(idx string, s []byte) (scriptID, error) {
	if fixCase("Zzzz", s) {
		if i := index(idx, s); i != -1 {
			return scriptID(i >> 2), nil
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

// String returns the script code in title case.
// It returns "Zzzz" for an unspecified script.
func (s scriptID) String() string {
	if s == 0 {
		return "Zzzz"
	}
	return get(script, int(s), 4)
}

// IsPrivateUse reports whether this script code is reserved for private use.
func (s scriptID) IsPrivateUse() bool {
	return scrQaaa <= s && s <= scrQabx
}

type currencyID uint16

func getCurrencyID(idx string, s []byte) (currencyID, error) {
	if fixCase("XXX", s) {
		if i := index(idx, s); i != -1 {
			return currencyID(i >> 2), nil
		}
		return 0, mkErrInvalid(s)
	}
	return 0, errSyntax
}

// String returns the upper case representation of the currency.
func (c currencyID) String() string {
	if c == 0 {
		return "XXX"
	}
	return get(currency, int(c), 3)
}

func round(index string, c currencyID) int {
	return int(index[c<<2+3] >> 2)
}

func decimals(index string, c currencyID) int {
	return int(index[c<<2+3] & 0x03)
}
