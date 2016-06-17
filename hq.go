/*
hq - html query tool, that works like jQuery in browser
wrap the library as a binary, for convenience use in shell.

@Author: jackxoing@tencent.com
@Date: 2016-6-18
*/
package main
import (
    "fmt"
    "log"
    "os"
    "flag"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

var (
    url = flag.String("u", "-", "URI or FilePath to scrape. default STDIN, so we can pipe sth :). URL must start with 'http'")
    html = flag.Bool("html", false, "print the innerHTML of the node")
    ohtml = flag.Bool("ohtml", false, "print the outerHTML of the node")
    text = flag.Bool("text", false, "print the TEXT part of the node. same as <-attr 'text'>")
    attr = flag.String("attr", "", "print the attribute <string> in node, <string> are comma seperated, and output is joined with tab. eg: -attr href,target")

    debug = flag.Bool("d", false, "debug or not, if debug, some more will be output")
    //strip = flag.Bool("s", false, "strip white space at the beginning and tail, or not") // always TrimSpace


    selector string = ""
    fun string = ""
)

const (
    USAGE_MSG = `
Example usage: %s [options] <-html|-ohtml|-text|-attr <name1,name2,...> > <selector>
    selector: jQuery style selector. eg: "head script"
    -html|-ohtml|-text|-attr: must specify at least one of these functions

    When u want to print multiple field that combined with text part and attribute, such as href and textbody,  you can <-attr 'href, text'>.
`
)

func Must(v interface{}, v2 interface{}) interface{} {
    return v
}

func run() {
    var (
        doc *goquery.Document
        doc_err, err error
        tmp_url string
        fd *os.File
    )
    tmp_url = strings.ToLower(*url)
    if strings.HasPrefix(tmp_url,"http") {
        doc, doc_err = goquery.NewDocument(*url)
    } else if *url=="-" {
        doc, doc_err = goquery.NewDocumentFromReader(os.Stdin)
    } else {
        if fd,err = os.Open(*url); err==nil {
            doc, doc_err = goquery.NewDocumentFromReader(fd)
        } else {
            log.Fatal("open local file failed:", *url, err)
        }
        defer fd.Close()
    }
    if doc_err != nil {
        log.Fatal("goquery NewDocument err:",err)
    }

    if *debug {
        fmt.Printf(">>>tag=[%s]\n",selector)
    }

    doc.Find(selector).Map(func(i int, sel *goquery.Selection) string {
        output := ""
        switch fun {
        case "html":
            if output, err = sel.Html(); err != nil {
                log.Fatal("select err:",err)
            }
            output = strings.TrimSpace(output)
        case "ohtml":
            if output, err = goquery.OuterHtml(sel); err != nil {
                log.Fatal("select err:",err)
            }
            output = strings.TrimSpace(output)
        case "text":
            output = strings.TrimSpace(sel.Text())
        case "attr":
            attr_list := strings.Split(*attr, ",")
            for _, attr_i := range(attr_list) {
                var output_i string
                if attr_i == "text" { // a hardcode case for convenience scrapy
                    output_i = strings.TrimSpace(sel.Text())
                } else {
                    output_i = sel.AttrOr(strings.TrimSpace(attr_i),"-")
                }
                if output_i == "" { output_i = "-" }
                output += output_i + "\t"
            }
        }

        fmt.Println(output)
        return ""
    })
}

func init() {
    fc := flag.Usage
    flag.Usage = func() {
        fc()
        fmt.Fprintf(os.Stderr, USAGE_MSG, os.Args[0])
    }

    flag.Parse()

    if flag.NArg() > 0 {
        selector = flag.Arg(0)
    }else {
        flag.Usage()
        os.Exit(-1)
    }

    switch {
        case *html :
            fun = "html"
        case *ohtml :
            fun = "ohtml"
        case *text :
            fun = "text"
        case *attr != "":
            fun = "attr"
        default:
            fun = ""
    }
    if fun == "" {
        fmt.Fprintln(os.Stderr, "!! no function specified.")
        flag.Usage()
        os.Exit(-1)
    }
}

func main() {
    run()
}
