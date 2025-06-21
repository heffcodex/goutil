package uslice

import (
	"cmp"
	"maps"
	"slices"

	"github.com/heffcodex/goutil/v2/uflag"
)

type (
	// ClusterKey is a comparable type for cluster keys.
	ClusterKey interface{ cmp.Ordered }

	// ClusterSet as a map of clusters.
	ClusterSet[K ClusterKey, V any, C ~[]V] map[K]C
)

// SortedKeys returns a sorted slice of cluster keys by direction provided by the uflag.Order.
func (s ClusterSet[K, V, C]) SortedKeys(order uflag.Order) []K {
	keys := maps.Keys(s)

	switch order {
	case uflag.DESC:
		return slices.SortedFunc(keys, func(a, b K) int {
			return cmp.Compare(b, a)
		})
	case uflag.ASC:
		fallthrough
	default:
		return slices.SortedFunc(keys, cmp.Compare)
	}
}

// SortedClusters returns a sorted slice of clusters by direction provided by the uflag.Order.
func (s ClusterSet[K, V, C]) SortedClusters(order uflag.Order) []C {
	keys := s.SortedKeys(order)
	clusters := make([]C, len(keys))

	for i, k := range keys {
		clusters[i] = s[k]
	}

	return clusters
}

// Cluster groups elements of given slice into clusters according to their keys, provided by KeyFn functions, producing a ClusterSet.
// Each given keyFn will be applied to each item thus you can produce multiple keys for same element.
func Cluster[K ClusterKey, V any, S ~[]V](s S, keyFn ...KeyFn[K, V]) ClusterSet[K, V, S] {
	res := make(ClusterSet[K, V, S])

	for _, item := range s {
		usedKeys := make(map[K]struct{}, len(keyFn))

		for _, fn := range keyFn {
			key := fn(item)

			if _, ok := usedKeys[key]; !ok {
				res[key] = append(res[key], item)
				usedKeys[key] = struct{}{}
			}
		}
	}

	return res
}
