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

package engine

import (
	"github.com/r7wx/luna-dns/internal/config"
	"github.com/r7wx/luna-dns/internal/entry"
	"github.com/r7wx/luna-dns/internal/logger"
	"github.com/r7wx/luna-dns/internal/tree"
)

// Engine - DNS Engine
type Engine struct {
	overrides *tree.Tree
	addr      string
	dns       []string
}

// NewEngine - Create a new engine
func NewEngine(config *config.Config) (*Engine, error) {
	logger.Info("Initializing engine...")

	overrides := []entry.Entry{}
	for _, override := range config.Overrides {
		entry, err := entry.NewEntry(override.Domain, override.IP)
		if err != nil {
			return nil, err
		}
		overrides = append(overrides, *entry)
	}

	logger.Info("Engine ready")
	return &Engine{
		overrides: tree.NewTree(overrides),
		addr:      config.Addr,
		dns:       config.DNS,
	}, nil
}