package main

func coPilotMerge(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of nums1 and nums2
	i, j, k := m-1, n-1, m+n-1

	// While there are still elements in both nums1 and nums2
	for i >= 0 && j >= 0 {
		// If the current element in nums1 is greater than the current element in nums2
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			// Otherwise, the current element in nums2 is greater
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are still elements in nums2
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func chatGPTMerge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}

	// Copy remaining elements from nums2
	for p2 >= 0 {
		nums1[p3] = nums2[p2]
		p2--
		p3--
	}
}
