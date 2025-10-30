package Clean

func CleanWorkSpace(workspace *[][]*string) *[][]*string {
	if workspace == nil {
		return workspace
	}
	for i := range *workspace {
		row := (*workspace)[i]
		for j := range row {
			row[j] = nil
		}
	}
	return workspace
}
