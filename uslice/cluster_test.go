package uslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/heffcodex/goutil/v2/uflag"
)

func TestCluster(t *testing.T) {
	t.Parallel()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cluster := Cluster(
		arr,
		func(item int) int { return item % 3 },
		func(item int) int { return item % 2 },
	)

	require.Len(t, cluster, 3)
	require.ElementsMatch(t, []int{2, 3, 4, 6, 8, 9, 10}, cluster[0])
	require.ElementsMatch(t, []int{1, 3, 4, 5, 7, 9, 10}, cluster[1])
	require.ElementsMatch(t, []int{2, 5, 8}, cluster[2])

	ascID := cluster.SortedKeys(uflag.ASC)
	assert.Equal(t, []int{0, 1, 2}, ascID)

	descID := cluster.SortedKeys(uflag.DESC)
	assert.Equal(t, []int{2, 1, 0}, descID)

	ascCluster := cluster.SortedClusters(uflag.ASC)
	assert.Equal(t, [][]int{{2, 3, 4, 6, 8, 9, 10}, {1, 3, 4, 5, 7, 9, 10}, {2, 5, 8}}, ascCluster)

	descCluster := cluster.SortedClusters(uflag.DESC)
	assert.Equal(t, [][]int{{2, 5, 8}, {1, 3, 4, 5, 7, 9, 10}, {2, 3, 4, 6, 8, 9, 10}}, descCluster)
}
