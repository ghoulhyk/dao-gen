package confBean

type OrmInfo struct {
	Type                string
	SessionGetterFunc   string
	SessionGetterImport string
}

func (receiver OrmInfo) IsXorm() bool {
	return receiver.Type == "xorm"
}
