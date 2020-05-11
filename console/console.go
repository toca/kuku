package console

import (
	"fmt"
	"kuku/operation"
	"syscall"
	"unsafe"
)

// prepare win32APIs  @see http://cuto.unirita.co.jp/gostudy/post/go-windowsapi/
func loadDll(name string) *syscall.DLL {
	dll, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	return dll
}
func findProc(dll *syscall.DLL, name string) *syscall.Proc {
	proc, err := dll.FindProc(name)
	if err != nil {
		panic(err)
	}
	return proc
}

// dll
var kernel32Dll = loadDll("kernel32.dll")

// functions
var (
	CreateConsoleScreenBuffer    = findProc(kernel32Dll, "CreateConsoleScreenBuffer")
	SetConsoleActiveScreenBuffer = findProc(kernel32Dll, "SetConsoleActiveScreenBuffer")
	WriteConsole                 = findProc(kernel32Dll, "WriteConsoleW")
	CloseHandle                  = findProc(kernel32Dll, "CloseHandle")
	SetConsoleCursorPosition     = findProc(kernel32Dll, "SetConsoleCursorPosition")
	SetConsoleWindowInfo         = findProc(kernel32Dll, "SetConsoleWindowInfo")
	GetStdHandle                 = findProc(kernel32Dll, "GetStdHandle")
	SetConsoleScreenBufferSize   = findProc(kernel32Dll, "SetConsoleScreenBufferSize")
	GetConsoleScreenBufferInfoEx = findProc(kernel32Dll, "GetConsoleScreenBufferInfoEx")
	ReadConsoleInput             = findProc(kernel32Dll, "ReadConsoleInputW")
)

// define win32 const
const _TRUE = 1
const _FALSE = 0
const _CONSOLE_TEXTMODE_BUFFER = 0x00000001
const _GENERIC_READ = 0x40000000
const _GENERIC_WRITE = 0x80000000
const _STD_OUTPUT_HANDLE = 0xFFFFFFF5
const _STD_INPUT_HANDLE = 0xFFFFFFFF6
const _KEY_EVENT = 0x0001

type VirtualKeyCode uint16

const (
	VK_NULL  VirtualKeyCode = 0x00
	VK_LEFT  VirtualKeyCode = 0x25
	VK_UP    VirtualKeyCode = 0x26
	VK_RIGHT VirtualKeyCode = 0x27
	VK_DOWN  VirtualKeyCode = 0x28
)

// type alias win32 APIs
type (
	_SHORT    = int16
	_USHORT   = uint16
	_LONG     = int32
	_ULONG    = uint32
	_WORD     = uint16
	_DWORD    = uint32
	_BOOL     = int32
	_COLORREF = uint32
)

// sutructure for win32 APIs ///////////////////////////////////////////////////////////////////////
type (
	_SMALL_RECT struct {
		left   _SHORT
		top    _SHORT
		right  _SHORT
		bottom _SHORT
	}
	_COORD struct {
		x _SHORT
		y _SHORT
	}
	// https://docs.microsoft.com/ja-jp/windows/console/console-screen-buffer-infoex
	_CONSOLE_SCREEN_BUFFER_INFOEX struct {
		cbSize               _ULONG
		dwSize               _COORD
		dwCursorPosition     _COORD
		wAttributes          _SHORT
		srWindow             _SMALL_RECT
		dwMaximumWindowSize  _COORD
		wPopupAttributes     _WORD
		bFullscreenSupported _BOOL
		ColorTable           [16]_COLORREF
	}
	_INPUT_RECORD struct {
		EventType   _WORD
		_padding    _WORD
		EventRecord [16]byte
		// union {
		//   keyEvent KEY_EVENT_RECORD          ;
		//   MOUSE_EVENT_RECORD        MouseEvent;
		//   WINDOW_BUFFER_SIZE_RECORD WindowBufferSizeEvent;
		//   MENU_EVENT_RECORD         MenuEvent;
		//   FOCUS_EVENT_RECORD        FocusEvent;
		// } Event;
	}
	_KEY_EVENT_RECORD struct {
		bKeyDown         _BOOL
		wRepeatCount     _WORD
		wVirtualKeyCode  _WORD
		wVirtualScanCode _WORD
		char             [2]byte
		// union {
		//   WCHAR UnicodeChar;
		//   CHAR  AsciiChar;
		// } uChar;
		dwControlKeyState _DWORD
	}
) // end structure definition ///////////////////////////////////////////////////////////////////////

// console struct
type Console struct {
	screenBuffers      [2]uintptr
	pointer            int
	stdOutHandle       uintptr
	stdInHandle        uintptr
	windowRect         *_SMALL_RECT
	originalWindowSize *_SMALL_RECT
}

// ctor
func NewConsole(width int, height int) (*Console, error) {
	// TODO error handling

	// rect for resize
	newWindowSize := _SMALL_RECT{0, 0, int16(width), int16(height)}

	// std out handle
	stdOutHandle, lastErr, err := GetStdHandle.Call(_STD_OUTPUT_HANDLE)
	if lastErr != 0 {
		fmt.Println(err)
		return nil, err
	}
	// std input handle
	stdInHandle, lastErr, err := GetStdHandle.Call(_STD_INPUT_HANDLE)
	if lastErr != 0 {
		fmt.Println(err)
		return nil, err
	}
	// get current window size
	screenBufferInfo := _CONSOLE_SCREEN_BUFFER_INFOEX{}
	screenBufferInfo.cbSize = uint32(unsafe.Sizeof(screenBufferInfo))
	res, lastError, err := GetConsoleScreenBufferInfoEx.Call(stdOutHandle, uintptr(unsafe.Pointer(&screenBufferInfo)))

	if res != _TRUE || lastError != 0 {
		fmt.Println(err)
		return nil, err
	}
	originalWindowSize := screenBufferInfo.srWindow

	// set window size
	_, lastErr, err = SetConsoleWindowInfo.Call(stdOutHandle, _TRUE, uintptr(unsafe.Pointer(&newWindowSize)))
	if lastErr != 0 {
		fmt.Println(err)
		return nil, err
	}

	// screen buffer 1
	handle1, lastErr, err := CreateConsoleScreenBuffer.Call(_GENERIC_READ|_GENERIC_WRITE, 0, 0, _CONSOLE_TEXTMODE_BUFFER, 0)
	if lastErr != 0 {
		fmt.Println(err)
		return nil, err
	}
	// screen buffer 2
	handle2, lastErr, err := CreateConsoleScreenBuffer.Call(_GENERIC_READ|_GENERIC_WRITE, 0, 0, _CONSOLE_TEXTMODE_BUFFER, 0)
	if lastErr != 0 {
		fmt.Println(err)
		return nil, err
	}

	return &Console{[2]uintptr{handle1, handle2}, 0, stdOutHandle, stdInHandle, &newWindowSize, &originalWindowSize}, nil
}

// write to console
func (this *Console) Write(buffer string) {
	this.pointer = (this.pointer + 1) % 2
	this.ResetWindowSize()
	utf16String, err := syscall.UTF16FromString(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	SetConsoleCursorPosition.Call(this.screenBuffers[this.pointer], 0)
	var written uintptr = 0
	_, lastErr, err := WriteConsole.Call(
		this.screenBuffers[this.pointer],
		uintptr(unsafe.Pointer(&utf16String[0])),
		uintptr(len(utf16String)),
		written,
		0)
	if lastErr != 0 {
		fmt.Println(err)
		return
	}
	_, lastErr, err = SetConsoleActiveScreenBuffer.Call(this.screenBuffers[this.pointer])
	if lastErr != 0 {
		fmt.Println(err)
	}
}

// closing
func (this *Console) Close() {
	SetConsoleWindowInfo.Call(this.stdOutHandle, _TRUE, uintptr(unsafe.Pointer(this.originalWindowSize)))
	CloseHandle.Call(this.screenBuffers[0])
	CloseHandle.Call(this.screenBuffers[1])
}

func (this *Console) ResetWindowSize() {
	// res, lastErr, err := SetConsoleWindowInfo.Call(this.screenBuffers[this.pointer], _TRUE, uintptr(unsafe.Pointer(this.windowRect)))
	// if lastErr != 0 || res == 0 {
	// 	fmt.Println(err)
	// }
}

func (this *Console) ReadInput() (bool, operation.KeyInput) {
	record := _INPUT_RECORD{}
	var len _DWORD = 1
	var numOfEvents _DWORD = 0

	res, _, err := ReadConsoleInput.Call(this.stdInHandle, uintptr(unsafe.Pointer(&record)), uintptr(len), uintptr(unsafe.Pointer(&numOfEvents)))
	if res == 0 {
		fmt.Println(err)
		return false, operation.KeyInput{operation.VK_NULL, 0}
	}
	if record.EventType != _KEY_EVENT {
		return false, operation.KeyInput{operation.VK_NULL, 0}
	}
	keyEventRecord := (*_KEY_EVENT_RECORD)(unsafe.Pointer(&record.EventRecord))
	// fmt.Println(record.EventRecord)
	// fmt.Printf("KD:%d, rep:%d, vk:0x%x, sc:0x%x, char:%X %X, state:%x\n",
	// 	keyEventRecord.bKeyDown,
	// 	keyEventRecord.wRepeatCount,
	// 	keyEventRecord.wVirtualKeyCode,
	// 	keyEventRecord.wVirtualScanCode,
	// 	keyEventRecord.char[0], keyEventRecord.char[1],
	// 	keyEventRecord.dwControlKeyState)
	if keyEventRecord.bKeyDown != _TRUE {
		return false, operation.KeyInput{operation.VK_NULL, 0}
	}
	return true, operation.KeyInput{operation.VirtualKeyCode(keyEventRecord.wVirtualKeyCode), int(keyEventRecord.wRepeatCount)}

}

func (this *Console) currentConsoleBuffer() uintptr {
	return this.screenBuffers[this.pointer]
}
