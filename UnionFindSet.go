package unionfindset

import (
	"errors"
	"fmt"
)

/*UnionFindSet structure that allows check path connectivity betwen elements of Set
*	Data structure : weighted tree
*	Uses : path commpresion
 */
type UnionFindSet struct {
	idsArray  []int
	sizeArray []int
}

var errFooInitNumberLessZero = errors.New("number of elements should be greater that 0")

var errFooIndexOutOfRange = errors.New("index out of range")

/*New Constructor for UnionFindSet
*Input
*	n : number of elements in UnionFindSet
*Output
	UnionFindSet
*/
func New(n int) (*UnionFindSet, error) {
	if n < 0 {
		return nil, errFooInitNumberLessZero
	}
	idsArray := make([]int, n)
	sizeArray := make([]int, n)

	for i := 0; i < n; i++ {
		idsArray[i] = i
	}
	return &UnionFindSet{idsArray: idsArray, sizeArray: sizeArray}, nil
}

/*Find function show to which component is belong element
*Input
*	i: element to check
*	ufs: UnionFindSet
*Output
*	index of component
 */
func (ufs *UnionFindSet) Find(i int) (int, error) {
	err := ufs.checkIndexRange(i)
	if err != nil {
		return 0, err
	}
	for i != ufs.idsArray[i] {
		ufs.idsArray[i] = ufs.idsArray[ufs.idsArray[i]] //This line make every node in path point to its grandparent. That keeps tree flat. It compreses path
		i = ufs.idsArray[i]
	}
	return i, nil
}

/*Union create connection between to element of UnionFindSet
*Input
*	p,q : element of UnionFindSet
 */
func (ufs *UnionFindSet) Union(p int, q int) error {
	err := ufs.checkIndexRange(p, q)
	if err != nil {
		return err
	}
	i, _ := ufs.Find(p)
	j, _ := ufs.Find(q)

	if i == j {
		return nil
	}
	// check and add weight to elements and tree of elements. element with less weight connect to element with greater weight
	if ufs.sizeArray[i] < ufs.sizeArray[j] {
		ufs.idsArray[i] = j
		ufs.sizeArray[j] += ufs.sizeArray[i]
	} else {
		ufs.idsArray[j] = i
		ufs.sizeArray[i] += ufs.sizeArray[j]
	}
	return nil
}

/*Connected check connection between two elements
*Input
*	p,q : element of UnionFindSet
*Output
*	is connected elements
 */
func (ufs *UnionFindSet) Connected(p int, q int) (bool, error) {
	err := ufs.checkIndexRange(p, q)
	if err != nil {
		return false, err
	}
	pRoot, _ := ufs.Find(p)
	qRoot, _ := ufs.Find(q)
	return pRoot == qRoot, nil
}

func (ufs *UnionFindSet) checkIndexRange(indexes ...int) error {
	for _, i := range indexes {
		if i >= len(ufs.idsArray) || i < 0 {
			return errFooIndexOutOfRange
		}
	}
	return nil
}

//ShowTree Display tree of Union find Set
func (ufs *UnionFindSet) ShowTree() {
	for i, v := range ufs.idsArray {
		fmt.Println(i, v)
	}
}
