/*
MIT License
Copyright (c) 2022 r7wx
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package dtree

import (
	"testing"

	"github.com/r7wx/luna-dns/internal/entry"
)

func TestBasics(t *testing.T) {
	testEntries := func(tree *DTree, entries []map[string]string, t *testing.T) {
		for _, e := range entries {
			testEntry, _ := entry.NewEntry(e["domain"], e["ip"])
			tree.Insert(testEntry)

			targetIP := testEntry.IP
			testEntry.IP = ""

			found := tree.Search(testEntry)
			if found == "" {
				t.Fatal()
				continue
			}

			if found != targetIP {
				t.Fatal()
			}
		}
	}

	tree := NewDTree()
	testEntries(tree, []map[string]string{
		{
			"domain": "test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "a.test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "test.a",
			"ip":     "127.0.0.1",
		},
	}, t)
}

func TestOthers(t *testing.T) {
	insertEntries := func(tree *DTree, entries []map[string]string) {
		for _, e := range entries {
			testEntry, _ := entry.NewEntry(e["domain"], e["ip"])
			tree.Insert(testEntry)
		}
	}

	searchDomains := func(tree *DTree, entries []map[string]any,
		t *testing.T) {
		for _, e := range entries {
			domain := e["domain"].(string)
			expected := e["expected"].(bool)

			testEntry, _ := entry.NewEntry(domain, "")
			found := tree.Search(testEntry)

			if found == "" && expected {
				t.Fatal()
			}
			if found != "" && !expected {
				t.Fatal()
			}
		}
	}

	tree := NewDTree()
	insertEntries(tree, []map[string]string{
		{
			"domain": "*.test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "*.tld",
			"ip":     "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"domain":   "unk.com",
			"expected": false,
		},
		{
			"domain":   "aaa.test.com",
			"expected": true,
		},
		{
			"domain":   "test.tld",
			"expected": true,
		},
		{
			"domain":   "test.xxx",
			"expected": false,
		},
	}, t)

	insertEntries(tree, []map[string]string{
		{
			"domain": "*",
			"ip":     "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"domain":   "google.com",
			"expected": true,
		},
	}, t)
}