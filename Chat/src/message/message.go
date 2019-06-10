package message

//定义消息结构体 使用反引号包含的文本是Go对象和JSON之间进行序列化和反序列化时需要的元数据
type Message struct {
	Email string `json:"email"`
	UserName string `json:"username"`
	Message string `json:"message"`
}
