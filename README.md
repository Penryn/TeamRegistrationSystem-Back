# 团队报名系统

[前端项目地址](https://github.com/HerveyB3B4/TeamRegistrationSystem-Front)

[后端项目地址](https://github.com/Penryn/TeamRegistrationSystem-Back)

[项目网址](http://47.115.209.120:5173)

## 团队报名系统

#### 基础要求

1.登录与注册功能

2.能够编辑自己的个人信息(包括基本信息以及联系方式)

3.可以创建团队(创建者为队长，能够管理团队和队员)

4.未组队的人能通过团队编号和密码加入团队，团队人数4-6人才可提交报名，且提交报名后不可更换队可以撤销提交状态

#### 提高要求

1.鉴权(`session`或`JWT`)

2.消息通知功能（成员加入、团队解散等）

3.能够上传团队头像

4.考虑管理员页面、比如查看团队提交状态、用户管理等

>  *注意事项*
>
> 1.优先保证基础要求有余力的团队可以考虑提高要求。
>
> 2.可以在原有要求之上进行自由拓展考虑和展现的东西更全面的话会有额外加分
>
> 3.使用`Apifox`作为接口文档。
>
> 4.使用`git`进行版本管理，并创建一个`GitHub`的库托管代码。
>
> 5.在开发完成后尝试将项目部署到服务器上（最低要求是本地前后端能一起展现）
> 
#### 后端完成进度

1.用户界面（登录，注册，发送邮箱验证码，找回密码，更新个人信息，上传个人头像，获取个人信息，获取个人消息）

2.团队界面（创建团队，解散团队，加入团队，退出团队，踢出团队，编辑团队信息，上传团队头像，获取团队信息，搜索团队，报名，取消报名）

3.管理员页面（获取全部用户名单，获取报名（未报名）团队名单，删除用户，向全体成员发送消息）
