package mailslot

import (
	"fmt"
	"unsafe"
)

const (
	MAILSLOT_WAIT_FOREVER int32 = -1
	MAILSLOT_NO_MESSAGE   int32 = -1

	GENERIC_READ          = 0x80000000
	GENERIC_WRITE         = 0x40000000
	FILE_SHARE_READ       = 0x00000001
	FILE_SHARE_WRITE      = 0x00000002
	OPEN_EXISTING         = 3
	FILE_ATTRIBUTE_NORMAL = 0x00000080
)

type MailSlot struct {
	handle  uintptr
	timeout int32
}

type MailSlotFile struct {
	handle uintptr
}

type MailSlotInfo struct {
	MaxSize     int32
	NextSize    int32
	Count       int32
	ReadTimeout int32
}

func (info MailSlotInfo) String() string {
	return fmt.Sprint("MaxMessageSize: ", info.MaxSize, ", NextSize: ", info.NextSize, ", Count: ", info.Count, ", ReadTimeout: ", info.ReadTimeout)
}

func New(name string, max int32, timeout int32) (*MailSlot, error) {
	handle, _, err := CreateMailslot(name, max, timeout)
	if handle == 0 {
		return nil, err
	}

	return &MailSlot{
		handle:  handle,
		timeout: timeout,
	}, nil
}

func (ms *MailSlot) Info() (MailSlotInfo, error) {
	info := MailSlotInfo{}

	ok, _, err := GetMailslotInfo(ms.handle,
		uintptr(unsafe.Pointer(&info.MaxSize)),
		uintptr(unsafe.Pointer(&info.NextSize)),
		uintptr(unsafe.Pointer(&info.Count)),
		uintptr(unsafe.Pointer(&info.ReadTimeout)))

	if !ok {
		return info, err
	}
	return info, nil
}

func (ms *MailSlot) SetTimeout(timeout int32) error {
	ok, _, err := SetMailslotInfo(ms.handle, timeout)
	if !ok {
		return err
	}
	ms.timeout = timeout
	return nil
}

func (ms *MailSlot) Read(p []byte) (n int, err error) {
	return ReadFile(ms.handle, p)
}

func (ms *MailSlot) Close() error {
	ok, _, err := CloseHandle(ms.handle)
	if ok {
		return nil
	}
	return err
}

func Open(name string) (*MailSlotFile, error) {
	handle, err := CreateFile(name, GENERIC_READ|GENERIC_WRITE, FILE_SHARE_READ|FILE_SHARE_WRITE, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL)
	if int(handle) == -1 {
		return nil, err
	}
	return &MailSlotFile{handle: handle}, nil
}

func (ms *MailSlotFile) Write(p []byte) (n int, err error) {
	return WriteFile(ms.handle, p)
}

func (ms *MailSlotFile) Close() error {
	ok, _, err := CloseHandle(ms.handle)
	if ok {
		return nil
	}
	return err
}
