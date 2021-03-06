package Search

import "Algorithm/Sort"

// # 二分探索
// すでに整列済みの集合に対して探索することで逐次探索よりも高速に検索できる。検索の目的は「検索対象が集合に存在するかを調べる」ことであって
// その目的をできるだけ早く達成するために検索対象の集合に前処理を加えることはズルではない。が、前処理にかかる時間等も考慮する必要がある。
// すでに検索対象が整列済みであることを利用すれば、存在するかを調べたい要素がいるはずである集合の部分を決め込んで行くことができる。これに
// よって調査対象が狭まるので無駄な調査が省けるので高速になる。
// 二分探索とは「検索手法」であってこれをサポートするデータ構造は複数考えられる。有名なところでは配列と二分探索木である。
// 集合の要素が変化しないような場合であれば配列を使うのが良い。一方で集合の要素の削除や追加が頻発するような場合は木構造にしたほうが良い。
// これは二分探索の計算量の話ではなくて、整列済みの集合を作るためのオーベーヘッドが採用するデータ構造によって変わってくるからである。
// 二分探索の変形として、「探索して存在しなかったら適切な位置に探索対象を挿入する」というものがある。二分探索の結果の返戻値を対象のインデックス
// として、存在しないものであれば負数で返すなどのように工夫すれば容易に実現できる。が、整列されている配列への挿入の時間コストが大きくなる。

func BinarySearch(a []int, t int, sortFlag bool) bool {
	n := len(a)
	// 整列を探索内に含めるとオーバーヘッドがでかいことがわかる。
	if sortFlag {
		Sort.QuickSort(&a, 0, n-1) // this is my own implementation.
		// sort.Ints(a)
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if t < a[mid] {
			high = mid - 1
		} else if a[mid] < t {
			low = mid + 1
		} else {
			return true // t == a[mid]
		}
	}
	return false
}
