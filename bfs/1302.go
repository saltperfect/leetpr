package bfs

import(
	"github.com/saltperfect/goleet"
)

func deepestLeavesSum(root *goleet.TreeNode) int {
  var q, nq []*goleet.TreeNode
  nq = []*goleet.TreeNode{root}
  sum := 0
  for len(nq) > 0 {
    q, nq = nq, make([]*goleet.TreeNode, 0)
    sum = 0
    for _, v := range q {
      if v.Left != nil {
        nq = append(nq, v.Left)    
      }
      if v.Right != nil {
        nq = append(nq, v.Right)    
      }
      sum += v.Val
    }
  }
  return sum
}