package db

type TableName string

var Table = struct {
	Flags TableName
}{
	Flags: "flags",
}