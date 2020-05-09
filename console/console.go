package console

import (
	"fmt"
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
)

// define win32 const
const _TRUE = 1
const _FALSE = 0
const _CONSOLE_TEXTMODE_BUFFER = 0x00000001
const _GENERIC_READ = 0x40000000
const _GENERIC_WRITE = 0x80000000
const _STD_OUTPUT_HANDLE = 0xFFFFFFF5

// type alias win32 APIs
type (
	_SHORT    = int16
	_USHORT   = uint16
	_LONG     = int32
	_ULONG    = uint32
	_WORD     = uint16
	_DWORE    = uint32
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
) // end structure definition ///////////////////////////////////////////////////////////////////////

// console
type Console struct {
	screenBuffers      [2]uintptr
	pointer            int
	stdOutHandle       uintptr
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
	// fmt.Printf("stdout:%v %v %v\n", stdOutHandle, lastErr, err)
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

	return &Console{[2]uintptr{handle1, handle2}, 0, stdOutHandle, &newWindowSize, &originalWindowSize}, nil
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
