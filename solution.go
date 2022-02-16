package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	res = emoji.Sprint("Hello :world_map:!")
	fmt.Println(res)
}
