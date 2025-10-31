// Package trid provides file type identification compatible with TrID definition files.
// This is a clean-room Go implementation, independent of the original Python code.
//
// Based on TrID concept by Marco Pontello - https://mark0.net
// Go implementation Copyright (c) 2025 JMcrafter26
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Based on version 2.43 of TrID by Marco Pontello - https://mark0.net

//go:build ignore

package trid

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"sort"
	"strings"
)

const (
	HeaderFrontSize = 2048
	MaxFileSize     = 10 * 1024 * 1024
)

// FileType represents an identified file type
type FileType struct {
	Name        string
	Extension   string
	MimeType    string
	Description string
	URL         string
	Remarks     string
	Author      string
	Email       string
	Homepage    string
	FileCount   uint32
}

// Pattern represents a file signature pattern at specific offset
type Pattern struct {
	Offset int
	Data   []byte
}

// Definition represents a complete file type definition
type Definition struct {
	FileType FileType
	Patterns []Pattern
	Strings  [][]byte
}

// Result represents file identification result
type Result struct {
	Confidence float64
	Score      int
	Definition *Definition
}

// Analyzer handles file type identification
type Analyzer struct {
	definitions []*Definition
}

// NewAnalyzer creates a new file analyzer
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		definitions: make([]*Definition, 0),
	}
}

// LoadDefinitions loads TRD definitions from a file
func (a *Analyzer) LoadDefinitions(filename string) error {
	defs, err := loadTrdDefinitions(filename)
	if err != nil {
		return err
	}
	a.definitions = defs
	return nil
}

// LoadDefinitionsFromBytes loads TRD definitions from byte data
func (a *Analyzer) LoadDefinitionsFromBytes(data []byte) error {
	defs, err := parseTrdPackage(data)
	if err != nil {
		return err
	}
	a.definitions = defs
	return nil
}

// AnalyzeFile identifies the file type of the given file
func (a *Analyzer) AnalyzeFile(filename string) ([]Result, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return a.AnalyzeReader(file, info.Size())
}

// AnalyzeBytes identifies the file type from byte data
func (a *Analyzer) AnalyzeBytes(data []byte) ([]Result, error) {
	reader := bytes.NewReader(data)
	return a.AnalyzeReader(reader, int64(len(data)))
}

// AnalyzeReader identifies the file type from a reader
func (a *Analyzer) AnalyzeReader(reader io.ReaderAt, size int64) ([]Result, error) {
	// Read header for pattern matching
	headerSize := minInt64(size, HeaderFrontSize)
	header := make([]byte, headerSize)
	_, err := reader.ReadAt(header, 0)
	if err != nil {
		return nil, err
	}

	// Prepare data for string matching if needed
	var fullData []byte
	if size <= MaxFileSize {
		fullData = make([]byte, size)
		_, err := reader.ReadAt(fullData, 0)
		if err != nil {
			return nil, err
		}
		fullData = bytes.ToUpper(fullData)
	}

	return a.analyze(header, fullData, size)
}

// GetSupportedTypes returns all supported file types
func (a *Analyzer) GetSupportedTypes() []FileType {
	types := make([]FileType, len(a.definitions))
	for i, def := range a.definitions {
		types[i] = def.FileType
	}
	return types
}

// GetDefinitionCount returns the number of loaded definitions
func (a *Analyzer) GetDefinitionCount() int {
	return len(a.definitions)
}

func (a *Analyzer) analyze(header, fullData []byte, fileSize int64) ([]Result, error) {
	var results []*Result
	totalScore := 0

	foundCache := make(map[string]bool)
	stopCache := make(map[string]bool)

	for _, def := range a.definitions {
		score := a.calculateDefinitionScore(def, header, fullData, fileSize, foundCache, stopCache)
		if score > 0 {
			totalScore += score
			results = append(results, &Result{
				Score:      score,
				Definition: def,
			})
		}
	}

	// Calculate confidence percentages
	for _, result := range results {
		result.Confidence = float64(result.Score) * 100 / float64(totalScore)
	}

	// Sort by score (descending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	// Convert to non-pointer slice
	finalResults := make([]Result, len(results))
	for i, r := range results {
		finalResults[i] = *r
	}

	return finalResults, nil
}

func (a *Analyzer) calculateDefinitionScore(def *Definition, header, fullData []byte, fileSize int64, foundCache, stopCache map[string]bool) int {
	score := 0

	// Check patterns
	for _, pattern := range def.Patterns {
		if fileSize < int64(pattern.Offset+len(pattern.Data)) {
			return 0
		}

		end := pattern.Offset + len(pattern.Data)
		if end > len(header) {
			return 0
		}

		if !bytes.Equal(pattern.Data, header[pattern.Offset:end]) {
			return 0
		}

		if pattern.Offset == 0 {
			score += len(pattern.Data) * 1000
		} else {
			score += len(pattern.Data)
		}
	}

	// Check strings if patterns matched
	if score > 0 && len(def.Strings) > 0 && fullData != nil {
		for _, str := range def.Strings {
			strKey := string(str)
			if stopCache[strKey] {
				return 0
			}
		}

		for _, str := range def.Strings {
			strKey := string(str)
			if foundCache[strKey] {
				score += len(str) * 500
			} else {
				if bytes.Contains(fullData, str) {
					score += len(str) * 500
					foundCache[strKey] = true
				} else {
					stopCache[strKey] = true
					return 0
				}
			}
		}
	}

	return score
}

// Load default definitions from common locations
func (a *Analyzer) LoadDefaultDefinitions() error {
	possiblePaths := []string{
		"triddefs.trd",
		"/usr/share/trid/triddefs.trd",
		"/usr/local/share/trid/triddefs.trd",
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return a.LoadDefinitions(path)
		}
	}

	return errors.New("no TRD definitions file found in default locations")
}

// Helper function to get the best match from results
func GetBestMatch(results []Result) *Result {
	if len(results) == 0 {
		return nil
	}
	return &results[0]
}

// Helper function to check if results contain a specific file type
func HasFileType(results []Result, fileType string) bool {
	for _, result := range results {
		if strings.EqualFold(result.Definition.FileType.Name, fileType) {
			return true
		}
	}
	return false
}

// TRD file parsing functions
func loadTrdDefinitions(filename string) ([]*Definition, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseTrdPackage(data)
}

func parseTrdPackage(data []byte) ([]*Definition, error) {
	if len(data) < 12 || string(data[0:4])+string(data[8:12]) != "RIFFTRID" {
		return nil, errors.New("invalid TRD package signature")
	}

	pkgLen := binary.LittleEndian.Uint32(data[4:8])
	if int(pkgLen)+8 != len(data) {
		return nil, errors.New("TRD package length mismatch")
	}

	defsNum := binary.LittleEndian.Uint32(data[20:24])
	blockLen := binary.LittleEndian.Uint32(data[28:32])

	defsBlock := data[32 : 32+blockLen]
	var definitions []*Definition

	pos := 0
	for i := 0; i < int(defsNum) && pos < len(defsBlock)-8; i++ {
		chunkID := string(defsBlock[pos : pos+4])
		if chunkID == "DEF " {
			chunkLen := int(binary.LittleEndian.Uint32(defsBlock[pos+4 : pos+8]))
			defBlock := defsBlock[pos+8 : pos+8+chunkLen]
			pos += 8 + chunkLen

			def, err := parseDefinitionBlock(defBlock)
			if err != nil {
				return nil, err
			}
			definitions = append(definitions, def)
		} else {
			chunkLen := int(binary.LittleEndian.Uint32(defsBlock[pos+4 : pos+8]))
			pos += 8 + chunkLen
		}
	}

	return definitions, nil
}

func parseDefinitionBlock(block []byte) (*Definition, error) {
	def := &Definition{}

	pos := 0
	for pos < len(block)-8 {
		chunkID := string(block[pos : pos+4])
		chunkLen := int(binary.LittleEndian.Uint32(block[pos+4 : pos+8]))
		chunk := block[pos+8 : pos+8+chunkLen]
		pos += 8 + chunkLen

		switch chunkID {
		case "DATA":
			if err := parseDataChunk(chunk, def); err != nil {
				return nil, err
			}
		case "INFO":
			if err := parseInfoChunk(chunk, &def.FileType); err != nil {
				return nil, err
			}
		}
	}

	return def, nil
}

func parseDataChunk(chunk []byte, def *Definition) error {
	pos := 0
	for pos < len(chunk)-8 {
		subChunkID := string(chunk[pos : pos+4])
		subChunkLen := int(binary.LittleEndian.Uint32(chunk[pos+4 : pos+8]))
		subChunk := chunk[pos+8 : pos+8+subChunkLen]
		pos += 8 + subChunkLen

		switch subChunkID {
		case "PATT":
			patterns, err := parsePatternsBlock(subChunk)
			if err != nil {
				return err
			}
			def.Patterns = patterns
		case "STRN":
			strs, err := parseStringsBlock(subChunk)
			if err != nil {
				return err
			}
			def.Strings = strs
		}
	}
	return nil
}

func parseInfoChunk(chunk []byte, fileType *FileType) error {
	pos := 0
	for pos < len(chunk)-6 {
		infoType := string(chunk[pos : pos+4])
		infoLen := int(binary.LittleEndian.Uint16(chunk[pos+4 : pos+6]))
		infoText := chunk[pos+6 : pos+6+infoLen]
		pos += 6 + infoLen

		switch infoType {
		case "TYPE":
			fileType.Name = string(infoText)
		case "EXT ":
			fileType.Extension = string(infoText)
		case "MIME":
			fileType.MimeType = string(infoText)
		case "NAME":
			fileType.Description = string(infoText)
		case "FNUM":
			fileType.FileCount = binary.LittleEndian.Uint32(infoText)
		case "RURL":
			fileType.URL = string(infoText)
		case "USER":
			fileType.Author = string(infoText)
		case "MAIL":
			fileType.Email = string(infoText)
		case "HOME":
			fileType.Homepage = string(infoText)
		case "REM ":
			fileType.Remarks = string(infoText)
		}
	}
	return nil
}

func parsePatternsBlock(chunk []byte) ([]Pattern, error) {
	if len(chunk) < 2 {
		return nil, nil
	}

	patternCount := int(binary.LittleEndian.Uint16(chunk[0:2]))
	patterns := make([]Pattern, 0, patternCount)

	pos := 2
	for i := 0; i < patternCount && pos < len(chunk)-4; i++ {
		offset := int(binary.LittleEndian.Uint16(chunk[pos : pos+2]))
		length := int(binary.LittleEndian.Uint16(chunk[pos+2 : pos+4]))
		patternData := chunk[pos+4 : pos+4+length]
		pos += 4 + length

		patterns = append(patterns, Pattern{
			Offset: offset,
			Data:   patternData,
		})
	}

	return patterns, nil
}

func parseStringsBlock(chunk []byte) ([][]byte, error) {
	if len(chunk) < 2 {
		return nil, nil
	}

	stringCount := int(binary.LittleEndian.Uint16(chunk[0:2]))
	strs := make([][]byte, 0, stringCount)

	pos := 2
	for i := 0; i < stringCount && pos < len(chunk)-4; i++ {
		length := int(binary.LittleEndian.Uint32(chunk[pos : pos+4]))
		stringData := chunk[pos+4 : pos+4+length]
		pos += 4 + length

		strs = append(strs, stringData)
	}

	return strs, nil
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
