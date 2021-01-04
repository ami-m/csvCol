package main

type Filter func(r Record) Record

func buildFilter(p runParams) Filter {
	if (0 == len(p.cols)) && (0 == len(p.colNames)) {
		return buildNopFilter()
	}

	if p.invert {
		return buildInvertFilter(p)
	}

	return buildRegularFilter(p)
}

func buildNopFilter() Filter {
	return func(r Record) Record {
		return r
	}
}

func buildRegularFilter(p runParams) Filter {
	return func(r Record) Record {
		var res Record
		for _, cell := range p.cols {
			if len(r) > cell {
				res = append(res, r[cell])
			} else {
				res = append(res, "")
			}
		}
		return res
	}
}

func buildInvertFilter(p runParams) Filter {
	colDic := make(map[int]bool)
	for _, cell := range p.cols {
		colDic[cell] = true
	}
	return func(r Record) Record {
		var res Record
		for cell, col := range r {
			if !colDic[cell] {
				res = append(res, col)
			}
		}
		return res
	}
}
