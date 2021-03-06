/*
 * Copyright 2014-Present Pivotal Software, Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filecopy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal/go-ape/pkg/filecopy"
	"github.com/pivotal/go-ape/pkg/test_support"
)

var _ = Describe("Checker", func() {
	var (
		checker filecopy.Checker
		tempDir string
		path    string
		exists  bool
	)

	BeforeEach(func() {
		checker = filecopy.NewChecker()
		tempDir = test_support.CreateTempDir()
	})

	JustBeforeEach(func() {
		exists = checker.Exists(path)
	})

	AfterEach(func() {
		test_support.CleanupDirs(GinkgoT(), tempDir)
	})

	Context("when the input file is a directory", func() {
		BeforeEach(func() {
			path = tempDir
		})

		It("should report that the directory exists", func() {
			Expect(exists).To(BeTrue())
		})
	})

	Context("when the input file is a file", func() {
		BeforeEach(func() {
			path = test_support.CreateFile(tempDir, "src.file")
		})

		It("should report that the file exists", func() {
			Expect(exists).To(BeTrue())
		})
	})

	Context("when the input file does not exist", func() {
		BeforeEach(func() {
			path = "nosuch"
		})

		It("should report that the file does not exist", func() {
			Expect(exists).To(BeFalse())
		})
	})
})
