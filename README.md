# hq - html query tool, with jQuery notation on dom operation

### About
This is a tool that works like jQuery notation in browser.
一个按jQuery风格的，工作在命令行下的 html dom 筛选器工具

### Changelog
* 20160626: auto detect the file encoding by `<meta http-equiv=...>`  or `<meta charset=...>` tag, then re-encode the output to utf8, if options [-noenc] is not specified.
* 20160626: open the -u STRING first as file, if failed , as url. and if not prefix with http, add http:// before request

### Usage
```
Usage of ./hq:
  -attr string
        print the attribute <string> in node, <string> are comma seperated, and output is joined with tab. eg: -attr href,target
  -d    debug or not, if debug, some more will be output
  -html
        print the innerHTML of the node
  -noenc
        DO NOT care about the output encoding. without this option, we try to detect and encode the output to utf8
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
./hq -u 'http://www.qq.com' -attr 'href,text' 'div#newsInfoQuanguo a[target="_blank"]'  
```
** Note **
* The selector chose the 'XXX' in  `<div id="newsInfoQuanguo">...<a target="_blank">xx</a>...</div>`. Learn more about the jQuery selector (here http://www.w3school.com.cn/jquery/jquery_selectors.asp)
* Take care about the encoding, u should make sure that it fits your own env. we consider that you works in utf8 envirionment

It will produce a list like this below . the original html was stored at index.html (on 26 Jun 2016), just try it

        http://news.qq.com/a/20160625/030480.htm        习近平同普京会谈
        http://news.qq.com/a/20160626/007469.htm        中俄联合声明:加强全球战略稳定
        http://news.qq.com/a/20160626/011341.htm        推进信息网络空间发展
        http://news.qq.com/p/topic/20140818025149/index.html    专题
        http://news.qq.com/zt2016/zhesannian/index.htm  治国理政
        http://news.qq.com/p/topic/20160620059297/index.html    建党95周年
        http://news.qq.com/a/20160625/029599.htm        李克强历届达沃斯演讲传递什么信号
        http://news.qq.com/a/20160625/030573.htm        会见普京
        http://finance.qq.com/zt2016/davos2016s/index.htm       专题
        http://news.qq.com/a/20160626/014784.htm        国资委副主任等3位省部级官员被通报违反八项规定
        http://news.qq.com/a/20160626/021171.htm        公安部副部长率队赴湖南宜章调查客车起火事故
        http://news.qq.com/a/20160626/001566.htm        长征七号火箭首飞成功
        http://news.qq.com/a/20160626/001201.htm        揭秘：为何尾焰多了蓝色
        http://news.qq.com/a/20160626/008391.htm#p=1    后悔了？英国民众吁二次公投 请愿签名超157万
        http://news.qq.com/a/20160626/014658.htm        老太被蜱虫叮咬生命垂危 320小时血液过滤救回
        http://news.qq.com/a/20160626/008072.htm#p=1    -
        http://news.qq.com/a/20160626/008072.htm#p=1    张家界玻璃桥“安检”：小车碾压 30人铁锤猛砸
        http://news.qq.com/a/20160626/006650.htm        江苏盐城救援战士昼夜奋斗 累了睡废墟
        http://news.qq.com/a/20160626/015564.htm        中纪委机关报：党员不交党费就是脱党 就应除名
        http://news.qq.com/a/20160626/001233.htm        长征七号到登月有多远
        http://news.qq.com/newspedia/changqi.htm        动画演示火箭有多牛
        http://news.qq.com/a/20160626/013499.htm        “百名红通人员”30人到案 都是怎么追回来的
        http://news.qq.com/a/20160626/014704.htm#p=1    朝鲜举行大规模群众集会 纪念“反美斗争日”
        http://news.qq.com/a/20160626/013335.htm        揭KTV冰妹：陪人吸毒1次赚上千 心肝等器官患病
        http://sports.qq.com/isocce/2016copaamerica/    美洲杯-哥伦比亚1-0美国获季军 巴卡破门制胜
        http://sports.qq.com/nba        NBA-杜兰特七大下家曝光 美男篮12人名单出炉
        http://news.qq.com/a/20160626/016906.htm#p=1    -
        http://news.qq.com/a/20160626/016906.htm#p=1    夫妻因租不起房 带着婴儿在网吧吃住半个月
        http://news.qq.com/a/20160626/006940.htm#p=1    陕西神木一镇政府新房烂尾 借房办公
        http://news.qq.com/a/20160626/000214.htm        公务员谈仕途变化：51岁还能晋升 不再跑官要官
        http://news.qq.com/a/20160626/002534.htm        90岁抗战老兵公园内乞讨 官方：不符优抚政策
        http://news.qq.com/a/20160626/011184.htm#p=1    龙卷风到来时，他在废墟下“舍身护妻”
        http://news.qq.com/a/20160626/005212.htm        高考状元清华博士毕业回西安：北京房价太贵
        http://news.qq.com/a/20160626/019287.htm        蓝孔雀男厕洗手台照镜子 网友：美男子
        http://ent.qq.com/a/20160625/031437.htm 黄渤演讲拿自己开涮：教你如何抹平“颜值差”
        http://ent.qq.com/a/20160626/010372.htm 小贝赴港出席活动 徐子淇半亿黄钻抢镜
        http://ent.qq.com/a/20160625/031942.htm 【存照】美丽俏女神，公公眼中好儿媳  
