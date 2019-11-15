package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Build int
}

var LatestVersion = Version{1, 0, 0}

func NewVersion(ver string) Version {
	verNums := strings.Split(ver, ".")
	if len(verNums) != 3 {
		return Version{0, 0, 0}
	}

	var err error
	var major, minor, build int
	major, err = strconv.Atoi(verNums[0])
	minor, err = strconv.Atoi(verNums[1])
	build, err = strconv.Atoi(verNums[2])
	if err != nil {
		return Version{0, 0, 0}
	}

	return Version{major, minor, build}
}

func (v1 *Version) EqualTo(v2 *Version) bool {
	return v1.Major == v2.Major && v1.Minor == v2.Minor && v1.Build == v2.Build
}

func (v1 *Version) GreaterThan(v2 *Version) bool {
	if v1.Major != v2.Major {
		return v1.Major > v2.Major
	}
	if v1.Minor != v2.Minor {
		return v1.Minor > v2.Minor
	}
	if v1.Build != v2.Build {
		return v1.Build > v2.Build
	}
	return false
}

func (v1 *Version) GreaterThanOrEqualTo(v2 *Version) bool {
	return v1.GreaterThan(v2) || v1.EqualTo(v2)
}

func (v1 *Version) LessThan(v2 *Version) bool {
	return !v1.GreaterThanOrEqualTo(v2)
}

func (v1 *Version) LessThanOrEqualTo(v2 *Version) bool {
	return !v1.GreaterThan(v2)
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
}
