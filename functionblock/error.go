package functionblock

import "fmt"

func error(nowFb interface{}) {
	fmt.Println(nowFb.(*EMerge).FbPrivate.(*EMergeAndServiceValue).FbTtl)
}
