package prompt

import (
	"fmt"
	"strings"
)

type FixedBlocks []FixedBlock

func (s FixedBlocks) String() string {
	var sb strings.Builder
	for _, block := range s {
		sb.WriteString(block.String(block.Key, block.Val))
	}
	return sb.String()
}

type FixedBlock struct {
	Key  string
	Val  string
	Vals []string
}

func (s *FixedBlock) String(key, value string) string {
	vals := s.Vals
	if s.Val != "" {
		vals = append(s.Vals, s.Val)
	}
	var sb strings.Builder
	for _, v := range vals {
		sb.WriteString(s.buildOne(key, v))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s *FixedBlock) buildOne(key, value string) string {
	key = strings.ToUpper(key)
	return fmt.Sprintf("### %s START###\n%s\n### %s END###\n", key, value, key)
}
