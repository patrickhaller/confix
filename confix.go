/* Copy all the members of your struct that match a prefix
   into their struct.

   https://stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go
   https://stackoverflow.com/questions/6395076/using-reflect
*/

package confix

import (
	"fmt"
	"reflect"
)

// Confix grafts one config struct onto another
func Confix(prefix string, src interface{}, dst interface{}) error {
	_a := reflect.ValueOf(src)
	a := _a.Elem()
	_b := reflect.ValueOf(dst)
	b := _b.Elem()

	for i := 0; i < a.NumField(); i++ {
		akey := a.Type().Field(i).Name
		aval := a.Field(i)
		atyp := a.Field(i).Type()
		for j := 0; j < b.NumField(); j++ {
			bkey := b.Type().Field(j).Name
			if akey != prefix+bkey {
				continue
			}
			if b.Field(j).CanSet() == false {
				return fmt.Errorf(
					"Member is unsettable (likely unexported): %s", bkey)
			}
			if atyp != b.Field(j).Type() {
				return fmt.Errorf("Type of %s != %s", akey, bkey)
			}
			b.Field(j).Set(aval)
		}
	}
	return nil
}
