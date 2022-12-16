package configure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetArray(t *testing.T) {
	_ = ReadInConfig()
	assert.Equal(t, "", GetString("a.b.c"))
	SetDefault("a.b.c", "123")
	assert.Equal(t, "123", GetString("a.b.c"))

	assert.Equal(t, "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fss_demo?charset=utf8&parseTime=True&loc=Local", GetString("Database.default"))

	arr := GetSlice("A")
	assert.Len(t, arr, 3)
	assert.Equal(t, "a1", arr[0])
}

// 测试环境变量
func TestEnv(t *testing.T) {
	_ = ReadInConfig()
	fmt.Println(GetString("Database.default"))
	os.Setenv("DATABASE_DEFAULT", "aaa")
	assert.Equal(t, "aaa", GetString("Database.default"))
}
