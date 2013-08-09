// Copyright (c) 2012 - Cloud Instruments Co., Ltd.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package seelog

import (
	"fmt"
	"io"
	"testing"
)

// fileWriterTestCase is declared in writers_filewriter_test.go

func createRollingSizeFileWriterTestCase(
	files []string,
	fileName string,
	fileSize int64,
	maxRolls int,
	writeCount int,
	resFiles []string) *fileWriterTestCase {

	return &fileWriterTestCase{files, fileName, rollingTypeSize, fileSize, maxRolls, "", writeCount, resFiles}
}

func createRollingDatefileWriterTestCase(
	files []string,
	fileName string,
	datePattern string,
	writeCount int,
	resFiles []string) *fileWriterTestCase {

	return &fileWriterTestCase{files, fileName, rollingTypeDate, 0, 0, datePattern, writeCount, resFiles}
}

func TestRollingFileWriter(t *testing.T) {
	t.Logf("Starting rolling file writer tests")
	newFileWriterTester(rollingfileWriterTests, rollingFileWriterGetter, t).test()
}

//===============================================================

func rollingFileWriterGetter(testCase *fileWriterTestCase) (io.WriteCloser, error) {
	if testCase.rollingType == rollingTypeSize {
		return newRollingFileWriterSize(testCase.fileName, rollingArchiveNone, "", testCase.fileSize, testCase.maxRolls)
	} else if testCase.rollingType == rollingTypeDate {
		return newRollingFileWriterDate(testCase.fileName, rollingArchiveNone, "", testCase.datePattern)
	}

	return nil, fmt.Errorf("Incorrect rollingType")
}

//===============================================================
// TODO: Enable when other errors are fixed.
var rollingfileWriterTests []*fileWriterTestCase = []*fileWriterTestCase{
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 10, 1, []string{"log.testlog"}),
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 10, 2, []string{"log.testlog", "log.testlog.1"}),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.1"}, "log.testlog", 10, 10, 2, []string{"log.testlog", "log.testlog.1", "log.testlog.2"}),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.1"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.2"}),
	createRollingSizeFileWriterTestCase([]string{}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.1"}),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.9"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.10"}),
	createRollingSizeFileWriterTestCase([]string{"log.testlog.a", "log.testlog.1b"}, "log.testlog", 10, 1, 2, []string{"log.testlog", "log.testlog.1", "log.testlog.a", "log.testlog.1b"}),
	createRollingSizeFileWriterTestCase([]string{}, `dir/log.testlog`, 10, 10, 1, []string{`dir/log.testlog`}),
	createRollingSizeFileWriterTestCase([]string{}, `dir/log.testlog`, 10, 10, 2, []string{`dir/log.testlog`, `dir/log.testlog.1`}),
	createRollingSizeFileWriterTestCase([]string{`dir/dir/log.testlog.1`}, `dir/dir/log.testlog`, 10, 10, 2, []string{`dir/dir/log.testlog`, `dir/dir/log.testlog.1`, `dir/dir/log.testlog.2`}),
	createRollingSizeFileWriterTestCase([]string{`dir/dir/dir/log.testlog.1`}, `dir/dir/dir/log.testlog`, 10, 1, 2, []string{`dir/dir/dir/log.testlog`, `dir/dir/dir/log.testlog.2`}),
	createRollingSizeFileWriterTestCase([]string{}, `./log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.1`}),
	createRollingSizeFileWriterTestCase([]string{`././././log.testlog.9`}, `log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.10`}),
	createRollingSizeFileWriterTestCase([]string{"dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}, "dir/dir/log.testlog", 10, 1, 2, []string{"dir/dir/log.testlog", "dir/dir/log.testlog.1", "dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}),
	createRollingSizeFileWriterTestCase([]string{}, `././dir/log.testlog`, 10, 10, 1, []string{`dir/log.testlog`}),
	createRollingSizeFileWriterTestCase([]string{}, `././dir/log.testlog`, 10, 10, 2, []string{`dir/log.testlog`, `dir/log.testlog.1`}),
	createRollingSizeFileWriterTestCase([]string{`././dir/dir/log.testlog.1`}, `dir/dir/log.testlog`, 10, 10, 2, []string{`dir/dir/log.testlog`, `dir/dir/log.testlog.1`, `dir/dir/log.testlog.2`}),
	createRollingSizeFileWriterTestCase([]string{`././dir/dir/dir/log.testlog.1`}, `dir/dir/dir/log.testlog`, 10, 1, 2, []string{`dir/dir/dir/log.testlog`, `dir/dir/dir/log.testlog.2`}),
	createRollingSizeFileWriterTestCase([]string{}, `././log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.1`}),
	createRollingSizeFileWriterTestCase([]string{`././././log.testlog.9`}, `log.testlog`, 10, 1, 2, []string{`log.testlog`, `log.testlog.10`}),
	createRollingSizeFileWriterTestCase([]string{"././dir/dir/log.testlog.a", "././dir/dir/log.testlog.1b"}, "dir/dir/log.testlog", 10, 1, 2, []string{"dir/dir/log.testlog", "dir/dir/log.testlog.1", "dir/dir/log.testlog.a", "dir/dir/log.testlog.1b"}),
	// ====================
	//createRollingDatefileWriterTestCase([]string{}, "log.txt", "02.01.2006", 1, []string{}),
	//createRollingDatefileWriterTestCase([]string{}, "log.txt", "02.01.2006.000000", 2, []string{}),
}
