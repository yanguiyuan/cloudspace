package excel2tag

import (
	"github.com/xuri/excelize/v2"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
	"io"
	"log"
	"slices"
	"strings"
)

type TagTypeRecord struct {
	ID    uint64
	Value string
}
type TagRecord struct {
	ID     uint64
	TypeID uint64
	Value  string
}
type WordTagLink struct {
	WordID uint64
	TagID  uint64
}
type WordRecord struct {
	ID         uint64
	Value      string
	Paraphrase string
}
type Option interface {
	apply(worker *Worker)
}
type Worker struct {
	file                    *excelize.File
	titleRowIndex           int
	primaryKeyColumn        int
	primaryKeyCommentColumn int
	tagTypeRecord           []TagTypeRecord
	tagRecord               []TagRecord
	wordRecord              []WordRecord
	wordTagLink             []WordTagLink
}

func New(r io.Reader, options ...Option) *Worker {
	f, err := excelize.OpenReader(r)
	if err != nil {
		log.Println(err)
		return nil
	}
	w := &Worker{
		f,
		0,
		1,
		-1,
		make([]TagTypeRecord, 0),
		make([]TagRecord, 0),
		make([]WordRecord, 0),
		make([]WordTagLink, 0),
	}
	for _, option := range options {
		option.apply(w)
	}
	return w
}

type commentColumnIndexOption struct {
	index int
}

func (c commentColumnIndexOption) apply(worker *Worker) {
	worker.primaryKeyCommentColumn = c.index
}

type primaryKeyColumnIndexOption struct {
	index int
}

func (c primaryKeyColumnIndexOption) apply(worker *Worker) {
	worker.primaryKeyColumn = c.index
}
func WithCommentColumnIndex(i int) commentColumnIndexOption {
	return commentColumnIndexOption{i}
}
func WithPrimaryKeyColumnIndex(i int) primaryKeyColumnIndexOption {
	return primaryKeyColumnIndexOption{i}
}
func (w *Worker) Parse() error {

	// 解析标题行

	sheets := w.file.GetSheetList()
	rows, err := w.file.GetRows(sheets[0])
	if err != nil {
		return err
	}
	titles := rows[0]
	// 构建 TagTypeRecord 切片
	var tagTypeRecords []TagTypeRecord
	for col, title := range titles {
		if title == "" || col == w.primaryKeyColumn || col == w.primaryKeyCommentColumn {
			continue
		}
		tagTypeRecords = append(tagTypeRecords, TagTypeRecord{
			ID:    id.One(),
			Value: title,
		})
	}

	// 构建 TagRecord 切片
	var tagRecords []TagRecord
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue // 跳过标题行
		}
		for colIndex, cell := range row {
			if colIndex == w.primaryKeyColumn || cell == "" || colIndex == w.primaryKeyCommentColumn {
				continue
			}
			splitWord := strings.Split(cell, "、")
			for _, word := range splitWord {
				if !slices.ContainsFunc(tagRecords, func(record TagRecord) bool {
					return word == record.Value
				}) {
					tempIndex := colIndex
					if colIndex > w.primaryKeyColumn {
						tempIndex--
					}
					if colIndex > w.primaryKeyCommentColumn && w.primaryKeyCommentColumn >= 0 {
						tempIndex--
					}
					tagRecords = append(tagRecords, TagRecord{
						ID:     id.One(),
						TypeID: tagTypeRecords[tempIndex].ID,
						Value:  word,
					})
				}
			}

		}
	}

	// 构建 WordRecord 切片
	var wordRecords []WordRecord
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue // 跳过标题行
		}
		wordRecord := WordRecord{
			ID:    id.One(),
			Value: row[w.primaryKeyColumn],
		}
		if len(row) > w.primaryKeyCommentColumn && w.primaryKeyCommentColumn >= 0 {
			wordRecord.Paraphrase = row[w.primaryKeyCommentColumn]
		}
		wordRecords = append(wordRecords, wordRecord)
	}

	// 构建 WordTagLink 切片
	var wordTagLinks []WordTagLink
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue // 跳过标题行
		}
		for colIndex, cell := range row {
			if colIndex == w.primaryKeyColumn || cell == "" || colIndex == w.primaryKeyCommentColumn {
				continue
			}
			splitWord := strings.Split(cell, "、")
			wordID := wordRecords[rowIndex-1].ID
			for _, sw := range splitWord {
				tagID := tagRecords[GetTagIndex(tagRecords, sw)].ID //fix err
				wordTagLinks = append(wordTagLinks, WordTagLink{
					WordID: wordID,
					TagID:  tagID,
				})
			}

		}
	}
	w.wordRecord = wordRecords
	w.tagTypeRecord = tagTypeRecords
	w.tagRecord = tagRecords
	w.wordTagLink = wordTagLinks

	return nil
}
func GetTagIndex(records []TagRecord, value string) int {
	for i, record := range records {
		if record.Value == value {
			return i
		}
	}
	return -1
}
func (w *Worker) GetTagTypeList() []TagTypeRecord {
	return w.tagTypeRecord
}
func (w *Worker) GetTagList() []TagRecord {
	return w.tagRecord
}
func (w *Worker) GetWordList() []WordRecord {
	return w.wordRecord
}
func (w *Worker) GetWordTagLink() []WordTagLink {
	return w.wordTagLink
}
