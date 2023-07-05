package uarray

import (
	"sort"

	"github.com/heffcodex/goutil/v2/uflag"
	"github.com/heffcodex/goutil/v2/utype"
)

type (
	// ClusterID is a comparable type for cluster IDs.
	ClusterID interface{ utype.ID }

	// ClusterFn is a function to map element to its ClusterID.
	ClusterFn[T any, I ClusterID] func(item T) I

	// ClusterSet as a map of clusters.
	ClusterSet[T any, I ClusterID] map[I][]T
)

// SortedIDs returns a sorted slice of ClusterID-s by direction provided by the uflag.Order.
func (s ClusterSet[T, I]) SortedIDs(order uflag.Order) []I {
	clusterIDs := make([]I, 0, len(s))

	for clusterID := range s {
		clusterIDs = append(clusterIDs, clusterID)
	}

	sort.Slice(clusterIDs, func(i, j int) bool {
		b := clusterIDs[i] < clusterIDs[j]
		if order == uflag.DESC {
			b = !b
		}

		return b
	})

	return clusterIDs
}

// SortedClusters returns a sorted slice of clusters by direction provided by the uflag.Order.
func (s ClusterSet[T, I]) SortedClusters(order uflag.Order) [][]T {
	clusterIDs := s.SortedIDs(order)
	clusters := make([][]T, 0, len(s))

	for _, clusterID := range clusterIDs {
		clusters = append(clusters, s[clusterID])
	}

	return clusters
}

// Cluster groups items into clusters according to their ids, provided by ClusterFn functions, producing a ClusterSet.
// Argument `fn` can contain multiple functions, each of which will be applied to a single item to produce multiple ClusterID.
func Cluster[T any, I ClusterID](arr []T, fn ...ClusterFn[T, I]) ClusterSet[T, I] {
	res := make(ClusterSet[T, I])

	for _, item := range arr {
		usedIDs := make(map[I]struct{}, len(fn))

		for _, f := range fn {
			clusterID := f(item)

			if _, ok := usedIDs[clusterID]; !ok {
				res[clusterID] = append(res[clusterID], item)
				usedIDs[clusterID] = struct{}{}
			}
		}
	}

	return res
}
