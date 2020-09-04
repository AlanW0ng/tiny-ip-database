package ipdb

import (
	"fmt"
	"math"
)

func (seg IpSegments) Len() int {
	return len(seg)
}

func (seg IpSegments) Less(i, j int) bool {
	return seg[i].Start < seg[j].Start
}

func (seg IpSegments) Swap(i, j int) {
	seg[i], seg[j] = seg[j], seg[i]
}

func (ipSegs *IpSegments) Find(ip uint32) (idx uint32, exists bool) {
	// 使用二分法查找
	length := len(*ipSegs)
	if length < 1 {
		return idx, false
	}
	left, right, mid := 0, length-1, 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		midIpSeg := (*ipSegs)[mid]
		if midIpSeg.Start <= ip && midIpSeg.End >= ip {
			return uint32(mid), true
		} else if midIpSeg.Start > ip {
			// 如果当前元素大于ip，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else {
			// (*ipSegs)[mid].End < ip
			// 如果当前元素小于ip，那么把left指针移到mid + 1的位置
			left = mid + 1
		}
		// 判断如果left大于right，那么这个元素是不存在的
		if left > right {
			return idx, false
		}
	}
}

func (ipSegs *IpSegments) update(idx uint32, ipSeg IpSegment) error {
	// 插入一个元素到Slice
	*ipSegs = append(*ipSegs, IpSegment{})
	copy((*ipSegs)[idx+1:], (*ipSegs)[idx:])
	(*ipSegs)[idx] = ipSeg
	return nil
}

func (ipSegs *IpSegments) delete(idx uint32) error {
	// 从Slice删除一个元素
	*ipSegs = append((*ipSegs)[:idx], (*ipSegs)[idx+1:]...)
	return nil
}

func (ipSegs *IpSegments) Update(ip uint32) error {
	// 使用二分法定位包含该IP的IP段或者最接近的IP段
	length := len(*ipSegs)
	if length < 1 {
		newIpSeg := IpSegment{
			Start: ip,
			End:   ip,
		}
		return ipSegs.update(0, newIpSeg)
	}
	var left, right, mid int
	left, right, mid = 0, len(*ipSegs)-1, 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		midIpSeg := (*ipSegs)[mid]
		if midIpSeg.Start <= ip && midIpSeg.End >= ip {
			// IP已存在
			return nil
		} else if midIpSeg.Start > ip {
			// 如果当前元素大于ip，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else {
			// (*ipSegs)[mid].End < ip
			// 如果当前元素小于ip，那么把left指针移到mid + 1的位置
			left = mid + 1
		}

		fmt.Println(left, mid, right, left > right)
		// 判断如果left大于right，那么这个元素是不存在的
		if left > right {
			lIpSeg, rIpSeg := &(*ipSegs)[Max(right, 0)], &(*ipSegs)[Min(left, len(*ipSegs)-1)]
			if lIpSeg.End+1 == ip {
				lIpSeg.End += 1
				return nil
			} else if rIpSeg.Start-1 == ip {
				rIpSeg.Start -= 1
				return nil
			} else {
				newIpSeg := IpSegment{
					Start: ip,
					End:   ip,
				}
				if ip > midIpSeg.Start {
					return ipSegs.update(uint32(mid+1), newIpSeg)
				} else {
					return ipSegs.update(uint32(mid), newIpSeg)
				}
			}

		}
	}
}

func (ipSegs *IpSegments) Delete(ip uint32) error {
	idx, exists := ipSegs.Find(ip)
	if !exists {
		return nil
	}
	ipSeg := &(*ipSegs)[idx]
	if ipSeg.Start == ipSeg.End {
		// 删除此ipSeg
		return ipSegs.delete(idx)
	} else if ip == ipSeg.Start {
		ipSeg.Start += 1
		return nil
	} else if ip == ipSeg.End {
		ipSeg.End -= 1
		return nil
	} else {
		// 包含在一个IpSegment中且不在两端
		newIpSeg := IpSegment{
			Start: ip + 1,
			End:   ipSeg.End,
		}
		ipSeg.End = ip - 1
		return ipSegs.update(idx+1, newIpSeg)
	}
}
