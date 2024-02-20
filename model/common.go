package model

// CommonDict 公共字典
type CommonDict struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type TreeNode struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Pid      string     `json:"pid"`
	Disabled bool       `json:"disabled"`
	Children []TreeNode `json:"children"`
	Sort     uint32     `json:"sort"`
}

func NewTreeNode(id string, name string, pid string, disabled bool, sort uint32) TreeNode {
	return TreeNode{
		Id:       id,
		Name:     name,
		Pid:      pid,
		Disabled: disabled,
		Sort:     sort,
	}
}

func ListToTree(list []TreeNode) []TreeNode {
	treeList := []TreeNode{}
	for i := range list {
		if list[i].Pid == "0" {
			treeList = append(treeList, finChildren(list[i], list))
		}
	}
	return treeList
}

func finChildren(t TreeNode, list []TreeNode) TreeNode {
	for i := range list {
		if list[i].Pid == t.Id {
			if t.Children == nil || len(t.Children) == 0 {
				t.Children = []TreeNode{}
			}
			t.Children = append(t.Children, finChildren(list[i], list))
		}
	}
	return t
}

// 分页入参
type PageDTO struct {
	PageNum  uint32 `json:"pageNum"`
	PageSize uint32 `json:"pageSize"`
}

// 返回结果
type PageInfoVO struct {
	PageNum  uint32 `json:"pageNum"`
	PageSize uint32 `json:"pageSize"`
	Pages    uint32 `json:"pages"`
	Total    uint64 `json:"total"`
}

// 数据build条件，下载和查询数据使用
type DataBuildFilter struct {
	IsPage    bool `json:"isPage"`    // 是否分页
	AllColumn bool `json:"allColumn"` // 是否全部字段
}
