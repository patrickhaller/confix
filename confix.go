/* Copy all the members of your struct that match a prefix
   into their struct.

   https://stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go
   https://stackoverflow.com/questions/6395076/using-reflect
*/

package confix

import (
	"reflect"
)

// Confix grafts one config struct onto another
func Confix(prefix string, src interface{}, dst interface{}) {
	_a := reflect.ValueOf(src)
	a := _a.Elem()
	_b := reflect.ValueOf(dst)
	b := _b.Elem()

	for i := 0; i < a.NumField(); i++ {
		akey := a.Type().Field(i).Name
		aval := a.Field(i)
		for j := 0; j < b.NumField(); j++ {
			if b.Field(j).CanSet() == false {
				continue
			}
			bkey := b.Type().Field(j).Name
			if akey == prefix+bkey {
				b.Field(j).Set(aval)
			}
		}
	}
}
