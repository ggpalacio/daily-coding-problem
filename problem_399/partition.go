package problem_399

type Partition struct {
	Items []uint
	Total uint
}

func NewPartition(numbers []uint, partitionLength, partitionNumber int) Partition {
	firstIndex := partitionLength * partitionNumber
	lastIndex := partitionLength * (partitionNumber + 1)
	if lastIndex > len(numbers) {
		lastIndex = len(numbers)
	}
	partition := Partition{
		Items: numbers[firstIndex:lastIndex],
	}
	partition.Total = partition.Sum()
	return partition
}

func (ref *Partition) RemoveFirst() (uint, bool) {
	firstItem := ref.Items[0]
	if len(ref.Items) == 0 {
		return 0, false
	} else if len(ref.Items) == 1 {
		ref.Items = []uint{}
	} else {
		ref.Items = ref.Items[1:]
	}
	ref.Total -= firstItem
	return firstItem, true
}

func (ref *Partition) RemoveLast() (uint, bool) {
	lastItem := ref.Items[len(ref.Items)-1]
	if len(ref.Items) == 0 {
		return 0, false
	} else if len(ref.Items) == 1 {
		ref.Items = []uint{}
	} else {
		ref.Items = ref.Items[:len(ref.Items)-1]
	}
	ref.Total -= lastItem
	return lastItem, true
}

func (ref *Partition) Append(item uint) {
	ref.Items = append(ref.Items, item)
	ref.Total += item
}

func (ref *Partition) Prepend(item uint) {
	ref.Items = append([]uint{item}, ref.Items...)
	ref.Total += item
}

func (ref *Partition) Sum() uint {
	var total uint
	for _, number := range ref.Items {
		total += number
	}
	return total
}
