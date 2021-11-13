//go:build rp2040
// +build rp2040

package stack

import (
//"device/arm"
)

/*
func UpdateStackPointer(loc string) {
	var sp int32

	arm.AsmFull(`
	push {r0}
	mov r0,sp
	str r0,{result}
	pop {r0}
`, map[string]interface{}{
		"result": &sp,
	})

	if sp <= MinStackPointer {
		MinStackPointer = sp
		MinStackLocation = loc
	}
}
*/
