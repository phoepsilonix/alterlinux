package boot

import (
	"fmt"
	"log/slog"

	"github.com/Hayao0819/nahi/exutils"
)

func byteToKib(b int64) int64 {
	return b / 1024
}

func mibToKiB(m int64) int64 {
	return m * 1024
}

// TODO: Implement MakeEfiBootImg
func MakeEfiBootImg(dest string, size int64) error {
	sizeMib := mibToKiB((byteToKib(size) + 1024) / 1024)
	slog.Debug("Creating EFI boot image...", "dest", dest, "size", sizeMib)

	mkfs := exutils.CommandWithStdio("mkfs.fat", "-C", "-n", "ALTERISOEFI", dest, fmt.Sprint(sizeMib))
	if err := mkfs.Run(); err != nil {
		return err
	}

	mmd := exutils.CommandWithStdio("mmd", "-i", dest, "::/EFI", "::/EFI/BOOT")
	if err := mmd.Run(); err != nil {
		return err
	}

	return nil
}
