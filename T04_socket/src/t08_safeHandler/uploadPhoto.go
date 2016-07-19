package main
import (
	"io"
	"os"
	"log"
	"net/http"
	"path"
	"io/ioutil"
	"html/template"
)
const ( 
	LOCAL_DIR = "E:/runtime-New_configuration/T04_socket/src/"
	UPLOAD_DIR = "E:/runtime-New_configuration/T04_socket/src/uploads"
	TEMPLATE_DIR = LOCAL_DIR+"views"
)
var templates = make(map[string]*template.Template)
func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func main() {
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//方式1:
		if err := renderHtml(w, "upload", nil); err != nil{
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,http.StatusFound)
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
//		方式3:	
	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
	}
}
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{})(err error ){
	err = templates[tmpl+".html"].Execute(w, locals)
	return err
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
/**
	在上述这段代码中，我们巧妙地使用了defer关键字搭配recover()方法终结 panic 的肆行。
	safeHandler() 接收一个业务逻辑处理函数作为参数，同时调用这个业务逻辑处理函数。该业
	务逻辑函数执行完毕后， safeHandler()中defer指定的匿名函数会执行。倘若业务逻辑处理
	函数里边引发了panic，则调用recover()对其进行检测，若为一般性的错误，则输出HTTP 50x
	出错信息并记录日志，而程序将继续良好运行。
*/
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Println("WARN: panic in %v - %v", fn, e)
				//log.Println(string(debug.Stack())) 
			}
		}()
		fn(w, r)
	}
}
