/*
hq - html query tool, that works like jQuery in browser
wrap the library as a binary, for convenience use in shell.

@Author: jackxoing@tencent.com
@Date: 2016-6-18
*/
package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"log"
	"os"
	"strings"
)

var (
	url   = flag.String("u", "-", "URI or FilePath to scrape. default STDIN, so we can pipe sth :). URL must start with 'http'")
	html  = flag.Bool("html", false, "print the innerHTML of the node")
	ohtml = flag.Bool("ohtml", false, "print the outerHTML of the node")
	text  = flag.Bool("text", false, "print the TEXT part of the node. same as <-attr 'text'>")
	attr  = flag.String("attr", "", "print the attribute <string> in node, <string> are comma seperated, and output is joined with tab. eg: -attr href,target")
	noenc  = flag.Bool("noenc", false, "DO NOT care about the output encoding. without this option, we try to detect and encode the output to utf8")

	debug = flag.Bool("d", false, "debug or not, if debug, some more will be output")
	//strip = flag.Bool("s", false, "strip white space at the beginning and tail, or not") // always TrimSpace

	selector string = ""
	fun      string = ""
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

//eg1: <meta content="text/html; charset=gb2312" http-equiv="Content-Type">
func get_html_enc_instant1(doc *goquery.Document) string {
	content, ok := doc.Find("head meta[http-equiv='Content-Type']").Attr("content")
	if !ok {
		return ""
	}
	fields := strings.Split(content, "=")
	if len(fields) > 1 && strings.HasSuffix(fields[0], "charset") {
		return fields[1]
	} else {
		return ""
	}
}

//eg2: <meta charset="UTF-8">
func get_html_enc_instant2(doc *goquery.Document) string {
	content, ok := doc.Find("head meta[charset]").Attr("charset")
	if !ok {
		return ""
	} else {
		return content
	}
}
func get_html_enc(doc *goquery.Document) string {
	var content string = ""
	content = get_html_enc_instant1(doc)
	if content == "" {
		content = get_html_enc_instant2(doc)
	}
	content = strings.ToLower(strings.Replace(content, "-", "", -1))
	if *debug {
		log.Printf("detected file encoding is [%s]\n", content)
	}
	return content
}
func run() {
	var (
		doc           *goquery.Document
		doc_err, err  error
		fd            *os.File
		file_encoding string = "-"
	)
        if *url == "-" {
            doc, doc_err = goquery.NewDocumentFromReader(os.Stdin)
        } else if fd, err = os.Open(*url); err == nil {
            doc, doc_err = goquery.NewDocumentFromReader(fd)
            defer fd.Close()
        } else {
            tmp_url := *url
            if !strings.HasPrefix(tmp_url,"http") {
                tmp_url = "http://" + tmp_url
            }
            doc, doc_err = goquery.NewDocument(tmp_url)
        }
	if doc_err != nil {
		log.Fatal("goquery NewDocument err:", doc_err)
	}

	if *debug {
		log.Printf("tag=[%s]\n", selector)
	}

	if !*noenc {
		file_encoding = get_html_enc(doc)
	}

	doc.Find(selector).Map(func(i int, sel *goquery.Selection) string {
		output := ""
		switch fun {
		case "html":
			if output, err = sel.Html(); err != nil {
				log.Fatal("select err:", err)
			}
			output = strings.TrimSpace(output)
		case "ohtml":
			if output, err = goquery.OuterHtml(sel); err != nil {
				log.Fatal("select err:", err)
			}
			output = strings.TrimSpace(output)
		case "text":
			output = strings.TrimSpace(sel.Text())
		case "attr":
			attr_list := strings.Split(*attr, ",")
			for _, attr_i := range attr_list {
				var output_i string
				if attr_i == "text" { // a hardcode case for convenience scrapy
					output_i = strings.TrimSpace(sel.Text())
				} else {
					output_i = sel.AttrOr(strings.TrimSpace(attr_i), "-")
				}
				if output_i == "" {
					output_i = "-"
				}
				output += output_i + "\t"
			}
		}

		if !*noenc && file_encoding != "" && file_encoding != "utf8" {
			if output, err = iconv.ConvertString(output, file_encoding, "utf8"); err != nil {
				log.Fatal("encoding invalid", err)
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
	} else {
		flag.Usage()
		os.Exit(-1)
	}

	switch {
	case *html:
		fun = "html"
	case *ohtml:
		fun = "ohtml"
	case *text:
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	run()
}
