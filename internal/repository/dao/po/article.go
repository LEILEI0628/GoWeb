package po

// Article 制作库
type Article struct {
	Id      int64  `gorm:"primaryKey,autoIncrement"`
	Title   string `gorm:"type=varchar(1024)"` // 长度 1024
	Content string `gorm:"type=BLOB"`          // BLOB比较适合大文本（甚至图片也可以）

	// 在 author_id 和 create_time 上创建联合索引
	//AuthorId   int64 `gorm:"index=aid_ctime"`
	//CreateTime int64 `gorm:"index=aid_ctime"`

	// 在 author_id 上创建索引
	AuthorId   int64 `gorm:"index"`
	CreateTime int64 // 创建时间：毫秒数
	UpdateTime int64 // 修改时间：毫秒数

	// 如何设计索引：
	// 帖子的查询场景？
	// 1.对于创作者来说，查看草稿箱所有自己的文章（产品经理要求按照创建时间（或更新时间）的倒序排序）
	// SELECT * FROM articles WHERE author_id = ? ORDER BY `create_time` DESC;
	// 2.单独查询某一篇（主键本来就有索引） SELECT * FROM articles WHERE id = ?
	// 在查询接口深入讨论这个问题（Explain 命令）：
	// - 最佳实践：在 author_id 和 create_time 上创建联合索引
	// - 考虑某一作者文章不会很多：在 author_id 上创建索引

	// 在 articles 表中准备十万/一百万条数据，author_id 各不相同（或者部分相同）
	// 在某一 author_id 插入两百条数据
	// 执行 SELECT * FROM articles WHERE author_id = 123 ORDER BY `create_time` DESC
	// 比较两种索引的性能
}
