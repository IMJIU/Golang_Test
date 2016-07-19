package main
import (
	"io"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
)
const ( 
	LOCAL_DIR = "E:/runtime-New_configuration/T04_socket/src/"
	UPLOAD_DIR = "E:/runtime-New_configuration/T04_socket/src/uploads"
)
func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
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
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
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
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	t, err := template.ParseFiles(LOCAL_DIR+tmpl + ".html")
	if err != nil {
		return
	}
	err = t.Execute(w, locals)
	return err
}

