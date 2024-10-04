package handlers

import (
	"bytes"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rchmachina/bpjs-tk/dto"
	"github.com/rchmachina/bpjs-tk/repositories"
	env "github.com/rchmachina/bpjs-tk/utils/env"
	"github.com/rchmachina/bpjs-tk/utils/helper"
)

type soal1Handler struct {
	soal1 repositories.Soal1Repository
	redis *redis.Client
}

func HandlerSoal1(soal1 repositories.Soal1Repository, client *redis.Client) *soal1Handler {
	return &soal1Handler{
		soal1: soal1,
		redis: client,
	}
}

func worker(w int, jobs <-chan struct{}, results chan<- int, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	for range jobs {
		// Generate a random UUID for each job
		id := uuid.New()
		format := "2006-01-02"

		// Call the helper function to generate date
		start, end := helper.CreateRandomDates(format)
		startDate := start
		endDate := end

		url := fmt.Sprint(env.GetConfigWithDefaultSetting("app.server.address_api", "http://localhost:8888"))
		url += "/soal1" // URL to POST data
		saldo, err := helper.RandomSaldo(900000, 1000000)

		if err != nil {
			log.Println("here")
		}
		//struct to send
		data := dto.PostDataSoal1{
			ID:        id,
			StartDate: startDate,
			EndDate:   endDate,
			Nominal:   saldo,
		}
		jsonData, err := json.Marshal(data) // Marshal the Data struct to JSON
		if err != nil {
			log.Printf("Worker %d: Error marshaling JSON: %v\n", id, err)
			results <- http.StatusInternalServerError
			continue
		}

		// Send POST request with UUID in the body
		resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Worker %d: Error fetching %s: %v\n", id, url, err)
			results <- http.StatusInternalServerError
			continue
		}
		defer resp.Body.Close()
		results <- resp.StatusCode // Send the HTTP status code to results channel
		if resp.StatusCode != 200 {
			fmt.Printf("Worker %d: Fetched %s, Status: %d, UUID: %s\n", w, url, resp.StatusCode, id.String())
		}
	}
}

// this api only for testing concurent
func (h *soal1Handler) PostData(c echo.Context) error {
	var dataPost dto.GetDataSoal1
	if err := c.Bind(&dataPost); err != nil {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	data, err := json.Marshal(dataPost)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("key-concurent-%s", dataPost.ID.String())
	err = h.redis.Set(c.Request().Context(), key, data, 0).Err() // Store UUID as the value
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save data to Redis")
	}

	keys, err := h.redis.Keys(c.Request().Context(), "key-concurent-*").Result()
	if err != nil {
		log.Printf("Could not retrieve keys: %v", err)
		return helper.JSONResponse(c, 400, "there is something wrong")
	}

	// Count the number of keys
	count := len(keys)

	// If the count is >= 1000, process them asynchronously
	if count >= 1000 {

		ctx := context.Background()
		go h.processKeysAsync(ctx, keys)

	}

	return helper.JSONResponse(c, 200, "Data received successfully")

}

func (h *soal1Handler) SendConcurentData(c echo.Context) error {
	const totalRequests = 1000
	var maxWorkers = (env.GetConfigWithDefaultSetting("app.db.worker.max_Workers_Post", 4))

	intMaxWorker, ok := maxWorkers.(int)
	if !ok {
		return helper.JSONResponse(c, 501, "there something wrong with server settings")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var wg sync.WaitGroup

	// Create channels for jobs and results
	jobs := make(chan struct{}, totalRequests)
	results := make(chan int, totalRequests)

	// Launch the workers for sending concurrent requests
	for w := 1; w <= intMaxWorker; w++ {
		wg.Add(1)
		go worker(w, jobs, results, client, &wg)
	}

	go func() {
		for i := 0; i < totalRequests; i++ {
			jobs <- struct{}{} // Send a signal to indicate a job
		}
		close(jobs)
	}()

	// Wait for all workers to finish

	wg.Wait()      // Wait for all workers to finish
	close(results) // Close results channel

	// Return immediately while workers process in the background
	return helper.JSONResponse(c, 200, "Workers started. Check console for output.")
}

func (h *soal1Handler) GetRedisData(c echo.Context) error {
	keys, err := h.redis.Keys(c.Request().Context(), "key-concurent-*").Result()
	if err != nil {
		log.Printf("Could not retrieve keys: %v", err)
		return helper.JSONResponse(c, 200, err)
	}
	return helper.JSONResponse(c, 200, keys)
}
func (h *soal1Handler) GetChildData(parentData dto.GetDataSoal1) ([]dto.PostChildData, error) {

	var childDatas []dto.PostChildData

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, parentData.StartDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
	}

	endDate, err := time.Parse(layout, parentData.EndDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return nil, err
	}
	dateRanges := helper.GenerateMonthlyRanges(startDate, endDate)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// Calculate the total number of days across all date ranges
	totalDays := 0
	for _, dr := range dateRanges {
		totalDays += helper.CountDays(dr.StartDate, dr.EndDate)
	}

	// Print the date ranges and protelar per month
	for _, dr := range dateRanges {
		var childData dto.PostChildData
		daysInMonth := helper.CountDays(dr.StartDate, dr.EndDate)
		protelar := (float64(daysInMonth) / float64(totalDays)) * parentData.Nominal
		childData.StartDate = dr.StartDate.Format("2006-01-02")
		childData.EndDate = dr.EndDate.Format("2006-01-02")
		childData.Nominal = math.Round(protelar*100) / 100

		//childData.ParentId = parentData.ID
		childDatas = append(childDatas, childData)
	}
	return childDatas, nil
}

func (h *soal1Handler) DeleteDataSoal1(c echo.Context) error {

	err := h.soal1.DeleteSoal1Repository()
	if err != nil {
		return helper.JSONResponse(c, 501, err.Error())

	}
	return helper.JSONBulkResponse(c, 200, "success deleting data")

}

func (h *soal1Handler) processKeysAsync(ctx context.Context, keys []string) {
	const numWorkers = 3 // Adjust the number of workers as needed
	var wg sync.WaitGroup
	var mu sync.Mutex // Create a mutex for synchronizing access to bulkDatas

	// Initialize struct for bulk data
	var bulkDatas []dto.DataSoal1

	// Create channels for jobs (keys) and results
	jobs := make(chan string, len(keys))

	// Worker function to process keys
	worker := func() {
		defer wg.Done() // Ensure that Done is called when the worker is done

		for key := range jobs {
			// Retrieve the value for each key using the provided context
			value, err := h.redis.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Could not get value for key %s: %v\n", key, err)
				continue
			}

			var parentData dto.GetDataSoal1
			err = json.Unmarshal([]byte(value), &parentData)
			if err != nil {
				log.Printf("Error unmarshalling value for key %s: %v\n", key, err)
				continue
			}
			mu.Lock()
			getChildData, err := h.GetChildData(parentData)
			if err != nil {
				log.Printf(err.Error())
			}

			bulkData := dto.DataSoal1{
				ParentData: parentData,
				ChildData:  getChildData,
			}

			// Lock the mutex before appending to bulkDatas

			bulkDatas = append(bulkDatas, bulkData)
			mu.Unlock() // Unlock the mutex after appending
		}
	}

	// Launch the workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker()
	}

	// Send keys to workers
	go func() {
		for _, key := range keys {
			jobs <- key // Send the key to the jobs channel
		}
		close(jobs) // Close the jobs channel when done
	}()

	// Wait for all workers to finish
	wg.Wait()

	// Prepare data for PostSoal1Repository
	var data dto.BulkDataSoal1
	data.Data = bulkDatas

	// PostSoal1Repository
	res := h.soal1.PostSoal1Repository(data)
	// Log the final aggregated parent data
	log.Println("Parent data bulk", res)

	_, err := h.redis.Del(ctx, keys...).Result() //delete all keys after success add data to db
	if err != nil {
		fmt.Errorf("could not delete keys: %v", err)
	}
}
