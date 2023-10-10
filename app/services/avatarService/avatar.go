package avatarService


import (
    "crypto/md5"
    "fmt"
    "strings"
)
// Hash值获取方式 先删除前导和尾随的空格，再将所有字符置为小写，最后进行MD5加密
func Hash(email string) [16]byte {
    return md5.Sum([]byte(strings.ToLower(strings.TrimFunc(email, func(r rune) bool {
        return r == ' '
    }))))
}
// 头像获取地址为 https://cravatar.cn/avatar/{HASH}
func EmailToCravatarURL(email string) string {
    return fmt.Sprintf("https://cravatar.cn/avatar/%x", Hash(email))
}
