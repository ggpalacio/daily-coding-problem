package problem_399

import (
	"math"
)

const (
	numberOfPartitions = 3
)

func Solution(numbers []uint) [][]uint {
	sumTotalValue := sumNumbers(numbers)
	if int(sumTotalValue)%numberOfPartitions != 0 {
		return nil
	}

	found := true
	sumPartitionValue := sumTotalValue / numberOfPartitions
	partitions := createPartitions(numbers, numberOfPartitions)

	for index, partition := range partitions {
		if index < len(partitions)-1 {
			if partition.Total > sumPartitionValue {
				for partition.Total > sumPartitionValue {
					if item, ok := partition.RemoveLast(); ok {
						partitions[index+1].Prepend(item)
					}
				}
			} else if partition.Total < sumPartitionValue {
				for partition.Total < sumPartitionValue {
					if item, ok := partitions[index+1].RemoveFirst(); ok {
						partition.Append(item)
					}
				}
			}
		}
		if partition.Total != sumPartitionValue {
			found = false
			break
		}
	}
	if !found {
		return nil
	}

	result := make([][]uint, len(partitions))
	for index, partition := range partitions {
		result[index] = partition.Items
	}
	return result
}

func createPartitions(numbers []uint, numberOfPartitions int) []*Partition {
	partitions := make([]*Partition, numberOfPartitions)
	partitionLength := int(math.Ceil(float64(len(numbers)) / float64(numberOfPartitions)))
	for partitionNumber := 0; partitionNumber < numberOfPartitions; partitionNumber++ {
		partition := NewPartition(numbers, partitionLength, partitionNumber)
		partitions[partitionNumber] = &partition
	}
	return partitions
}

func sumNumbers(numbers []uint) uint {
	var total uint
	for _, number := range numbers {
		total += number
	}
	return total
}
