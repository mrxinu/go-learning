/*
	Getting values out of a web form.
*/

package main

import "fmt"
import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Form["username"])
	fmt.Fprint(w, `
    <h1>A header</h1>
    <form>
      <div>
        <input type="text" name="navn">
				<input type="text" name="etternavn">
				<input type="submit" value="Submit">
			</div>
    </form>
	`)
	r.ParseForm()
	myName := r.FormValue("navn")
	mySurName := r.FormValue("etternavn")
	fmt.Println("Name given  = ", myName)
	fmt.Println("Surname = ", mySurName)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":7000", nil)
}

/*
<form>
  <div>
    <label for="uname">Choose a username: </label>
    <input type="text" id="uname" name="name">
  </div>
  <div>
    <button>Submit</button>
  </div>
</form>

*/
