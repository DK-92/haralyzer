package windows

import (
	"log"
	"strconv"
	"strings"

	"github.com/DK-92/harviewer/screens/model/table"
	"github.com/DK-92/harviewer/structs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type EntryWindow struct {
	*walk.MainWindow
	tabWidget *walk.TabWidget
}

var ew *EntryWindow

func MakeEntityWindow(filename string, entry *structs.Entry) {
	ew = new(EntryWindow)
	var db *walk.DataBinder

	if err := (MainWindow{
		AssignTo: &ew.MainWindow,
		Title:    filename + " - " + entry.Request.Method + " " + entry.Request.URL,
		MinSize:  Size{320, 240},
		Size:     Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "entry",
			DataSource:     entry,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			TabWidget{
				AssignTo: &ew.tabWidget,
				Pages: []TabPage{
					TabPage{
						Title:  "Request",
						Layout: VBox{},
						Children: []Widget{
							GroupBox{
								Title:  "Request information",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Method"},
									LineEdit{Text: Bind("Request.Method"), ReadOnly: true},
									Label{Text: "URL"},
									LineEdit{Text: Bind("Request.URL"), ReadOnly: true},
									Label{Text: "HTTP Version"},
									LineEdit{Text: Bind("Request.HTTPVersion"), ReadOnly: true},
								},
							},
							GroupBox{
								Title:  "Headers",
								Layout: VBox{},
								Children: []Widget{
									TableView{
										Columns: []TableViewColumn{
											{Title: "Name", Width: 200},
											{Title: "Value", Width: 300},
											{Title: "Comment", Width: 100},
										},
										Model: table.NewHeaderModel(&entry.Request.Headers),
									},
								},
							},
							GroupBox{
								Title:  "Query string(s)",
								Layout: VBox{},
								Children: []Widget{
									TableView{
										Columns: []TableViewColumn{
											{Title: "Name", Width: 200},
											{Title: "Value", Width: 300},
											{Title: "Comment", Width: 100},
										},
										Model: table.NewQueryStringModel(&entry.Request.QueryString),
									},
								},
							},
							GroupBox{
								Title:  "Post data",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									VSplitter{
										Column: 1,
										Children: []Widget{
											Composite{
												Layout: Grid{Columns: 2},
												Children: []Widget{
													Label{Text: "MimeType"},
													LineEdit{Text: Bind("Request.PostData.MimeType"), ReadOnly: true},
													Label{Text: "Comment"},
													LineEdit{Text: Bind("Request.PostData.Comment"), ReadOnly: true},
												},
											},
											TableView{
												Columns: []TableViewColumn{
													{Title: "Name", Width: 100},
													{Title: "Value", Width: 100},
													{Title: "Filename", Width: 100},
													{Title: "Content-type", Width: 100},
													{Title: "Comment", Width: 100},
												},
												Model: table.NewParamModel(&entry.Request.PostData.Params),
											},
										},
									},
									VSplitter{
										Column: 2,
										Children: []Widget{
											Label{Text: "Body"},
											TextEdit{
												Text:     Bind("Request.PostData.Text"),
												ReadOnly: true,
												Font:     Font{Family: "Consolas", PointSize: 10},
												HScroll:  true,
												VScroll:  true,
											},
										},
									},
								},
							},
						},
					},
					TabPage{
						Title:  "Response",
						Layout: VBox{},
						Children: []Widget{
							GroupBox{
								Title:  "Response information",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Response"},
									LineEdit{Text: strconv.Itoa(entry.Response.Status) + " " + entry.Response.StatusText, ReadOnly: true}, // LineEdit doesn't work with ints
									Label{Text: "HTTP Version"},
									LineEdit{Text: Bind("Response.HTTPVersion"), ReadOnly: true},
								},
							},
							GroupBox{
								Title:  "Headers",
								Layout: VBox{},
								Children: []Widget{
									TableView{
										Columns: []TableViewColumn{
											{Title: "Name", Width: 200},
											{Title: "Value", Width: 300},
											{Title: "Comment", Width: 100},
										},
										Model: table.NewHeaderModel(&entry.Response.Headers),
									},
								},
							},
							GroupBox{
								Title:  "Content",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Size in bytes"},
									LineEdit{Text: strconv.Itoa(entry.Response.Content.Size), ReadOnly: true},
									Label{Text: "Bytes saved by compression"},
									LineEdit{Text: strconv.Itoa(entry.Response.Content.Compression), ReadOnly: true},
									Label{Text: "MimeType"},
									LineEdit{Text: Bind("Response.Content.MimeType"), ReadOnly: true},
								},
							},
							GroupBox{
								Title:  "Response body",
								Layout: VBox{},
								Children: []Widget{
									ScrollView{
										Layout: VBox{MarginsZero: true},
										Children: []Widget{
											TextEdit{
												Text:     strings.ReplaceAll(entry.Response.Content.Text, "\n", "\r\n"),
												ReadOnly: true,
												Font:     Font{Family: "Consolas", PointSize: 10},
											},
										},
									},
								},
							},
						},
					},
					TabPage{
						Title:  "Cookies",
						Layout: Grid{Columns: 2},
						Children: []Widget{
							TableView{
								Columns: []TableViewColumn{
									{Title: "Name", Width: 200},
									{Title: "Value", Width: 300},
									{Title: "Path", Width: 100},
									{Title: "Domain", Width: 100},
									{Title: "Expires", Width: 100},
									{Title: "HTTPOnly", Width: 100},
									{Title: "Secure", Width: 100},
									{Title: "Comment", Width: 100},
								},
								Model: table.NewCookieModel(&entry.Response.Cookies),
							},
						},
					},
					TabPage{
						Title:  "Timing",
						Layout: VBox{Alignment: AlignHNearVNear},
						Children: []Widget{
							GroupBox{
								Title:  "Timings (ms)",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Blocked"},
									LineEdit{Text: strconv.Itoa(entry.Timings.Blocked), ReadOnly: true},
									Label{Text: "DNS"},
									LineEdit{Text: strconv.Itoa(entry.Timings.DNS), ReadOnly: true},
									Label{Text: "Connect"},
									LineEdit{Text: strconv.Itoa(entry.Timings.Connect), ReadOnly: true},
									Label{Text: "Send"},
									LineEdit{Text: strconv.Itoa(entry.Timings.Send), ReadOnly: true},
									Label{Text: "Wait"},
									LineEdit{Text: strconv.Itoa(entry.Timings.Wait), ReadOnly: true},
									Label{Text: "Receive"},
									LineEdit{Text: strconv.Itoa(entry.Timings.Receive), ReadOnly: true},
									Label{Text: "SSL"},
									LineEdit{Text: strconv.Itoa(entry.Timings.SSL), ReadOnly: true},
									Label{Text: "Comment"},
									LineEdit{Text: Bind("Timings.Comment"), ReadOnly: true},
								},
							},
						},
					},
					TabPage{
						Title:  "Cache",
						Layout: HBox{Alignment: AlignHNearVNear},
						Children: []Widget{
							GroupBox{
								Title:  "Cache state before request",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Expires"},
									LineEdit{Text: Bind("Cache.BeforeRequest.Expires"), ReadOnly: true},
									Label{Text: "Last access"},
									LineEdit{Text: Bind("Cache.BeforeRequest.LastAccess"), ReadOnly: true},
									Label{Text: "ETag"},
									LineEdit{Text: Bind("Cache.BeforeRequest.ETag"), ReadOnly: true},
									Label{Text: "Hit count"},
									LineEdit{Text: strconv.Itoa(entry.Cache.BeforeRequest.HitCount), ReadOnly: true},
									Label{Text: "Comment"},
									LineEdit{Text: Bind("Cache.BeforeRequest.Comment"), ReadOnly: true},
								},
							},
							GroupBox{
								Title:  "Cache state after request",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Expires"},
									LineEdit{Text: Bind("Cache.AfterRequest.Expires"), ReadOnly: true},
									Label{Text: "Last access"},
									LineEdit{Text: Bind("Cache.AfterRequest.LastAccess"), ReadOnly: true},
									Label{Text: "ETag"},
									LineEdit{Text: Bind("Cache.AfterRequest.ETag"), ReadOnly: true},
									Label{Text: "Hit count"},
									LineEdit{Text: strconv.Itoa(entry.Cache.AfterRequest.HitCount), ReadOnly: true},
									Label{Text: "Comment"},
									LineEdit{Text: Bind("Cache.AfterRequest.Comment"), ReadOnly: true},
								},
							},
						},
					},
				},
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	ew.tabWidget.SetCurrentIndex(0)

	ew.Run()
}
