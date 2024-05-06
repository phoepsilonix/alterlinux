package boot

import (
	"fmt"
	"log/slog"

	"github.com/FascodeNet/alterlinux/alteriso5/utils"
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

	mkfs := utils.CommandWithStdio("mkfs.fat", "-C", "-n", "ALTERISO_EFI", dest, fmt.Sprint(sizeMib))
	if err := mkfs.Run(); err != nil {
		return err
	}

	mmd := utils.CommandWithStdio("mmd", "-i", dest, "::/EFI", "::/EFI/BOOT")
	if err := mmd.Run(); err != nil {
		return err
	}

	return nil
}
