package operation

type VirtualKeyCode uint16

const (
	VK_NULL  VirtualKeyCode = 0x00
	VK_LEFT  VirtualKeyCode = 0x25
	VK_UP    VirtualKeyCode = 0x26
	VK_RIGHT VirtualKeyCode = 0x27
	VK_DOWN  VirtualKeyCode = 0x28
)

type KeyInput struct {
	Key    VirtualKeyCode
	Repeat int
}
