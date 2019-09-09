package main
import(
    "bufio"
    "flag"
    "fmt"
    "encoding/csv"
    "io"
    "io/ioutil"
    "os"
    "strings"
)


func read_csv(csv_file string) ([]string, []string) {
    dat, _ := ioutil.ReadFile(csv_file)
    r := csv.NewReader(strings.NewReader(string(dat)))
    var questions []string
    var answers []string
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        questions = append(questions, record[0])
        answers = append(answers, record[1])
    }
    return questions, answers
}


func main() {
    // pdf viewer's name is "Document Viewer"
    var csv_file = flag.String("csv_file", "problems.csv",
                               "Path to the data csv file")
    flag.Parse()
    fmt.Printf("%q\n", *csv_file)
    reader := bufio.NewReader(os.Stdin)
    var correct int
    questions, answers := read_csv(*csv_file)
    for i := range questions {
        fmt.Print(questions[i],":")
        answer, _ := reader.ReadString('\n')
        if strings.Compare(answer, answers[i]) == 0 {
            correct ++;
        }
    }
    fmt.Printf("%d/%d\n", correct, len(questions))
}