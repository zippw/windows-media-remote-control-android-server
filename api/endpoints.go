package api

import (
	"net/http"
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procSendInput    = user32.NewProc("SendInput")
)

const (
	VK_VOLUME_MUTE_KEY = 0xAD
	VK_VOLUME_DOWN_KEY = 0xAE
	VK_VOLUME_UP_KEY   = 0xAF
	VK_MEDIA_NEXT_TRACK_KEY = 0xB0
	VK_MEDIA_PREV_TRACK_KEY = 0xB1
	VK_MEDIA_PLAY_PAUSE_KEY = 0xB3
)

type keyBdInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uintptr
}

type input struct {
	inputType uint32
	ki        keyBdInput
	padding   uint64
}

const (
	inputKeyboard = 1
	keyeventfKeyup = 0x0002
)

func sendKey(vk uint16) {
	var keyDown input
	keyDown.inputType = inputKeyboard
	keyDown.ki.wVk = vk

	var keyUp input
	keyUp.inputType = inputKeyboard
	keyUp.ki.wVk = vk
	keyUp.ki.dwFlags = keyeventfKeyup

	inputs := []input{keyDown, keyUp}	
	procSendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(inputs[0])),
	)
}

func RegisterRoutes() {
	http.HandleFunc("/api/status", AuthMiddleware(statusHandler))
	http.HandleFunc("/api/play", AuthMiddleware(playHandler))
	http.HandleFunc("/api/next", AuthMiddleware(nextHandler))
	http.HandleFunc("/api/prev", AuthMiddleware(prevHandler))
	http.HandleFunc("/api/volup", AuthMiddleware(volUpHandler))
	http.HandleFunc("/api/voldown", AuthMiddleware(volDownHandler))
	http.HandleFunc("/api/mute", AuthMiddleware(muteHandler))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_MEDIA_PLAY_PAUSE_KEY)
	w.WriteHeader(http.StatusOK)
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_MEDIA_NEXT_TRACK_KEY)
	w.WriteHeader(http.StatusOK)
}

func prevHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_MEDIA_PREV_TRACK_KEY)
	w.WriteHeader(http.StatusOK)
}

func volUpHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_VOLUME_UP_KEY)
	w.WriteHeader(http.StatusOK)
}

func volDownHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_VOLUME_DOWN_KEY)
	w.WriteHeader(http.StatusOK)
}

func muteHandler(w http.ResponseWriter, r *http.Request) {
	sendKey(VK_VOLUME_MUTE_KEY)
	w.WriteHeader(http.StatusOK)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}