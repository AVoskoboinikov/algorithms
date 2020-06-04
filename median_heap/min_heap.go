package main

import (
	"math"
)

type MinHeap struct {
	objs []int
}

func (h *MinHeap) Insert(val int) {
	h.objs = append(h.objs, val)
	h.bubbleUp()
}

func (h *MinHeap) GetMin() int {
	min := h.objs[0]

	if len(h.objs) > 1 {
		h.objs[0] = h.objs[len(h.objs)-1]
		h.objs = h.objs[0:len(h.objs)-1]
		h.bubbleDown()
	} else {
		h.objs = []int{}
	}

	return min
}

func (h *MinHeap) WhatIsMin() int {
	return h.objs[0]
}

func (h *MinHeap) bubbleUp() {
	if len(h.objs) == 1 {
		return
	}

	h.validateUp(len(h.objs) - 1)
}

func (h *MinHeap) validateUp(i int) {
	if i == 0 {
		return
	}

	pi := h.getParentIndex(i)

	if !h.doesExist(i) || !h.doesExist(pi) {
		return
	}

	val := h.objs[i]
	pval := h.objs[pi]

	if pval <= val {
		return
	}

	h.objs[i] = pval
	h.objs[pi] = val

	h.validateUp(pi)
}

func (h *MinHeap) bubbleDown()  {
	if len(h.objs) == 1 {
		return
	}

	h.validateDown(0)
}

func (h *MinHeap) validateDown(i int) {
	if !h.doesExist(i) {
		return
	}

	if yes, ni := h.shouldSwapDown(i); yes {
		val := h.objs[i]

		h.objs[i] = h.objs[ni]
		h.objs[ni] = val

		h.validateDown(ni)
	}

	return
}

func (h *MinHeap) shouldSwapDown(i int) (bool, int) {
	val := h.objs[i]

	li := h.getLeftChildIndex(i)
	ri := h.getRightChildIndex(i)

	if !h.doesExist(li) && !h.doesExist(ri) {
		return false, 0
	}

	if h.doesExist(li) && !h.doesExist(ri) {
		if h.objs[li] < val {
			return true, li
		}
	}

	if !h.doesExist(li) && h.doesExist(ri) {
		if h.objs[ri] < val {
			return true, ri
		}
	}

	if h.doesExist(li) && h.doesExist(ri) {
		if h.objs[li] < val || h.objs[ri] < val {
			if h.objs[li] < h.objs[ri] {
				return true, li
			}

			if h.objs[li] > h.objs[ri] {
				return true, ri
			}
		}
	}

	return false, 0
}

func (h *MinHeap) getParentIndex(i int) int {
	i += 1

	if i%2 == 0 {
		return i/2 - 1
	}

	return int(math.Floor(float64(i)/2)) - 1
}

func (h *MinHeap) getLeftChildIndex(i int) int {
	i += 1
	return 2*i - 1
}

func (h *MinHeap) getRightChildIndex(i int) int {
	i += 1
	return 2*i
}

func (h *MinHeap) doesExist(i int) bool {
	return i + 1 <= len(h.objs)
}