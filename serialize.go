// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cookiejar

import (
	"sort"
	"time"
)

func (j *Jar) Export() []Entry {
	j.deleteExpired(time.Now())
	return j.allPersistentEntries()
}

func (j *Jar) Import(entries []Entry) {
	j.merge(entries)
	j.deleteExpired(time.Now())
}

// allPersistentEntries returns all the entries in the jar, sorted by primarly by canonical host
// name and secondarily by path length.
func (j *Jar) allPersistentEntries() []Entry {
	var entries []Entry
	for _, submap := range j.entries {
		for _, e := range submap {
			if e.Persistent {
				entries = append(entries, e)
			}
		}
	}
	sort.Sort(byCanonicalHost{entries})
	return entries
}
