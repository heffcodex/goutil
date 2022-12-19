package uarray

import (
	"sort"

	"github.com/heffcodex/goutil/v2/uflag"
	"github.com/heffcodex/goutil/v2/utype"
)

type (
	ClusterID                      interface{ utype.ID }
	ClusterFn[T any, I ClusterID]  func(item T) I
	ClusterSet[T any, I ClusterID] map[I][]T
)

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

func (s ClusterSet[T, I]) SortedClusters(order uflag.Order) [][]T {
	clusterIDs := s.SortedIDs(order)
	clusters := make([][]T, 0, len(s))

	for _, clusterID := range clusterIDs {
		clusters = append(clusters, s[clusterID])
	}

	return clusters
}

func Cluster[T any, I ClusterID](arr []T, fn ClusterFn[T, I]) ClusterSet[T, I] {
	res := make(ClusterSet[T, I])

	for _, item := range arr {
		clusterID := fn(item)
		res[clusterID] = append(res[clusterID], item)
	}

	return res
}
