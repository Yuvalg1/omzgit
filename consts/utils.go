package consts

func TrimRight(name string, width int) string {
	return name[:min(len(name), width)]
}
