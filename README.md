# hq - html query tool, with jQuery notation on dom operation

### About
This is a tool that works like jQuery notation in browser.
一个按jQuery风格的，工作在命令行下的 html dom 筛选器工具

### Usage
```
Usage of ./hq:
  -attr string
        print the attribute <string> in node, <string> are comma seperated, and output is joined with tab. eg: -attr href,target
  -d    debug or not, if debug, some more will be output
  -html
        print the innerHTML of the node
  -ohtml
        print the outerHTML of the node
  -text
        print the TEXT part of the node. same as <-attr 'text'>
  -u string
        URI or FilePath to scrape. default STDIN, so we can pipe sth :). URL must start with 'http' (default "-")

Example usage: ./hq [options] <-html|-ohtml|-text|-attr <name1,name2,...> > <selector>
    selector: jQuery style selector. eg: "head script"
    -html|-ohtml|-text|-attr: must specify at least one of these functions

    When u want to print multiple field that combined with text part and attribute, such as href and textbody,  you can <-attr 'href, text'>.
```

### Example

```bash
./hq -u 'http://news.qq.com' -attr 'href,text' '.Q-tpList div.text a.linkto'  |iconv -f gbk -t utf8
```
Take care about the encoding, u should ICONV it to fit your own env, for the html content is encoded in gbk. : ).

It will produce a list like this below . the original html was stored at index.html (on 18 Jun 2016), just try it

        http://news.qq.com/a/20160617/057214.htm        发改委回应女学生裸条借贷：惩戒失信但要依法
        http://sports.qq.com/nba/       NBA-詹皇41+8+11库里遭逐 骑士大胜总分3-3
        http://news.qq.com/a/20160617/054420.htm        中国军队应越方请求 将协助搜救越南失事飞机
        http://news.qq.com/a/20160617/053030.htm        美国不愿介入中国军舰进日“领海”事件 称宜由中日解决
        http://news.qq.com/a/20160617/053311.htm        黑龙江一厅官妻在家门口收百万 钱太多需送上楼
        http://news.qq.com/a/20160617/056645.htm        押钞员辞职干起抢劫 还爱上被抢的女大学生
        http://news.qq.com/a/20160617/053759.htm        兰博基尼高速上自燃成空壳 车子刚买没几天
        http://news.qq.com/a/20160617/053613.htm        田纪云：建国之后急于消灭个体私营经济是一个重大失误
        http://news.qq.com/a/20160617/049088.htm        安徽高考眼镜被收走女生：希望得到公开道歉
        http://news.qq.com/a/20160617/016183.htm        油漆工冒充国企领导骗80余万 同时与4名女子网恋
        http://news.qq.com/a/20160617/041863.htm        消防员火场救出40居民 把面罩让给婴儿自己被熏黑
        http://news.qq.com/a/20160617/034519.htm        国防部回应海军舰艇航经日本邻近海域：符合国际法
        http://news.qq.com/a/20160617/060541.htm        治国理政新实践山东篇 品牌“打擂” 创新先行
        http://news.qq.com/a/20160617/060396.htm        共谋发展 共享繁荣 中塞友谊再谱新篇章
        http://news.qq.com/a/20160617/028600.htm        中国孕妇赴美生子 所有手续齐全仍被原机遣返
        http://news.qq.com/a/20160617/046431.htm        妻子出轨后坦白 丈夫不信被戴“绿帽”反而更爱她
        http://news.qq.com/a/20160617/051092.htm        黑龙江省人口与计划生育条例修改 再婚也可享婚假
        http://news.qq.com/a/20160617/032488.htm        沪昆高铁全线轨通 设计时速350公里上海到昆明仅8小时
        http://news.qq.com/a/20160617/037466.htm        赞比亚小报就“中国向非洲卖人肉”事件道歉
        http://news.qq.com/a/20160616/059756.htm        创纪录贪近2.5亿 白恩培有啥捞钱法宝？
        http://news.qq.com/a/20160617/004460.htm        京沪等多地调整社保缴费基数 你的社保缴费增加了吗？
        http://news.qq.com/a/20160617/007868.htm        云南大学生疑因唱歌遭室友杀害 校方无明确表态
        http://news.qq.com/a/20160616/060254.htm        陕西一派出所副所长唱K后带走卖酒女 经理接人被打断肋骨
        http://news.qq.com/a/20160616/064177.htm        广西警方回应“近百村民凌晨被抓 ”：因抓嫌犯时遭阻挠
        http://news.qq.com/a/20160616/060058.htm        解放军1月内43人晋升少将 含首位空军女师长
        http://news.qq.com/a/20160617/008370.htm        上海迪士尼开园首日频现游客不文明行为：插队、乱丢垃圾
        http://news.qq.com/a/20160617/009069.htm        英国女议员脱欧公投前一周遭枪击身亡 一名50多岁嫌犯被抓
        http://news.qq.com/a/20160616/058718.htm        央视专访美国枪击案枪手父亲：儿子不是恐怖分子
        http://news.qq.com/a/20160617/007836.htm        广东一碰瓷团伙打断同伙骨头诈骗 断骨者仅分几百提成
        http://news.qq.com/a/20160617/008009.htm        澳洲为防房价涨太快对海外购房者提税 被指针对中国买家
        http://news.qq.com/a/20160616/052995.htm        黑龙江籍3团伙47人在京碰瓷被抓：开豪车撞酒驾司机
        http://news.qq.com/a/20160617/008695.htm        媒体：坊间传闻称李云峰与杨卫泽是“连襟”关系
        http://news.qq.com/a/20160617/008276.htm        媒体：半年5市高官落马 河南反腐也是“蛮拼的”
        http://news.qq.com/a/20160617/004564.htm        北京昌平一铲车冲向人行道 一位老人被撞身亡(图)
        http://news.qq.com/a/20160617/008302.htm        媒体揭秘纪检干部受过啥威胁：有人写信“不回去没好下场”
        http://news.qq.com/a/20160617/007789.htm        陕西咸阳警方：举报社会知名人士吸毒最高奖5千
        http://news.qq.com/a/20160617/005612.htm        媒体：这件事，德国日本都能做，中国凭啥做不好？
        http://news.qq.com/a/20160616/058979.htm        泰国厕所天花板掉下2只巨蜥 吓坏如厕女子(图)
        http://news.qq.com/a/20160617/004420.htm        网传山东中考辅导老师建群作弊 19岁嫌疑人被刑拘
        http://news.qq.com/a/20160617/006669.htm        少年从上海乘飞机偷渡迪拜续：浦东机场新增航线被叫停
        http://news.qq.com/a/20160616/060425.htm        日本男子失踪近20年被发现 曾疑遭朝鲜特工绑架
