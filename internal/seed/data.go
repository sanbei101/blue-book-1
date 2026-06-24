package seed

import "strconv"

// 用户数据池
var userSeeds = []struct {
	Username string
	Bio      string
}{
	{"xiangqian_dev", "前端开发 | React/Vue 双修 | 偶尔写写 Go"},
	{"go_fanatic", "Go 语言重度爱好者 | 开源贡献者 | 偶尔写 Rust"},
	{"product_li", "产品经理 | 需求挖掘机 | 画 Figma 的人"},
	{"designer_wang", "UI/UX 设计师 | Figma 重度用户 | 配色强迫症"},
	{"backend_zhao", "后端工程师 | 分布式系统 | 微服务拆分爱好者"},
	{"student_chen", "大三计算机在读 | 刷题选手 | 求个实习 💼"},
	{"devops_sun", "DevOps 工程师 | K8s 玩家 | CI/CD 流水线搭建"},
	{"mobile_zhou", "移动端开发 | Flutter/iOS/Android 全栈 | 性能优化"},
	{"data_analyst", "数据分析师 | Python/Pandas | 用数据讲故事"},
	{"ai_researcher", "AI/NLP 研究员 | LLM 微调 | 偶尔跑跑 diffusion"},
	{"fullstack_liu", "全栈工程师 | Next.js + Go | 独立开发者"},
	{"security_wu", "安全工程师 | 渗透测试 | CTF 选手 🏴‍☠️"},
	{"test_master", "测试工程师 | 自动化测试 | 质量守护者"},
	{"linux_hacker", "Linux 爱好者 | Arch 用户 btw | 终端美学"},
	{"newbie_coder", "转码新人 | 正在学 Go | 每天进步一点点 🌱"},
	{"tech_lead", "技术负责人 | 架构设计 | 代码审查狂魔"},
	{"open_source", "开源爱好者 | GitHub 重度用户 | Star 收集器 ⭐"},
	{"cloud_native", "云原生工程师 | Serverless | 边缘计算"},
	{"game_dev", "游戏开发 | Unity/Godot | 独立游戏制作者"},
	{"startup_founder", "创业者 | 全栈开发 | 产品+技术双修"},
}

// 帖子数据池
var postSeeds = []struct {
	Title   string
	Content string
}{
	{
		"Go 1.22 泛型性能实测:比 interface{} 快多少?",
		"最近在项目中把一些热点路径从 interface{} 迁移到了泛型,做了个简单的 benchmark 对比。\n\n结论:在高频调用场景下,泛型版本快了约 30-40%,主要是省去了类型断言和反射的开销。对于低频调用场景差异不大。\n\n分享一下测试代码和结果,欢迎大家讨论 👇",
	},
	{
		"面试翻车记录:字节跳动后端一面",
		"今天面了字节的后端开发岗,记录一下翻车现场:\n\n1. 手撕 LRU Cache,写到一半忘了双向链表的删除操作\n2. 问 MySQL 索引原理,B+ 树说了个大概但细节没答好\n3. 系统设计题:设计一个短链接服务,只说了基本思路没深入\n\n总结:基础还是不扎实,继续加油吧 😭",
	},
	{
		"为什么我从 React 转向了 Vue",
		"用了三年 React,最近新项目尝试了 Vue 3 + Composition API,说说感受:\n\n优点:\n- 模板语法对后端同学更友好\n- 响应式系统用起来确实更直觉\n- 生态整合度高(Router/Pinia/DevTools)\n\n缺点:\n- TypeScript 支持虽然进步了但还是不如 React\n- 复杂组件的逻辑组织不如 hooks 灵活\n\n总的来说,两个框架都很优秀,选择取决于团队和项目。",
	},
	{
		"分享一下我的终端配置(附 dotfiles)",
		"折腾了一周终于把终端环境配好了,分享一下配置:\n\n- Shell: zsh + oh-my-zsh\n- 终端: Wezterm(跨平台、GPU 加速)\n- 主题: Catppuccin Mocha\n- 字体: JetBrains Mono Nerd Font\n- 通知: starship prompt\n- 模糊搜索: fzf + ripgrep\n\ndotfiles 已经开源了,链接在评论区 👇",
	},
	{
		"Docker Compose 搭建本地开发环境指南",
		"团队新人来了之后环境搭建总是出问题,索性写了个 docker-compose.yml 一键拉起所有依赖:\n\n- PostgreSQL 16\n- Redis 7\n- MinIO(S3 兼容存储)\n- Mailhog(邮件测试)\n\nCompose 文件和初始化脚本都放在仓库里了,clone 下来 docker compose up -d 就能跑。",
	},
	{
		"产品经理又改需求了,聊聊如何优雅地应对",
		"周一开需求评审会,PM 突然说'这个功能我们加个小改动',然后从兜里掏出了一个 20 页的 PRD...\n\n分享一下我的应对策略:\n1. 先评估影响范围,列出改动点\n2. 和 PM 确认优先级,哪些必须这期做\n3. 技术方案先写 RFC,评审通过再动手\n4. 留 buffer,需求变更几乎是必然的\n\n大家有什么好的经验分享吗?",
	},
	{
		"用 Go 实现一个简单的 HTTP 中间件链",
		"最近在学习 Go 的 HTTP 中间件模式,写了一个简单的实现:\n\n```go\nfunc Chain(h http.Handler, middlewares ...Middleware) http.Handler {\n    for i := len(middlewares) - 1; i >= 0; i-- {\n        h = middlewares[i](h)\n    }\n    return h\n}\n```\n\n支持日志、认证、限流等中间件的组合,代码很简洁。详细实现和用法在文末链接。",
	},
	{
		"远程办公一年的真实感受",
		"从去年开始全职远程办公,分享一下真实体验:\n\n优点:\n- 省去通勤时间(每天多出 2 小时)\n- 工作时间灵活,效率反而更高\n- 可以回老家办公,生活成本降低\n\n缺点:\n- 沟通成本增加,很多事要异步处理\n- 容易工作生活界限模糊\n- 偶尔会感到孤独\n\n建议想远程的同学先试试混合办公模式。",
	},
	{
		"PostgreSQL 查询优化实战:从 3s 到 50ms",
		"线上有个接口响应很慢,排查发现是 SQL 查询的问题:\n\n原始查询:3.2s\n- 加了合适的索引:降到 800ms\n- 改写了子查询为 JOIN:降到 200ms\n- 加了覆盖索引:降到 50ms\n\n关键点:\n1. EXPLAIN ANALYZE 是你的好朋友\n2. 注意索引的选择性和区分度\n3. 避免 SELECT *,只查需要的字段\n4. 合理使用连接池",
	},
	{
		"大三学生的秋招准备清单",
		"距离秋招还有 3 个月,整理了一下需要准备的内容:\n\n算法:\n- LeetCode Hot 100 至少刷两遍\n- 重点:二叉树、动态规划、图论\n\n基础:\n- 操作系统:进程/线程、内存管理、文件系统\n- 计算机网络:TCP/UDP、HTTP/HTTPS、DNS\n- 数据库:索引、事务、MVCC\n\n项目:\n- 准备 2-3 个项目,能讲清楚难点和优化\n\n大家一起加油 💪",
	},
	{
		"聊聊微服务拆分的血泪教训",
		"去年把一个单体应用拆成了微服务,踩了不少坑:\n\n1. 拆分粒度太细,服务间调用链太长\n2. 分布式事务没处理好,数据一致性问题频发\n3. 服务治理没跟上,出了问题排查困难\n4. 部署复杂度暴增,CI/CD 经常挂\n\n总结:不是所有系统都需要微服务,拆分前想清楚。",
	},
	{
		"CSS Container Queries 实战体验",
		"终于在生产项目中用上了 CSS Container Queries,分享一下体验:\n\n以前响应式设计只能根据视口宽度来,现在可以根据父容器宽度来调整样式,这对组件化开发太友好了。\n\n比如一个卡片组件,在侧边栏里显示精简版,在主内容区显示完整版,不需要写 JS 逻辑。\n\n浏览器支持已经很好了,推荐大家试试。",
	},
	{
		"用 AI 辅助编程一个月的真实体验",
		"用 Copilot + Claude 写了一个月代码,说说感受:\n\n适合的场景:\n- 写样板代码、CRUD、测试用例\n- 解释不熟悉的代码\n- 生成正则表达式\n\n不太适合的场景:\n- 复杂的业务逻辑设计\n- 性能优化方案\n- 架构决策\n\n总体来说效率提升了大约 30%,但核心思考还是得靠自己。",
	},
	{
		"Redis 实现分布式锁的正确姿势",
		"面试常考题:如何用 Redis 实现分布式锁?\n\n错误示范:SETNX + EXPIRE(非原子操作)\n\n正确姿势:\n```\nSET lock_key random_value NX EX 30\n```\n\n释放锁时要用 Lua 脚本保证原子性:\n```lua\nif redis.call('get', KEYS[1]) == ARGV[1] then\n    return redis.call('del', KEYS[1])\nelse\n    return 0\nend\n```\n\n更复杂的场景建议用 Redlock 算法或 etcd。",
	},
	{
		"独立开发者的第一款产品上线了",
		"折腾了三个月,我的第一款独立产品终于上线了 🎉\n\n产品:一个 Markdown 笔记应用,支持实时协作和 AI 总结\n\n技术栈:\n- 前端:Next.js + Tiptap\n- 后端:Go + PostgreSQL\n- 部署:Vercel + Supabase\n\n第一个月数据:\n- 注册用户:500+\n- 日活:50+\n- 付费用户:12\n\n虽然不多但很开心,有用户真的在用!",
	},
	{
		"Git 工作流最佳实践:我们团队的选择",
		"对比了几种 Git 工作流,最后我们团队选择了 Trunk-Based Development:\n\n核心原则:\n- 主干分支保持随时可发布\n- 功能分支生命周期不超过 2 天\n- 用 Feature Flag 控制未完成功能\n- CI/CD 必须可靠,合并前必须通过所有检查\n\n配合 PR 模板和代码审查,效果很不错。",
	},
	{
		"从零搭建 Kubernetes 集群的踩坑记录",
		"帮公司从零搭建了一套 K8s 集群,记录一下踩过的坑:\n\n1. 网络插件选型:Calico vs Flannel,最后选了 Calico\n2. 存储方案:一开始用 hostPath,后来换成了 Longhorn\n3. 监控告警:Prometheus + Grafana 是标配\n4. 日志收集:EFK 方案,Fluent Bit 比 Fluentd 资源占用小\n5. Ingress Controller:Nginx Ingress 配置相对简单\n\n建议小团队直接用托管服务,自建维护成本太高。",
	},
	{
		"TypeScript 5.4 新特性速览",
		"TypeScript 5.4 正式发布了,几个值得关注的特性:\n\n1. NoInfer 工具类型:防止类型推断的意外行为\n2. 闭包中的类型收窄改进\n3. Object.groupBy 和 Map.groupBy 支持\n4. --moduleResolution bundler 改进\n\n个人最期待的是 NoInfer,之前在写泛型库的时候经常遇到类型推断太宽的问题。",
	},
	{
		"程序员的职业发展路径思考",
		"工作五年了,最近在思考职业发展方向:\n\n技术路线:\n- 深耕某个领域成为专家\n- 持续学习新技术保持竞争力\n- 技术影响力(开源、演讲、写作)\n\n管理路线:\n- 技术管理:带团队做项目\n- 需要培养沟通、协调、向上管理能力\n\n我的选择:先走技术路线,积累足够的深度再考虑转型。大家怎么看?",
	},
	{
		"写给新人的 Go 语言学习路线",
		"经常有同学问我 Go 怎么学,整理一个学习路线:\n\n入门(1-2 周):\n- Go Tour 官方教程\n- 《Go 程序设计语言》前 5 章\n\n进阶(1-2 月):\n- Go 标准库源码阅读\n- 并发编程(goroutine/channel)\n- 接口和组合的设计模式\n\n实战:\n- 写一个 CLI 工具\n- 实现一个 HTTP 服务器\n- 参与开源项目\n\n资源链接在评论区。",
	},
	{
		"系统设计面试:设计一个短链接服务",
		"高频面试题:设计一个短链接服务,分享一下我的思路:\n\n核心功能:\n- 长链接转短链接\n- 短链接重定向到原链接\n\n技术方案:\n- 生成算法:雪花算法 / 自增ID + Base62 编码\n- 存储:MySQL + Redis 缓存\n- 读写比:读远大于写,重点优化读性能\n\n关键指标:\n- QPS 预估:10 万读 / 1000 写\n- 存储预估:5 年数据约 10 亿条\n- 可用性:99.99%\n\n详细设计图在文末。",
	},
	{
		"聊聊代码审查的那些事",
		"作为 reviewer,分享一下代码审查的经验:\n\n关注点:\n1. 逻辑正确性:边界条件、异常处理\n2. 可读性:命名、注释、代码结构\n3. 性能:是否有明显的性能问题\n4. 安全:SQL 注入、XSS、权限校验\n\n原则:\n- 对事不对人\n- 提供建设性意见\n- 及时 review,不要积压\n- 大改动建议拆分",
	},
	{
		"前端性能优化实战:首屏加载从 5s 降到 1s",
		"优化了一个后台管理系统的首屏加载速度:\n\n优化前:5.2s\n优化措施:\n1. 路由懒加载:-1.5s\n2. 图片压缩 + WebP:-0.8s\n3. 代码分割 + Tree Shaking:-0.6s\n4. CDN 加速静态资源:-0.5s\n5. 接口数据缓存:-0.3s\n\n优化后:1.5s,用户体验提升明显。\n\n工具推荐:Lighthouse、Webpack Bundle Analyzer",
	},
	{
		"etcd 在微服务中的应用",
		"最近在项目中用 etcd 做服务注册发现和配置中心,分享一下经验:\n\n使用场景:\n- 服务注册与发现\n- 分布式配置管理\n- 分布式锁\n- Leader 选举\n\n注意事项:\n- 合理设置 TTL 和续约间隔\n- 监听 Watch 事件处理好重连\n- 集群部署至少 3 个节点\n- 定期做 compaction 避免数据膨胀",
	},
	{
		"程序员副业收入渠道总结",
		"盘点一下程序员常见的副业收入渠道:\n\n技术类:\n- 技术博客/公众号(广告+打赏)\n- 录制技术课程(网易云课堂、极客时间)\n- 接私活(朋友介绍为主)\n- 开发独立产品(SaaS、工具类)\n\n非技术类:\n- 技术咨询/面试辅导\n- 写技术书籍\n- 社群运营\n\n个人经验:副业不要影响主业,先做好本职工作。",
	},
	{
		"用 Go 实现一个简单的任务调度器",
		"分享一个用 Go 实现的简单任务调度器:\n\n支持功能:\n- 定时任务(Cron 表达式)\n- 延迟任务\n- 任务重试\n- 任务状态查询\n\n核心用到了 time.Ticker、channel 和 goroutine pool,代码量不大但涵盖了 Go 并发的很多知识点。\n\n源码和使用示例在 GitHub 仓库。",
	},
	{
		"聊聊数据库事务隔离级别",
		"面试高频题:MySQL 的四种事务隔离级别分别解决了什么问题?\n\n1. READ UNCOMMITTED:什么都没解决,会脏读\n2. READ COMMITTED:解决脏读,但不可重复读\n3. REPEATABLE READ(MySQL 默认):解决不可重复读,基本解决幻读\n4. SERIALIZABLE:解决所有问题,但性能最差\n\nInnoDB 在 RR 级别下通过 MVCC + Next-Key Lock 解决了大部分幻读问题。",
	},
	{
		"设计系统搭建经验分享",
		"给团队搭建了一套设计系统,分享一下经验:\n\n组件库:\n- 基础组件:Button、Input、Modal 等\n- 业务组件:根据业务场景封装\n\n设计规范:\n- 颜色系统:主色、辅助色、语义色\n- 间距系统:4px 基准网格\n- 字体系统:标题、正文、辅助文字\n\n文档:\n- Storybook 做组件展示\n- 使用文档 + 代码示例\n\n维护:\n- 版本管理\n- 变更日志\n- 向后兼容",
	},
	{
		"WASM 入门:用 Go 写前端可行吗?",
		"尝试用 Go + WASM 写了一个简单的前端应用,记录一下体验:\n\n优点:\n- 可以复用 Go 的知识和生态\n- 类型安全\n- 性能比 JS 好(计算密集场景)\n\n缺点:\n- 生成的 WASM 文件太大(最小 2MB+)\n- DOM 操作还是得通过 JS 互操作\n- 生态不成熟,库太少\n\n结论:目前不建议用于生产,但值得关注。",
	},
	{
		"程序员如何保持技术敏感度",
		"分享一下我保持技术敏感度的方法:\n\n信息源:\n- Hacker News(英文世界的技术风向标)\n- Twitter/X(关注技术大佬)\n- GitHub Trending(了解热门项目)\n- 技术博客(高质量的长文)\n\n习惯:\n- 每天花 30 分钟浏览技术资讯\n- 每周写一篇技术笔记\n- 每月尝试一个新技术\n- 每季度做一次技术复盘\n\n核心:保持好奇心,不要只埋头干活。",
	},
	{
		"云原生应用的 12 要素原则",
		"12 要素原则是构建现代 SaaS 应用的方法论:\n\nI. 代码库:一份代码,多份部署\nII. 依赖:显式声明依赖\nIII. 配置:在环境中存储配置\nIV. 后端服务:把后端服务当作附加资源\nV. 构建/发布/运行:严格分离构建和运行\nVI. 进程:以无状态进程运行\nVII. 端口绑定:通过端口绑定提供服务\nVIII. 并发:通过进程模型进行扩展\nIX. 易处理:快速启动和优雅终止\nX. 开发/生产环境等价\nXI. 日志:把日志当作事件流\nXII. 管理进程:后台管理任务作为一次性进程运行\n\n推荐阅读原文。",
	},
	{
		"写了一个 VS Code 插件,解决了我的痛点",
		"每次写 SQL 都要切到数据库客户端太麻烦了,于是写了一个 VS Code 插件:\n\n功能:\n- 直接在编辑器里执行 SQL\n- 连接多个数据库(MySQL/PostgreSQL/SQLite)\n- 结果展示为表格\n- 支持导出 CSV\n\n技术栈:\n- TypeScript\n- VS Code Extension API\n- node-postgres / mysql2\n\n已上架 VS Code 扩展市场,搜索 sql-runner 就能找到。",
	},
	{
		"聊聊 Go 的错误处理哲学",
		"Go 的错误处理一直是争议话题,分享一下我的看法:\n\nGo 的设计哲学:错误是值,不是异常\n\n优势:\n- 显式处理,不会遗漏\n- 控制流清晰\n- 性能好,没有异常栈开销\n\n劣势:\n- 代码冗长(if err != nil)\n- 容易被忽略(_ = err)\n\n最佳实践:\n- 在包边界定义哨兵错误\n- 用 fmt.Errorf 包装错误添加上下文\n- 使用 errors.Is/As 判断错误类型\n- 不要 panic,除非真的不可恢复",
	},
	{
		"前端工程化:Monorepo 实践",
		"把团队的前端项目迁移到了 Monorepo,分享一下实践:\n\n工具选择:pnpm workspace\n\n目录结构:\n- apps/:应用(web、admin、h5)\n- packages/:共享包(ui、utils、config)\n\n优点:\n- 代码复用方便\n- 依赖管理统一\n- 原子提交\n- CI/CD 简化\n\n挑战:\n- 构建工具配置复杂\n- 需要处理好包之间的依赖\n- CI 流水线需要优化",
	},
	{
		"SQL 注入攻击原理与防御",
		"安全基础:SQL 注入是怎么回事?\n\n攻击原理:\n```sql\n-- 原始查询\nSELECT * FROM users WHERE id = '用户输入'\n\n-- 恶意输入\n1' OR '1'='1\n\n-- 最终执行\nSELECT * FROM users WHERE id = '1' OR '1'='1'\n```\n\n防御措施:\n1. 参数化查询(最有效)\n2. ORM 框架\n3. 输入验证\n4. 最小权限原则\n5. WAF\n\n永远不要拼接 SQL!",
	},
	{
		"聊聊技术选型的方法论",
		"技术选型是架构师的核心能力之一,分享一下方法论:\n\n评估维度:\n1. 技术成熟度:社区活跃度、文档质量\n2. 团队能力:学习成本、现有技术栈\n3. 业务需求:性能、可扩展性、稳定性\n4. 生态系统:第三方库、工具支持\n\n决策流程:\n1. 明确需求和约束\n2. 列出候选方案\n3. 做 PoC 验证\n4. 评审讨论\n5. 决策并记录理由\n\n原则:没有最好的技术,只有最合适的技术。",
	},
	{
		"用 Rust 重写了 Node.js 工具,性能提升 10 倍",
		"把一个用 Node.js 写的文件处理工具用 Rust 重写了:\n\n原版(Node.js):\n- 处理 1GB 文件:45 秒\n- 内存占用:800MB\n\nRust 版:\n- 处理 1GB 文件:4.2 秒\n- 内存占用:120MB\n\n主要优化点:\n- 流式处理,避免全量加载\n- 零拷贝解析\n- 并行处理(rayon)\n\nRust 的学习曲线虽然陡,但在性能敏感场景真的很值得。",
	},
	{
		"程序员如何写好技术文档",
		"好的技术文档能减少 80% 的沟通成本,分享一下写作经验:\n\n文档类型:\n- README:项目介绍、快速开始\n- API 文档:接口说明、参数、示例\n- 设计文档:背景、方案、权衡\n- 运维手册:部署、监控、故障处理\n\n写作原则:\n1. 先写大纲再填充\n2. 用代码示例代替纯文字\n3. 图比文字清晰\n4. 保持更新,过时文档比没有更糟\n\n工具推荐:Notion、Confluence、GitHub Wiki",
	},
	{
		"聊聊分布式系统中的 CAP 定理",
		"CAP 定理是分布式系统的基石:\n\nC(Consistency):一致性\nA(Availability):可用性\nP(Partition Tolerance):分区容错性\n\n定理:三者最多只能同时满足两个\n\n实际应用:\n- CP 系统:ZooKeeper、etcd(强一致,牺牲部分可用性)\n- AP 系统:Cassandra、DynamoDB(高可用,最终一致)\n- CA 系统:单机数据库(无网络分区问题)\n\n现实:网络分区不可避免,所以要在 CP 和 AP 之间选择。",
	},
	{
		"游戏开发入门:用 Godot 做一个小游戏",
		"花了一周用 Godot 引擎做了一个小平台跳跃游戏:\n\n技术栈:\n- 引擎:Godot 4.2\n- 语言:GDScript\n- 美术:像素风格(Aseprite 绘制)\n\n学习心得:\n- Godot 的节点系统很直觉\n- GDScript 对 Python 用户很友好\n- 物理引擎开箱即用\n- 社区资源丰富\n\n成品已经上传 itch.io,可以免费玩。",
	},
	{
		"聊聊 API 设计的最佳实践",
		"好的 API 设计能大大减少前后端的沟通成本:\n\n命名规范:\n- 使用名词而非动词\n- 复数形式(/users, /posts)\n- kebab-case(/user-profiles)\n\n版本管理:\n- URL 路径(/api/v1/)\n- 请求头(Accept-Version)\n\n响应格式:\n```json\n{\n  \"code\": 0,\n  \"message\": \"success\",\n  \"data\": {}\n}\n```\n\n错误处理:\n- 使用标准 HTTP 状态码\n- 返回详细的错误信息\n- 提供错误码便于排查",
	},
	{
		"聊聊技术债务的管理",
		"技术债务是每个团队都会面临的问题:\n\n类型:\n- 故意的债务:为了快速上线,后续重构\n- 无意的债务:能力不足导致的烂代码\n- 老化的债务:技术栈过时\n\n管理策略:\n1. 识别和记录:维护一个技术债务清单\n2. 评估影响:业务影响 + 维护成本\n3. 制定计划:每个迭代分配 20% 时间\n4. 持续偿还:不要让债务越积越多\n\n原则:技术债务不可怕,可怕的是不管理。",
	},
	{
		"前端状态管理方案对比",
		"对比了几种主流的前端状态管理方案:\n\nReact 生态:\n- Redux:成熟稳定,但模板代码多\n- Zustand:轻量级,API 简洁\n- Jotai:原子化状态,灵活\n- React Query:服务端状态管理\n\nVue 生态:\n- Pinia:Vuex 的继任者,推荐\n- VueUse:组合式工具集\n\n选择建议:\n- 小项目:Context + useReducer\n- 中型项目:Zustand / Pinia\n- 大型项目:Redux Toolkit / Pinia\n- 服务端状态:React Query / SWR",
	},
	{
		"聊聊 CI/CD 流水线设计",
		"一个好的 CI/CD 流水线能显著提升团队效率:\n\n流水线阶段:\n1. 代码检查:Lint、格式化\n2. 单元测试:覆盖率要求\n3. 构建:编译、打包\n4. 集成测试:E2E 测试\n5. 部署:staging -> production\n\n工具选择:\n- GitHub Actions:开源项目首选\n- GitLab CI:自托管首选\n- Jenkins:老牌,插件丰富\n\n最佳实践:\n- 快速反馈(<10 分钟)\n- 失败即停止\n- 环境一致性\n- 自动化回滚",
	},
}

// 评论数据池
var commentSeeds = []string{
	"学到了!收藏收藏 📚",
	"大佬能分享一下源码吗?",
	"这个方案确实不错,我们项目也遇到了类似的问题",
	"讲得很清楚,比官方文档好懂多了",
	"求出个续集!想看更深入的分析",
	"有遇到过 xx 问题吗?怎么解决的?",
	"感谢分享,正好用得上 👍",
	"这个工具太棒了,已经 star 了",
	"说得很对,深有同感",
	"能详细讲讲第三点吗?没太理解",
	"我们团队也在用这套方案,效果确实好",
	"新手表示看不太懂,有没有入门级的资料?",
	"这个坑我也踩过 😂",
	"效率提升很明显,值得尝试",
	"有性能测试数据吗?想看看具体指标",
	"终于有人把这个讲清楚了!",
	"可以再举个实际案例吗?",
	"我已经在生产环境用了,确实稳定",
	"这个库的文档确实不太友好",
	"面试的时候被问到过,当时没答好",
	"收藏了,回头仔细研究一下",
	"大佬带带我 🙏",
	"代码风格很好,学习了",
	"有没有遇到兼容性问题?",
	"这个思路很巧妙,没想到可以这样",
	"确实,踩过坑才知道这些经验的价值",
	"能分享一下踩坑经历吗?",
	"写得很用心,支持一下 ❤️",
	"我觉得还可以用 xx 方案来做",
	"这个工具我用了半年了,推荐",
}

// 嵌套回复数据池
var replySeeds = []string{
	"感谢回复!已经解决了",
	"对,就是这个问题",
	"确实,我也这么认为",
	"好的,我去试试",
	"你说得对,我理解错了",
	"是的,后续会更新",
	"收到,谢谢指正",
	"哈哈,确实是这样",
	"明白,我去补一下这块知识",
	"谢谢大佬解答 🙏",
}

// 帖子媒体 URL 模板(使用 picsum.photos)
func mediaURL(seed int) string {
	return "https://picsum.photos/seed/" + strconv.Itoa(seed) + "/800/600"
}

// 头像 URL 模板(使用 Dicebear)
func avatarURL(username string) string {
	return "https://api.dicebear.com/7.x/thumbs/svg?seed=" + username
}
