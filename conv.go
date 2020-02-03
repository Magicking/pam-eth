// +build darwin linux

package main

/*
#cgo LDFLAGS: -lpam
#include <stdlib.h>
#include <stdio.h>
#include <security/pam_appl.h>

typedef const struct pam_message** MsgsT;
typedef struct pam_response* RespsT;

int do_conv(pam_handle_t* pamh, char* msg, RespsT* resp) {
	int err;
	struct pam_message _msg = { .msg_style = PAM_PROMPT_ECHO_ON, .msg = msg };
	struct pam_message *msgs = &_msg;
	struct pam_conv* conv;

	err = pam_get_item(pamh, PAM_CONV, (const void**)&conv);
	if(err != PAM_SUCCESS) {
		return err;
	}

	return conv->conv(1, (const MsgsT)&msgs, resp, conv->appdata_ptr);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Conversation passes on the specified messages.
func Conversation(pamh *C.pam_handle_t, msg string) (string, error) {
	var resp C.RespsT
	code := C.do_conv(pamh, C.CString(msg), &resp)
	if code != C.PAM_SUCCESS {
		return "", fmt.Errorf("Got non-success from the function: %d", code)
	}
	if resp == nil {
		return "", fmt.Errorf("Empty response")
	}

	ret := C.GoString((*resp).resp)
	C.free(unsafe.Pointer((*resp).resp))

	return ret, nil
}
