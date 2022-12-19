package uarray

import (
	"testing"

	"github.com/heffcodex/goutil/v2/uflag"

	"github.com/stretchr/testify/require"
)

func TestCluster(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cluster := Cluster(arr, func(item int) int { return item % 3 })
	require.Equal(t, 3, len(cluster))
	require.ElementsMatch(t, []int{3, 6, 9}, cluster[0])
	require.ElementsMatch(t, []int{1, 4, 7, 10}, cluster[1])
	require.ElementsMatch(t, []int{2, 5, 8}, cluster[2])

	ascID := cluster.SortedIDs(uflag.ASC)
	require.Equal(t, []int{0, 1, 2}, ascID)

	descID := cluster.SortedIDs(uflag.DESC)
	require.Equal(t, []int{2, 1, 0}, descID)

	ascCluster := cluster.SortedClusters(uflag.ASC)
	require.Equal(t, [][]int{{3, 6, 9}, {1, 4, 7, 10}, {2, 5, 8}}, ascCluster)
	
	descCluster := cluster.SortedClusters(uflag.DESC)
	require.Equal(t, [][]int{{2, 5, 8}, {1, 4, 7, 10}, {3, 6, 9}}, descCluster)
}
