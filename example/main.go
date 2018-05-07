package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/jackwakefield/graff"
)

func main() {
	graph := graff.NewDirectedGraph()
	graph.AddNodes("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K")

	graph.AddEdge("A", "D")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "E")
	graph.AddEdge("E", "F")
	graph.AddEdge("F", "G")
	graph.AddEdge("H", "I")
	graph.AddEdge("H", "K")
	graph.AddEdge("I", "G")
	graph.AddEdge("J", "G")
	graph.AddEdge("K", "G")

	if err := previewGraph(graph); err != nil {
		log.Fatalln(err)
	}

	sorted, err := graph.DFSSort()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DFSIterativeSort:", sorted)
}

func previewGraph(graph *graff.DirectedGraph) error {
	inputFile, err := ioutil.TempFile("", "graff")
	if err != nil {
		return err
	}
	defer os.Remove(inputFile.Name())
	if err := ioutil.WriteFile(inputFile.Name(), []byte(graph.DOTGraph()), 0777); err != nil {
		return err
	}

	outputFile, err := ioutil.TempFile("", "graff")
	if err != nil {
		return err
	}
	defer os.Remove(outputFile.Name())

	if err := exec.Command("dot", "-Tsvg", inputFile.Name(), "-o", outputFile.Name()).Run(); err != nil {
		return err
	}

	output, err := ioutil.ReadFile(outputFile.Name())
	if err != nil {
		return err
	}
	url := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(output)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Unsupported platform")
	}
	if err != nil {
		return err
	}

	return nil
}
