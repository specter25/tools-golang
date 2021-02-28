// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v2

import (
	"fmt"
	"io"
	"sort"

	"github.com/spdx/tools-golang/spdx"
)

func renderFile2_2(f *spdx.File2_2, w io.Writer) error {
	if f.FileName != "" {
		fmt.Fprintf(w, "FileName: %s\n", f.FileName)
	}
	if f.FileSPDXIdentifier != "" {
		fmt.Fprintf(w, "SPDXID: %s\n", spdx.RenderElementID(f.FileSPDXIdentifier))
	}
	for _, s := range f.FileType {
		fmt.Fprintf(w, "FileType: %s\n", s)
	}
	for _, checksum := range f.FileChecksums {
		switch checksum.Algorithm {
		case 1:
			fmt.Fprintf(w, "FileChecksum: SHA1: %s\n", checksum.Value)
		case 2:
			fmt.Fprintf(w, "FileChecksum: SHA224: %s\n", checksum.Value)
		case 3:
			fmt.Fprintf(w, "FileChecksum: SHA256: %s\n", checksum.Value)
		case 4:
			fmt.Fprintf(w, "FileChecksum: SHA384: %s\n", checksum.Value)
		case 5:
			fmt.Fprintf(w, "FileChecksum: SHA512: %s\n", checksum.Value)
		case 6:
			fmt.Fprintf(w, "FileChecksum: MD2: %s\n", checksum.Value)
		case 7:
			fmt.Fprintf(w, "FileChecksum: MD4: %s\n", checksum.Value)
		case 8:
			fmt.Fprintf(w, "FileChecksum: MD5: %s\n", checksum.Value)
		case 9:
			fmt.Fprintf(w, "FileChecksum: MD6: %s\n", checksum.Value)
		}
	}

	if f.LicenseConcluded != "" {
		fmt.Fprintf(w, "LicenseConcluded: %s\n", f.LicenseConcluded)
	}
	for _, s := range f.LicenseInfoInFile {
		fmt.Fprintf(w, "LicenseInfoInFile: %s\n", s)
	}
	if f.LicenseComments != "" {
		fmt.Fprintf(w, "LicenseComments: %s\n", f.LicenseComments)
	}
	if f.FileCopyrightText != "" {
		fmt.Fprintf(w, "FileCopyrightText: %s\n", textify(f.FileCopyrightText))
	}
	for _, aop := range f.ArtifactOfProjects {
		fmt.Fprintf(w, "ArtifactOfProjectName: %s\n", aop.Name)
		if aop.HomePage != "" {
			fmt.Fprintf(w, "ArtifactOfProjectHomePage: %s\n", aop.HomePage)
		}
		if aop.URI != "" {
			fmt.Fprintf(w, "ArtifactOfProjectURI: %s\n", aop.URI)
		}
	}
	if f.FileComment != "" {
		fmt.Fprintf(w, "FileComment: %s\n", f.FileComment)
	}
	if f.FileNotice != "" {
		fmt.Fprintf(w, "FileNotice: %s\n", f.FileNotice)
	}
	for _, s := range f.FileContributor {
		fmt.Fprintf(w, "FileContributor: %s\n", s)
	}
	for _, s := range f.FileAttributionTexts {
		fmt.Fprintf(w, "FileAttributionText: %s\n", textify(s))
	}
	for _, s := range f.FileDependencies {
		fmt.Fprintf(w, "FileDependency: %s\n", s)
	}

	fmt.Fprintf(w, "\n")

	// also render any snippets for this file
	// get slice of Snippet identifiers so we can sort them
	snippetKeys := []string{}
	for k := range f.Snippets {
		snippetKeys = append(snippetKeys, string(k))
	}
	sort.Strings(snippetKeys)
	for _, sID := range snippetKeys {
		s := f.Snippets[spdx.ElementID(sID)]
		renderSnippet2_2(s, w)
	}

	return nil
}
