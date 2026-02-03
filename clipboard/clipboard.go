package clipboard

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Copy(s string) {
	b64 := base64.StdEncoding.EncodeToString([]byte(s))
	seq := fmt.Sprintf("\033]52;c;%s\a", b64)
	os.Stdout.WriteString(seq)
}
