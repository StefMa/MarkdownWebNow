package markdownwebnow

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
  markdownString, err := readReadme(r)
  if err != nil {
    fmt.Fprintf(w, "Error: " + err.Error())
    return
  }
  client := &http.Client{}
  markdownBuffer := bytes.NewBufferString(*markdownString)
	req, err := http.NewRequest("POST", "https://" + r.Host + "/converter/convert.js", markdownBuffer)
  req.Header.Set("Content-Type", "text/plain")
	response, err := client.Do(req)
	defer response.Body.Close()
  if err != nil {
    fmt.Fprintf(w, "Error: " + err.Error())
    return
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Fprintf(w, "Error: " + err.Error())
    return
  }
  fmt.Fprintf(w, string(body))
}

func readReadme(r *http.Request) (*string, error) {
  // repo should be StefMa/MarkdownWebNow
  repo := r.URL.Query().Get("repo")
  rawReadmeUrl := "https://raw.githubusercontent.com/"+ repo + "/master/README.md"
  response, error := http.Get(rawReadmeUrl)
  if (error != nil) {
    return nil, error
  }
  defer response.Body.Close()
  body, error := ioutil.ReadAll(response.Body)
  if (error != nil) {
    return nil, error
  }
  readmeContentPointer := string(body)
  return &readmeContentPointer, nil
}
