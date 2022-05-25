package webpagegenerator

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Name      string
	Biography string
}

func GenerateFile(users []User) {
	f, err := os.Create("html/users.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	fmt.Fprintln(writer, "<!doctype html>")

	fmt.Fprintln(writer, "<html lang=\"en\">")
	fmt.Fprintln(writer, "<head>")
	fmt.Fprintln(writer, "<meta charset=\"utf-8\">")
	fmt.Fprintln(writer, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">")

	fmt.Fprintln(writer, "<title>User biographies</title>")

	fmt.Fprintln(writer, "<!-- Bootstrap core CSS -->")
	fmt.Fprintln(writer, "<link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css\" integrity=\"sha384-9gVQ4dYFwwWSjIDZnLEWnxCjeSWFphJiwGPXr1jddIhOegiu1FwO5qRGvFXOdJZ4\" crossorigin=\"anonymous\">")

	fmt.Fprintln(writer, "<!-- Custom styles for this template -->")
	fmt.Fprintln(writer, "<link href=\"assets/cover.css\" rel=\"stylesheet\">")
	fmt.Fprintln(writer, "</head>")

	fmt.Fprintln(writer, "<body class=\"text-center\">")

	fmt.Fprintln(writer, "<div class=\"cover-container d-flex w-100 h-100 p-3 mx-auto flex-column\">")
	fmt.Fprintln(writer, "<header class=\"masthead mb-auto\">")
	fmt.Fprintln(writer, "<div class=\"inner\">")
	fmt.Fprintln(writer, "<h3 class=\"masthead-brand\">Users Biography</h3>")
	fmt.Fprintln(writer, "<nav class=\"nav nav-masthead justify-content-center\">")
	fmt.Fprintln(writer, "<a class=\"nav-link active\" href=\"#\">Home</a>")
	fmt.Fprintln(writer, "<a class=\"nav-link\" href=\"#\">Features</a>")
	fmt.Fprintln(writer, "<a class=\"nav-link\" href=\"#\">Contact</a>")
	fmt.Fprintln(writer, "</nav>")
	fmt.Fprintln(writer, "</div>")
	fmt.Fprintln(writer, "</header>")

	fmt.Fprintln(writer, "<main role=\"main\" class=\"inner cover\">")
	for _, user := range users {
		fmt.Fprintf(writer, "<h1 class=\"cover-heading\">%s</h1>", user.Name)
		fmt.Fprintf(writer, "<p class=\"lead\">%s</p>", user.Biography)
	}
	fmt.Fprintln(writer, "</main>")

	fmt.Fprintln(writer, "<footer class=\"mastfoot mt-auto\">")
	fmt.Fprintln(writer, "<div class=\"inner\">")
	fmt.Fprintln(writer, "<p>Sprout class kata created by Codium <a href=\"https://twitter.com/CodiumTeam\">@CodiumTeam</a>. Cover template for <a href=\"https://getbootstrap.com/\">Bootstrap</a>, by <a href=\"https://twitter.com/mdo\">@mdo</a>.</p>")
	fmt.Fprintln(writer, "</div>")
	fmt.Fprintln(writer, "</footer>")
	fmt.Fprintln(writer, "</div>")

	fmt.Fprintln(writer, "<!-- Bootstrap core JavaScript")
	fmt.Fprintln(writer, "        ================================================== -->")
	fmt.Fprintln(writer, "<!-- Placed at the end of the document so the pages load faster -->")
	fmt.Fprintln(writer, "<script src=\"https://code.jquery.com/jquery-3.3.1.slim.min.js\" integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\" crossorigin=\"anonymous\"></script>")
	fmt.Fprintln(writer, "<script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js\" integrity=\"sha384-cs/chFZiN24E4KMATLdqdvsezGxaGsi4hLGOzlXwp5UZB1LY//20VyM2taTB4QvJ\" crossorigin=\"anonymous\"></script>")
	fmt.Fprintln(writer, "<script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js\" integrity=\"sha384-uefMccjFJAIv6A+rW+L4AHf99KvxDjWSu1z9VI8SKNVmz4sk7buKt/6v9KI65qnm\" crossorigin=\"anonymous\"></script>")
	fmt.Fprintln(writer, "</body>")
	fmt.Fprintln(writer, "</html>")

	writer.Flush()
}
