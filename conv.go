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

// MessageStyle is a style of Message
type MessageStyle int

const (
	// MessageEchoOff is for messages that shouldn't gave an echo.
	MessageEchoOff = MessageStyle(C.PAM_PROMPT_ECHO_OFF)

	// MessageEchoOn is for messages that should have an echo.
	MessageEchoOn = MessageStyle(C.PAM_PROMPT_ECHO_ON)

	// MessageErrorMsg is for messages that should be displayed as an error.
	MessageErrorMsg = MessageStyle(C.PAM_ERROR_MSG)

	// MessageTextInfo is for textual blurbs to be spat out.
	MessageTextInfo = MessageStyle(C.PAM_TEXT_INFO)
)

// Message represents something to ask / show in a Conv.Conversation call.
type Message struct {
	Style MessageStyle
	Msg   string
}

// Conversation passes on the specified messages.
func Conversation(pamh *C.pam_handle_t, msg string) ([]string, error) {
	var resp C.RespsT
	code := C.do_conv(pamh, C.CString(msg), &resp)
	if code != C.PAM_SUCCESS {
		return nil, fmt.Errorf("Got non-success from the function: %d", code)
	}
	if resp == nil {
		return nil, fmt.Errorf("Empty response")
	}

	var ret []string
	ret = append(ret, C.GoString((*resp).resp))
	C.free(unsafe.Pointer((*resp).resp))

	return ret, nil
}
