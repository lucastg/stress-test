package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	statusCode int
	duration   time.Duration
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: A URL é obrigatória. Use --url para especificá-la.")
		return
	}

	var wg sync.WaitGroup
	results := make(chan Result, *requests)

	startTime := time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < *requests / *concurrency; j++ {
				start := time.Now()
				resp, err := http.Get(*url)
				duration := time.Since(start)

				if err != nil {
					fmt.Printf("[Worker %d] Request falhou: %s (Duração: %.2fs)\n", workerID, err.Error(), duration.Seconds())
					results <- Result{statusCode: 0, duration: duration}
					continue
				}

				fmt.Printf("[Worker %d] Request bem-sucedida: Status %d (Duração: %.2fs)\n", workerID, resp.StatusCode, duration.Seconds())
				results <- Result{statusCode: resp.StatusCode, duration: duration}
				resp.Body.Close()
			}
		}(i)
	}

	wg.Wait()
	close(results)

	endTime := time.Now()

	totalDuration := endTime.Sub(startTime)
	statusCounts := make(map[int]int)
	totalRequests := 0

	for res := range results {
		statusCounts[res.statusCode]++
		totalRequests++
	}

	report := fmt.Sprintf("\n========== Relatório de Testes de Carga ==========\n")
	report += fmt.Sprintf("URL Testada: %s\n", *url)
	report += fmt.Sprintf("Tempo Total de Execução: %.2fs\n", totalDuration.Seconds())
	report += fmt.Sprintf("Total de Requests Enviados: %d\n", totalRequests)
	report += fmt.Sprintf("\nDistribuição dos Status HTTP:\n")

	totalSucessos := 0
	for status, count := range statusCounts {
		if status == 200 {
			totalSucessos += count
		}
		report += fmt.Sprintf("  Status %d: %d requests\n", status, count)
	}
	report += fmt.Sprintf("\nTotal de Requests Bem-Sucedidos (Status 200): %d\n", totalSucessos)
	report += fmt.Sprintf("==================================================\n")

	fmt.Println(report)

	if err := saveReport("stress-test-report.txt", report); err != nil {
		fmt.Println("Erro ao salvar o relatório:", err)
	} else {
		fmt.Println("Relatório salvo em 'stress-test-report.txt'")
	}
}

func saveReport(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
