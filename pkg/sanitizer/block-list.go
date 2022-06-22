package sanitizer

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

// BlockList is a block-list structure
type BlockList struct {
	location string
	filter   *bloom.BloomFilter
}

// NewBlockList create new block-list from a text file
func NewBlockList(location string) (*BlockList, error) {
	if location == "" {
		return nil, nil
	}
	// open block-list file
	input, err := open(location)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	// count numer of lines
	var buf bytes.Buffer
	tee := io.TeeReader(input, &buf)
	size, err := countLines(tee)
	if err != nil {
		return nil, err
	}

	// initialize bloom filter
	filter := bloom.NewWithEstimates(size, 0.01)

	// read block-list file
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		filter.AddString(strings.TrimPrefix(line, "0.0.0.0 "))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &BlockList{
		location: location,
		filter:   filter,
	}, nil
}

// Contains check if the provided value is inside the block list
func (bl *BlockList) Contains(value string) bool {
	return bl.filter.TestString(value)
}

// Location of the block list
func (bl *BlockList) Location() string {
	return bl.location
}

// Size of the block list
func (bl *BlockList) Size() uint32 {
	return bl.filter.ApproximatedSize()
}
