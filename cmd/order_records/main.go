package main

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/samber/lo"
	"github.com/shanexu/go-quicktest/internal/simpleid"
)

func main() {
	var rows = make(map[int]*Row)
	for i := 0; i < 10; i++ {
		rows[i+1] = &Row{
			ID:    i + 1,
			After: nil,
			Ts:    0,
		}
	}
	rowList2 := lo.Values(rows)
	slices.SortFunc(rowList2, func(a, b *Row) int {
		return a.ID - b.ID
	})
	for i := 0; i < 100; i++ {
		who, after := rand.Intn(10)+1, rand.Intn(11)
		//who, after := c.Who, c.After
		fmt.Printf("Move [%d] after [%d]\n", who, after)
		Move(rows, who, after)
		rowList2 = Move2(rowList2, who, after)
		rowList := Ordered(rows)
		Move3(rows, who, after)
		rowList3 := Ordered3(rows)
		Move4(rows, who, after)
		rowList4 := Ordered4(rows)
		if len(rowList2) != len(rowList) || len(rowList3) != len(rowList) || len(rowList4) != len(rowList) {
			panic("length mismatch")
		}
		for i := 0; i < len(rowList2); i++ {
			if rowList2[i] != rowList[i] {
				panic("mismatch 2")
			}
			if rowList3[i] != rowList[i] {
				panic("mismatch 3")
			}
			if rowList4[i] != rowList[i] {
				panic("mismatch 4")
			}
		}

	}
}

type Row struct {
	ID      int
	After   *int
	Ts      int64
	Rank    int64
	AuxRank int
}

func Move(rows map[int]*Row, who, after int) {
	if who == after {
		return
	}
	whoRow := rows[who]
	children := lo.Filter(lo.MapToSlice(rows, func(id int, item *Row) *Row {
		return item
	}), func(item *Row, index int) bool {
		return item.After != nil && *item.After == who
	})
	if len(children) == 0 {
		whoRow.After = &after
		whoRow.Ts = simpleid.NextID()
		return
	}
	ordered := Ordered(rows)
	idx := slices.IndexFunc(ordered, func(row *Row) bool {
		return row.ID == who
	})
	if idx == 0 {
		lo.ForEach(children, func(item *Row, index int) {
			item.After = &idx
		})
	} else {
		id := ordered[idx-1].ID
		lo.ForEach(children, func(item *Row, index int) {
			item.After = &id
		})
	}
	whoRow.After = &after
	whoRow.Ts = simpleid.NextID()
	return
}

func Ordered(rows map[int]*Row) []*Row {
	zeroRow := &Row{
		ID:    0,
		After: nil,
		Ts:    0,
	}
	zeroRowTuple := &lo.Tuple2[*Row, []*Row]{
		A: zeroRow,
		B: nil,
	}
	tupleRows := lo.MapEntries(rows, func(id int, row *Row) (int, *lo.Tuple2[*Row, []*Row]) {
		var nextRows []*Row
		tuple := lo.T2(row, nextRows)
		return id, &tuple
	})
	for _, tuple := range tupleRows {
		row := tuple.A
		if row.After == nil {
			continue
		}
		var afterRowTuple *lo.Tuple2[*Row, []*Row]
		if *row.After == 0 {
			afterRowTuple = zeroRowTuple
		} else {
			afterRowTuple = tupleRows[*row.After]
		}
		afterRowTuple.B = append(afterRowTuple.B, row)
		slices.SortFunc(afterRowTuple.B, func(a, b *Row) int {
			return int(b.Ts - a.Ts)
		})
	}
	tupleRows[zeroRow.ID] = zeroRowTuple
	roots := []*Row{zeroRow}
	for _, row := range rows {
		if row.After == nil {
			roots = append(roots, row)
		}
	}
	slices.SortFunc(roots, func(a, b *Row) int {
		return a.ID - b.ID
	})
	var result []*Row
	for i, root := range roots {
		if i != 0 {
			result = append(result, root)
		}
		getChildren(root, rows, tupleRows, &result)
	}
	return result
}

func getChildren(row *Row, rows map[int]*Row, tupleRowMap map[int]*lo.Tuple2[*Row, []*Row], result *[]*Row) {
	tuple := tupleRowMap[row.ID]
	if len(tuple.B) == 0 {
		return
	}
	for _, r := range tuple.B {
		*result = append(*result, r)
		getChildren(r, rows, tupleRowMap, result)
	}
}

func Move2(rows []*Row, who, after int) []*Row {
	if who == after {
		return rows
	}
	var whoRow *Row
	rows1 := slices.DeleteFunc(rows, func(item *Row) bool {
		if item.ID == who {
			whoRow = item
			return true
		}
		return false
	})
	if after == 0 {
		return slices.Insert(rows1, 0, whoRow)
	}
	idx := slices.IndexFunc(rows1, func(item *Row) bool {
		return item.ID == after
	})
	return slices.Insert(rows1, idx+1, whoRow)
}

func Move3(rows map[int]*Row, who, after int) {
	if who == after {
		return
	}
	rowList := lo.MapToSlice(rows, func(id int, item *Row) *Row {
		return item
	})
	slices.SortFunc(rowList, func(a, b *Row) int {
		weightA := int64(a.ID<<30) + a.Rank
		weightB := int64(b.ID<<30) + b.Rank
		return int(weightA - weightB)
	})
	whoRow := rows[who]
	if after == 0 {
		row := rowList[0]
		if row.ID == who {
			return
		}
		weight := int64(row.ID<<30) + row.Rank
		whoRow.Rank = weight/2 - int64(whoRow.ID<<30)
		return
	}
	idx := slices.IndexFunc(rowList, func(item *Row) bool {
		return item.ID == after
	})
	afterRow := rows[after]
	afterRowWeight := int64(afterRow.ID<<30) + afterRow.Rank
	if idx == len(rowList)-1 {
		whoRow.Rank = (afterRowWeight+0x3fffffff<<30)/2 - int64(whoRow.ID<<30)
		return
	}
	afterAfterRow := rowList[idx+1]
	if afterAfterRow == whoRow {
		return
	}
	whoRow.Rank = (int64(afterAfterRow.ID<<30)+afterAfterRow.Rank+afterRowWeight)/2 - int64(whoRow.ID<<30)
}

func Ordered3(rows map[int]*Row) []*Row {
	rowList := lo.MapToSlice(rows, func(id int, item *Row) *Row {
		return item
	})
	slices.SortFunc(rowList, func(a, b *Row) int {
		weightA := int64(a.ID<<30) + a.Rank
		weightB := int64(b.ID<<30) + b.Rank
		return int(weightA - weightB)
	})
	return rowList
}

func Move4(rows map[int]*Row, who, after int) {
	if who == after {
		return
	}
	var newRank int
	if after == 0 {
		newRank = 1
	} else {
		afterRow := rows[after]
		afterRowRank := afterRow.ID + afterRow.AuxRank
		newRank = afterRowRank + 1
	}
	whoRow := rows[who]
	whoRow.AuxRank = newRank - whoRow.ID
	for _, row := range rows {
		if row == whoRow {
			continue
		}
		if row.AuxRank+row.ID >= newRank {
			row.AuxRank += 1
		}
	}
}

func Ordered4(rows map[int]*Row) []*Row {
	rowList := lo.MapToSlice(rows, func(id int, item *Row) *Row {
		return item
	})
	slices.SortFunc(rowList, func(a, b *Row) int {
		return (a.ID + a.AuxRank) - (b.ID + b.AuxRank)
	})
	return rowList
}
