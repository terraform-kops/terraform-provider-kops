package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"sync"
	"text/template"
)

const schemasTemplate = "hack/gen-tf-code/templates/schema.tpl"
const testsTemplate = "hack/gen-tf-code/templates/test.tpl"
const docsTemplate = "hack/gen-tf-code/templates/doc.tpl"

var (
	templateCache     = make(map[string]string)
	templateCacheMux  sync.RWMutex
	dirCreationCache  = make(map[string]bool)
	dirCreationMux    sync.Mutex
)

// getCachedTemplateContent reads template file content with caching
func getCachedTemplateContent(tplFile string) string {
	templateCacheMux.RLock()
	content, exists := templateCache[tplFile]
	templateCacheMux.RUnlock()

	if exists {
		return content
	}

	// Cache miss - read and cache
	templateCacheMux.Lock()
	defer templateCacheMux.Unlock()

	// Double-check after acquiring write lock
	if content, exists := templateCache[tplFile]; exists {
		return content
	}

	content = readFile(tplFile)
	templateCache[tplFile] = content
	return content
}

// ensureDirExists creates directory with caching to avoid repeated MkdirAll calls
func ensureDirExists(folder string) error {
	dirCreationMux.Lock()
	defer dirCreationMux.Unlock()

	if dirCreationCache[folder] {
		return nil
	}

	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}

	dirCreationCache[folder] = true
	return nil
}

func executeTemplate(t reflect.Type, tplFile, folder, file string, funcMaps ...template.FuncMap) {
	tmpl := template.New(tplFile)
	for _, funcMap := range funcMaps {
		tmpl = tmpl.Funcs(funcMap)
	}

	content := getCachedTemplateContent(tplFile)
	if tmpl, err := tmpl.Parse(content); err != nil {
		panic(err)
	} else {
		if err := ensureDirExists(folder); err != nil {
			panic(fmt.Sprintf("Failed to create directories for %s", folder))
		}
		f, err := os.Create(path.Join(folder, file))
		if err != nil {
			panic(fmt.Sprintf("Failed to create file %s", path.Join(folder, file)))
		}
		defer f.Close()
		if err := tmpl.Execute(f, t); err != nil {
			panic(err)
		}
	}
}
