package common

type Operator string

func (o Operator) String() string {
	return string(o)
}

const (
	// Comparison Operators
	Eq  Operator = "$eq"  // 等于
	Ne  Operator = "$ne"  // 不等于
	Gt  Operator = "$gt"  // 大于
	Gte Operator = "$gte" // 大于等于
	Lt  Operator = "$lt"  // 小于
	Lte Operator = "$lte" // 小于等于

	// Logical Operators
	And Operator = "$and" // 与
	Or  Operator = "$or"  // 或
	Not Operator = "$not" // 非
	Nor Operator = "$nor" // 或非

	// Element Operators
	Exists Operator = "$exists" // 字段存在
	Type   Operator = "$type"   // 字段类型

	// Evaluation Operators
	Expr       Operator = "$expr"       // 使用表达式
	JsonSchema Operator = "$jsonSchema" // JSON Schema 验证
	Mod        Operator = "$mod"        // 模运算
	Regex      Operator = "$regex"      // 正则表达式
	Text       Operator = "$text"       // 文本搜索
	Where      Operator = "$where"      // JavaScript 表达式

	// Array Operators
	All       Operator = "$all"       // 数组中所有元素匹配
	ElemMatch Operator = "$elemMatch" // 数组中至少一个元素匹配
	Size      Operator = "$size"      // 数组长度

	// Bitwise Operators
	BitsAllClear Operator = "$bitsAllClear" // 按位全部清除
	BitsAllSet   Operator = "$bitsAllSet"   // 按位全部设置
	BitsAnyClear Operator = "$bitsAnyClear" // 按位任意清除
	BitsAnySet   Operator = "$bitsAnySet"   // 按位任意设置

	// Geospatial Operators
	GeoIntersects Operator = "$geoIntersects" // 地理空间相交
	GeoWithin     Operator = "$geoWithin"     // 地理空间范围内
	Near          Operator = "$near"          // 附近
	NearSphere    Operator = "$nearSphere"    // 球体附近

	// Update Operators
	CurrentDate Operator = "$currentDate" // 当前日期
	Inc         Operator = "$inc"         // 递增
	Mul         Operator = "$mul"         // 乘法
	Rename      Operator = "$rename"      // 重命名字段
	Set         Operator = "$set"         // 设置字段
	SetOnInsert Operator = "$setOnInsert" // 插入时设置字段
	Unset       Operator = "$unset"       // 删除字段
	Min         Operator = "$min"         // 设置字段最小值
	Max         Operator = "$max"         // 设置字段最大值
	AddToSet    Operator = "$addToSet"    // 添加到集合
	Pop         Operator = "$pop"         // 删除数组中的第一个或最后一个元素
	Pull        Operator = "$pull"        // 删除数组中所有匹配的值
	Push        Operator = "$push"        // 添加到数组
	PullAll     Operator = "$pullAll"     // 删除数组中所有的指定值
	Each        Operator = "$each"        // 用于 $push 或 $addToSet，每个操作
	Slice       Operator = "$slice"       // 用于 $push 修剪数组
	Sort        Operator = "$sort"        // 用于 $push 排序数组
	Position    Operator = "$position"    // 用于 $push 指定位置
)
