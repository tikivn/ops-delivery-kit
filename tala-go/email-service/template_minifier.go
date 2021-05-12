// Copyright (c) 2020 TNSL Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package email_service

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

type Compressor struct {
	M *minify.M
}

func NewCompressor() *Compressor {
	m := minify.New()

	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)

	return &Compressor{
		M: m,
	}
}

var (
	compressor = NewCompressor()
)

//ReadAndMinifyTemplate
func (m *Compressor) ReadAndMinifyTemplate(name string) (tpl *template.Template, err error) {

	root := "./email"

	if val := os.Getenv("ROOT_EMAIL"); val != "" {
		root = val
	}

	fullPath := fmt.Sprintf(path.Join(root, "./tpl/%s.tpl.html"), name)
	file, err := os.Open(fullPath)

	if err != nil {
		return tpl, err
	}

	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)

	if err != nil {
		return tpl, err
	}

	out, err := m.M.Bytes("text/html", fileContent)

	if err != nil {
		return tpl, err
	}

	tpl, err = template.New(fullPath).Parse(string(out))

	if err != nil {
		return tpl, err
	}

	return tpl, err
}
