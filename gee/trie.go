package gee

import "strings"
// 实现路由模糊匹配的功能
type node struct {
	pattern 	string		// 待匹配路由 完整路径
	part 		string		// 当前这一层的路由中的一部分
	children	[]*node		// 存放所有的孩子节点
	isWild 		bool		// 是否孩子节点是否需要精确匹配
}

// 查找第一个匹配成功的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 将路由插入到树结构中
func (n *node) insert(pattern string, parts []string, height int){
	// 找到该存入的层
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	// 拿到当前层需要匹配的子路由
	part := parts[height]
	// 当前子路由是否在孩子节点中
	child := n.matchChild(part)
	if child == nil{
		child = &node{
			part: part,
			// 可以使用模糊匹配
			isWild: part[0] == ':' || part[0] == '*',
		}
		// 加入到孩子节点中
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// strings.HasPrefix 检测字符串是否以指定的前缀开头
	if len(parts) == height || strings.HasPrefix(n.part, "*"){
		if n.pattern == ""{
			return nil
		}
		return n
	}

	part := parts[height]
	// 查找每个孩子节点
	children := n.matchChildren(part)

	for _, child := range children{
		result := child.search(parts, height+1)
		if result != nil{
			return result
		}
	}
	return nil
}