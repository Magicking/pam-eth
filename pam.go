// +build darwin linux

/*
Copyright (c) 2017 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

// code in here can't be tested because it relies on cgo. :(

import (
	"fmt"
	"log"
	"unsafe"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

char *string_from_argv(int, char**);
char *get_user(pam_handle_t *pamh);
int get_uid(char *user);
*/
import "C"

func init() {
	if !disablePtrace() {
		pamLog("unable to disable ptrace")
	}
}

func sliceFromArgv(argc C.int, argv **C.char) []string {
	r := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {
		s := C.string_from_argv(C.int(i), argv)
		defer C.free(unsafe.Pointer(s))
		r = append(r, C.GoString(s))
	}
	return r
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	ph, err := initEth(sliceFromArgv(argc, argv))
	if err != nil {
		pamLog(err.Error())
		return C.PAM_AUTH_ERR
	}
	// Connect first as it may take some time
	cUsername := C.get_user(pamh)
	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	uId := int64(C.get_uid(cUsername))
	if uId < 0 {
		return C.PAM_USER_UNKNOWN
	}

	username := C.GoString(cUsername)
	token, bNum, err := ph.GetOTP(uId, username)
	if err != nil {
		pamLog(err.Error())
		return C.PAM_AUTH_ERR
	}

	msg := fmt.Sprintf("Helper link: http://127.0.0.1:8000#token=%s&b=%v\nSign %s\n", token, bNum, token)
	tokenSig, err := Conversation(pamh, msg)
	if err != nil {
		pamLog(err.Error())
		return C.PAM_AUTH_ERR
	}
	log.Println("TODO CHECK SIG => ", tokenSig)
	ok, err := ph.VerifyAuth(uId, username, tokenSig)
	if err != nil {
		pamLog(err.Error())
		return C.PAM_AUTH_ERR
	}
	if ok {
		return C.PAM_SUCCESS
	} else {
		return C.PAM_USER_UNKNOWN
	}
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	return C.PAM_SUCCESS
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	//TODO
	return C.PAM_SUCCESS
}

//export pam_sm_open_session
func pam_sm_open_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	//TODO
	return C.PAM_SUCCESS
}

//export pam_sm_close_session
func pam_sm_close_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	//TODO
	return C.PAM_SUCCESS
}
