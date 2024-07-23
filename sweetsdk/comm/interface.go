package comm

//	func (m *models.Com) Authentication() string {
//		return m.AppId
//	}
type ComInterface interface {
	// 鉴权
	Authentication() string
	// 获取平台分类
	// GetPclas()
	// 获取小说平台列表
	// GetPlaer() []map[string]interface{}
	// 获取账单
	// GetOrderNow() []map[string]interface{}
	// 获取回填信息
	// 获取关键词列表
	// 获取小说库
	// 关键词提交审核
	// 获取关键词结果
	// 获取结算账单
	// 获取收款地址
}
