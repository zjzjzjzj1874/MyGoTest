package base

import (
	"context"
	"strconv"

	"golang.org/x/sys/unix"
)

type UsageStat struct {
	Path              string
	Fstype            interface{}
	Total             uint64
	Free              uint64
	InodesTotal       uint64
	InodesFree        uint64
	InodesUsed        uint64
	Used              uint64
	InodesUsedPercent float64
	UsedPercent       float64
}

// Usage returns a file system usage. path is a filesystem path such
// as "/", not device file path like "/dev/vda1".  If you want to use
// a return value of disk.Partitions, use "Mountpoint" not "Device".
func Usage(path string) (*UsageStat, error) {
	return UsageWithContext(context.Background(), path)
}

func UsageWithContext(ctx context.Context, path string) (*UsageStat, error) {
	stat := unix.Statfs_t{}
	err := unix.Statfs(path, &stat)
	if err != nil {
		return nil, err
	}
	bsize := stat.Bsize

	ret := &UsageStat{
		Path:        unescapeFstab(path),
		Fstype:      stat,
		Total:       stat.Blocks * uint64(bsize),
		Free:        stat.Bavail * uint64(bsize),
		InodesTotal: stat.Files,
		InodesFree:  stat.Ffree,
	}

	// if could not get InodesTotal, return empty
	if ret.InodesTotal < ret.InodesFree {
		return ret, nil
	}

	ret.InodesUsed = ret.InodesTotal - ret.InodesFree
	ret.Used = (uint64(stat.Blocks) - uint64(stat.Bfree)) * uint64(bsize)

	if ret.InodesTotal == 0 {
		ret.InodesUsedPercent = 0
	} else {
		ret.InodesUsedPercent = (float64(ret.InodesUsed) / float64(ret.InodesTotal)) * 100.0
	}

	if (ret.Used + ret.Free) == 0 {
		ret.UsedPercent = 0
	} else {
		// We don't use ret.Total to calculate percent.
		// see https://github.com/shirou/gopsutil/issues/562
		ret.UsedPercent = (float64(ret.Used) / float64(ret.Used+ret.Free)) * 100.0
	}

	return ret, nil
}

// Unescape escaped octal chars (like space 040, ampersand 046 and backslash 134) to their real value in fstab fields issue#555
func unescapeFstab(path string) string {
	escaped, err := strconv.Unquote(`"` + path + `"`)
	if err != nil {
		return path
	}
	return escaped
}
