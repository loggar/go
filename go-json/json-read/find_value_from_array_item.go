package json_sample

func FindFirstPathMatch(arr []Array1Element, v string) string {
	for _, item := range arr {
		if item.Path == v {
			return item.NewPath
		}
	}
	return ""
}
