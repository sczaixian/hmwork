

{ 必须和 func 函数声明在同一行上，且位于末尾，不能独占一行，
x+y 中，可在 + 后换行，不能在 + 前换行
不需要在语句或者声明的末尾添加分号


声明一个变量： 
s := ""   // 只能用在函数内部，而不能用于包变量
var s string  // 依赖于字符串的默认初始化零值机制，被初始化为 ""
// 下面量值使用少
var s = ""
var s string = ""  // 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了

import "bufio"  // 处理输入和输出方便又高效 Scanner 类型是该包最有用的特性之一
input := bufio.NewScanner(os.Stdin)  // 接收一个标准输入 
for input.Scan() {                   // 读入下一行，并移除行末的换行符
	counts[input.Text()]++  
}


go doc http.ListenAndServe  // 相当远 man

// 除 数值、bool、string 外其他初始化为 nil

// 返回函数中局部变量的地址也是安全的

// 表名 添加这个方法指定表明，不然名字应该叫 users_models
func (u *UsersModel) TableName() string {
	return "tb_users"
}


// 编译时类型检查的技巧，用于确保SysUser类型实现了Login接口，编译时接口实现检查的技术
var _ Login = new(SysUser)


type BaseModel struct {
	// *gorm.DB 在模型方法中访问数据库操作
	*gorm.DB  `gorm:"-" json:"-"` // gorm:"-" 忽略这个字段，不映射到数据库列，json:"-" json 编解码时忽略这个字段
	Id        int64               `gorm:"primaryKey" json:"id"` // json序列化时使用id
	CreatedAt string              `json:"created_at"`           //日期时间字段统一设置为字符串即可
	UpdatedAt string              `json:"updated_at"`           // 这俩字段没有使用 gorm 约束，会按照默认规则映射到数据库列
	
}

type User struct{
	ID            uint           `gorm:"primarykey" json:"id"`
	Password      string         `json:"-"  gorm:"comment:用户登录密码"`   
	Enable        int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`     
	                                                    // foreignkey:AuthorityId;  AuthorityId为外键， references:AuthorityId 关联到 Authority 中的字段 AuthorityId
	Authority     Authority      `json:"authority" gorm:"foreignkey:AuthorityId;references:AuthorityId;comment:用户角色"`  // 一对多关系
	Authorities   []Authority    `json:"authorities" gorm:"many2many:user_authority;"`   //  多对多关联 通过中间表 UserAuthority 实现
} 

var user User
db.Preload("Authority").First(&user, "username = ?", "admin")
fmt.Println(user.Authority.AuthorityName) // 输出角色名

type UserAuthority struct {   
	UserId               uint `gorm:"column:user_id"`
	AuthorityAuthorityId uint `gorm:"column:authority_authority_id"`
}

type Authority struct {
	//....
	AuthorityId    uint          `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"
	Users          []User        `json:"-" gorm:"many2many:user_authority;"``
}



// 为用户添加附加角色
newRole := Authority{AuthorityId: 999}
db.Model(&user).Association("Authorities").Append(&newRole)

// 查询用户的所有角色
db.Preload("Authorities").Find(&user)
for _, role := range user.Authorities {
    fmt.Println(role.AuthorityName)
}


使用 Preload("Authority").Preload("Authorities")一次性加载所有角色数据，避免 N+1 查询问题














/*----------------------------   init    -----------------------------------------*/
init() 是个特殊的 函数 当包加载是默认调用， 如果有多个文件含有多个 init 根据字典序依次调用




/*----------------------------   sync.Map    -----------------------------------------*/
线程安全：内置并发安全，无需额外同步机制（如互斥锁）。
读写分离：通过 read和 dirty两个内部 map 实现读写分离。readmap 支持无锁读取，dirtymap 用于写入并在需要时加锁。
无锁读取：大多数读操作（命中 readmap）无需加锁，性能极高。
延迟同步：使用懒惰同步（Lazy Synchronization）策略，在读操作频繁时逐步将 dirtymap 的数据同步到 readmap。
原子操作：提供了 LoadOrStore、LoadAndDelete等原子操作方法。


读多写少、多 goroutine 读写不同键、无需复杂操作