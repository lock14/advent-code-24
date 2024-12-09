package day9

import (
	. "advent/util"
	"bufio"
	"fmt"
	"strings"
)

const (
	FreeID     = -1
	FreedUpID  = -2
	SentinalID = -10000
)

type FileInfo struct {
	ID     int
	blocks int
	next   *FileInfo
	prev   *FileInfo
}

func Part1(filename string) int64 {
	ID := 0
	file := true
	fileInfos := sentinel()
	cur := fileInfos
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		n := int(Must(ParseInt64(scanner.Text())))
		if file {
			insertAfter(cur, &FileInfo{ID: ID, blocks: n})
			ID++
		} else {
			insertAfter(cur, &FileInfo{ID: FreeID, blocks: n})
		}
		file = !file
		cur = cur.next
	}

	freeBlock := nextFreeBlock(fileInfos, fileInfos)
	fileBlock := nextFileBlock(fileInfos)
	for freeBlock != nil && fileBlock != nil {
		if freeBlock.blocks > fileBlock.blocks {
			newFreeBlock := &FileInfo{ID: FreeID, blocks: freeBlock.blocks - fileBlock.blocks}
			insertAfter(freeBlock, newFreeBlock)

			freeBlock.ID = fileBlock.ID
			freeBlock.blocks = fileBlock.blocks
			fileBlock.ID = FreedUpID
			fileBlock.blocks = 0

			freeBlock = newFreeBlock
			fileBlock = nextFileBlock(fileBlock)
		} else if freeBlock.blocks < fileBlock.blocks {
			freeBlock.ID = fileBlock.ID
			fileBlock.blocks -= freeBlock.blocks
			freeBlock = nextFreeBlock(freeBlock, fileBlock)
		} else {
			freeBlock.ID = fileBlock.ID
			fileBlock.ID = FreedUpID
			fileBlock.blocks = 0
			fileBlock = nextFileBlock(fileBlock)
			freeBlock = nextFreeBlock(freeBlock, fileBlock)

		}
	}

	var compacted []int
	cur = fileInfos.next
	for cur != fileInfos && cur.ID != FreeID && cur.ID != FreedUpID {
		for i := 0; i < cur.blocks; i++ {
			compacted = append(compacted, cur.ID)
		}
		cur = cur.next
	}

	var checksum int64
	for i := 0; i < len(compacted); i++ {
		checksum += int64(i * compacted[i])
	}

	return checksum
}

func Part2(filename string) int64 {
	return 0
}

func sentinel() *FileInfo {
	f := &FileInfo{ID: SentinalID}
	f.next = f
	f.prev = f
	return f
}

func insertAfter(infos *FileInfo, f *FileInfo) {
	f.prev = infos
	f.next = infos.next
	infos.next.prev = f
	infos.next = f
}

func nextFreeBlock(start, end *FileInfo) *FileInfo {
	cur := start.next
	for cur != end {
		if cur.ID == FreeID {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func nextFileBlock(fileInfo *FileInfo) *FileInfo {
	cur := fileInfo.prev
	for cur != fileInfo {
		if cur.ID != FreeID && cur.ID != SentinalID && cur.ID != FreedUpID {
			return cur
		}
		cur = cur.prev
	}
	return nil
}

func (f *FileInfo) String() string {
	var vals []string
	cur := f.next
	for cur != f {
		vals = append(vals, fmt.Sprintf("%+v", *cur))
		cur = cur.next
	}
	return "[" + strings.Join(vals, ", ") + "]"
}
